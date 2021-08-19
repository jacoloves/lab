package main

import 	(
	"bufio"
	"bytes"
	"fmt"
	"go/token"
	"io"
	"log"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

/*** defines ***/

const GILO_VERSION = "0.0.1"
const GILO_TAB_STOP = 8
const GILO_QUIT_TIMES = 3

const (
	HL_HIGHLIGHT_NUMBERS = 1 << 0
	HL_HIGHLIGHT_STRINGS = 1 << iota
)

const (
	BACKSPACE   = 127
	ARROW_LEFT  = 1000 + iota
	ARROW_RIGHT = 1000 + iota
	DEL_KEY     = 1000 + iota
	HOME_KEY    = 1000 + iota
	END_KEY     = 1000 + iota
	PAGE_UP     = 1000 + iota
	PAGE_DOWN   = 1000 + iota
)

const (
	HL_NORMAL    = 0
	HL_COMMENT   = iota 
	HL_MLCOMMENT = iota 
	HL_KEYWORD1  = iota 
	HL_KEYWORD2  = iota 
	HL_STRING	   = iota 
	HL_NUMBER    = iota 
	HL_MATCH     = iota 
)

/*** data ***/

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [20]byte
	Ispeed uint32
	Ospeed uint32
}

type editorSyntax struct {
	filetype               string
	filematch              []string
	keywords               []string
	singleLineCommentStart []byte
	multiLineCommentStart  []byte
	multiLineCommentEnd    []byte
	flags                  int
}

type erow struct {
	idx           int
	size          int
	rsize         int
	chars         []byte
	render        []byte
	hl            []byte
	hlOpenComment bool
}

type editorConfig struct {
	cx             int
	cy             int
	rx             int
	rowoff         int
	coloff         int
	screenRows     int
	screenCols     int
	numRows        int
	rows           []erow
	dirty          bool
	filename       string
	statusmsg      string
	statusmsg_time time.Time
	syntax         *editorSyntax
	origTermios    *Termios
}

type WinSize struct {
	Row uint16
	Col uint16
	Xpixel uint16
	Ypixel uint16
}

var E editorConfig

/*** filetypes ***/
var HLDB []editorSyntax = []editorSyntax {
	editorSyntax{
		filetype: "c",
		filematch: []string{".c", ".h", ".cpp"},
		keywords: []string{"switch", "if", "while", "for",
				"break", "continue", "return", "else", "struct",
				"union", "typedef", "static", "enum", "class", "case",
				"int|", "long|", "double|", "float|", "char|",
				"unsigned|", "signed|", "void|",
			},
		singleLineCommentStart: []byte{'/', '/'},
		multiLineCommentStart: []byte{'/', '*'},
		multiLineCommentEnd: []byte{'*', '/'},
		flags:HL_HIGHLIGHT_NUMBERS|HL_HIGHLIGHT_STRINGS,
	},
}

/*** terminal ***/

func die(err error) {
	disableRawMode()
	io.WriteString(os.Stdout, "\x1b[2J")
	io.WriteString(os.Stdout, "\x1b[H")
	log.Fatal(err)
}

func TcSetAttr(fd uintptr, termios *Termios) error {
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(syscall.TCSETS+1), uintptr(unsafe.Pointer(termios))); err != 0 {
		return err
	}
	return nil
}

func TcGetAttr(fd uintptr) *Termios {
	var termios = &Termios{}
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, syscall.TCGETS, uintptr(unsafe.Pointer(termios))); err != 0 {
		log.Fatalf("Problem getting terminal attributes: %s\n", err)
	}
	return termios
}

func enableRawMode() {
	E.origTermios = TcGetAttr(os.Stdin.Fd())
	var raw Termios
	raw = *E.origTermios
	raw.Iflag &^= syscall.BRKINT | syscall.ICRNL | syscall.INPCK | syscall.ISTRIP | syscall.IXON
	raw.Oflag &^= syscall.OPOST
	raw.Cflag |= syscall.CS8
	raw.Lflag &^= syscall.ECHO | syscall.ICANON | syscall.IEXTEN | syscall.ISIG
	raw.Cc[syscall.VMIN+1] = 0
	raw.Cc[syscall.VTIME+1] = 1
	if e := TcSetAttr(os.Stdin.Fd(), &raw); e != nil {
		log.Fatalf("Problem enabling raw mode: %s\n", e)
	}
}

func disableRawMode() {
	if e := TcSetAttr(os.Stdin.Fd(), E.origTermios); e != nil{
		log.Fatalf("Problem disabling raw mode: %s\n", e)
	}
}

func getCursorPosition(rows *int, cols *int) int {
	io.WriteString(os.Stdout, "\x1b[6n")
	var buffer [1]byte
	var buf []byte
	var cc int
	for cc, _ = os.Stdin.Read(buffer[:]); cc == 1; cc, _ = os.Stdin.Read(buffer[:]) {
		if buffer[0] == 'R' {
			break
		}
		buf = append(buf, buffer[0])
	}
	if string(buf[0:2]) != "\x1b[" {
		log.Printf("Failed to read rows;cols from tty\n")
		return -1
	}
	if n, e := fmt.Sscanf(string(buf[2:]), "%d;%d", rows, cols); n != 2 || e != nil {
		if e != nil {
			log.Printf("getCursorPosition: fmt.Sscanf() failed: %s\n", e)
		}
		if n != 2 {
			log.Printf("getCursorPosition: got %d items, wanted 2\n", n)
		}
		return -1
	}
	return 0
}

func getWindowSize(rows *int, cols *int) int {
	var w WinSize
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL,
			os.Stdout.Fd(),
			syscall.TIOCGWINSZ,
			uintptr(unsafe.Pointer(&w)),
		)
	if err != 0 {
		io.WriteString(os.Stdout, "\x1b[999C\x1b[999B")
		return getCursorPosition(rows, cols)
	} else {
		*rows = int(w.Row)
		*cols = int(w.Col)
		return 0
	}
	return -1
}

/*** syntax hightlighting ***/

func editorUpdateSyntax(row *erow) {
	row.hl = make([]byte, row.rsize)
	if E.syntax == nil { return }
	keywords := E.syntax.keywords[:]
	scs := E.syntax.singleLineCommentStart
	mcs := E.syntax.multiLineCommentStart
	mce := E.syntax.multiLineCommentEnd
	prevSep := true
	inComment := row.idx > 0 && E.rows[row.idx-1].hlOpenComment
	var inString byte = 0
	var skip = 0
	for i, c := range row.render {
		if skip > 0 {
			skip--
			continue
		}
		if inString == 0 && len(scs) > 0 && !inComment {
			if bytes.HasPrefix(row.render[i:], scs) {
				for j := i; j < row.rsize; j++ {
					row.hl[j] = HL_COMMENT
				}
				break
			}
		}
		if inString == 0 && len(mcs) > 0 && len(mce) > 0 {
			if inComment {
				row.hl[i] = HL_COMMENT
				if bytes.HasPrefix(row.render[i:], mce) {
					for l := i; l < i + len(mce); i++ {
						row.hl[l] = HL_MLCOMMENT
					}
					skip = len(mce)
					inComment = false
					prevSep = true
				}
				continue
		} else if bytes.HasPrefix(row.render[i:], mcs) {
			for l := i; l < i + len(mcs); i++ {
				roe.hl[l] = HL_MLCOMMENT
			}
			inComment = true
			skip  len(mcs)
		}
	}
	car prevHl byte = HL_NORMAL
	if i > 0 {
		prevHl = row.hl[i-1]
	}
	if (E.syntax.flags & HL_HIGHLIGHT_STRINGS) == HL_HIGHLIGHT_STRINGS {
		if inString != 0 {
			row.hl[i] = HL_STRING
			if c == '\\' && i + 1 < row.rsize {
				row.hl[i+1] = HL_STRING
				skip = 1
				continue
			}
			if c == inString { inString = 0 }
			prevSep = true
			continue
		} else {
			if c == '"' || c == '\'' {
				inString = c
				row.hl[i] = HL_STRING
				continue
			}
		}
	}
	if (E.syntax.flags & HL_HIGHLIGHT_NUMBERS) == HL_HIGHLIGHT_NUMBERS {
		if unicode.IsDigit(rune(c)) &&
				(prevSep || prevHl == HL_NUMBER) ||
				(c == '.' && prevHl == HL_NUMBER) {
					row.hl[i] = HL_NUMBER
					prevSep = false
					continue
				}
	}
	if prevSep {
		var j int
		var skw string
		for j, skw = range keywords {
			kw := []byte(skw)
			var color byte = HL_KEYWORD1
			idx := bytes.LastIndexByte(kw, '|')
			if idx > 0 {
				kw = kw[:idx]
				color = HL_KEYWORD2
			}
			klen := len(kw)
			if bytes.HasPrefix(row.render[i:], kw) &&
							(len(row.render[i:]) == klen ||
							isSeparator(row.render[i+klen]))	{
								for l := i; l < i+klen; l++ {
									row.hl[i] = color
								}
								skip = klen - 1
								break
							}
		}
		if j < len(keywords) - 1 {
			prevSep = false
			continue
		}
	}
	prevSep = isSeparator(c)
}

func editorSelectSyntaxHighlight() {
	if E.filename == "" { return }

	for _, s := range HLDB {
		for _, suffix := range s.filematch {
			if strings.HasSuffix(E.filename, suffix) {
				E.syntax = &s
				return
			}
		}
	}
}

/*** row operations ***/

func editorUpdateRow(row *erow) {
	tabs := 0
	for _, c := range row.chars {
		if c == '\t' {
			tabs++
		}
	}

	row.render = make([]byte, row.size + tabs*(GILO_TAB_STOP - 1))

	idx := 0
	for _, c := range row.chars {
		if c == '\t' {
			row.render[idx] = ' '
			idx++
			for (idx%GILO_TAB_STOP) != 0 {
				row.render[idx] = ' '
				idx++
			}
		} else {
			row.render[idx] = c
			idx++
		}
	}
	row.rsize = idx
	editorUpdateSyntax(row)
}

func editorInsertRow(at int, s []byte) {
	if at < 0 || at > E.numRows {
		return
	}
	var r erow
	r.chars = s
	r.size = len(s)
	r.idx = at

	if at == 0 {
		t := make([]erow, 1)
		t[0] = r
		E.rows = append(t, E.rows...)
	} else if at == E.numRows {
		E.rows = append(E.rows, r)
	} else {
		t := make([]erow, 1)
		t[0] = r
		E.rows = append(E.rows[:at], append(t, E.rows[at:]...)...)
	}

	for j := at + 1; j <= E.numRows; j++ {
		E.rows[j].idx++
	}

	editorUpdateRow(&E.rows[at])
	E.numRows++
	E.dirty = true

}

/*** file I/O ***/

func editorOpen(filename string) {
	E.filename = filename
	editorSelectSyntaxHighlight()
	fd, err := os.Open(filename)
	if err != nil {
		die(err)
	}
	defer fd.Close()
	fp := bufio.NewReader(fd)

	for line, err := fp.ReadBytes('\n'); err == nil; err = fp.ReadBytes('\n') {
		for c := line[len(line) -1]; len(line) > 0 && (c == '\n' || c == '\r'); {
			line = line[:len(line)-1]
			if len(line) > 0 {
				c = line[len(line) - 1]
			}
		}
		editorInsertRow(E.numRows, line)
	}

	if err != nil && err != io.EOF {
		die(err)
	}
	E.dirty = false
}

/*** init ***/

func initEditor() {
	if getWindowSize(&E.screenRows, &E.screenCols) == -1 {
		die(fmt.Errorf("couldn't get screen size"))
	}
	E.screenRows -= 2
}

func main() {
	enableRawMode()
	defer disableRawMode()
	initEditor()
	if len(os.Args) > 1 {
		editorOpen(os.Args[1])
	}
}

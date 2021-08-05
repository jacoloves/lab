package main

import (
	"log"
	"os"
	"syscall"
	"time"
	"unsafe"
)

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

var E editorConfig

/*** terminal ***/
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

func main() {
	enableRawMode()
}

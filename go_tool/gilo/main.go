package main

import (
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
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(syscall.TCSETS+1), uintptr(unsafe.Pointer(termios)))
	if err != 0 {
		return err
	}
	return nil
}

func enableRawMode() {
	E.origTermios = TcGetAttr(os.Stdin.Fd())
}

func main() {
	enableRawMode()
}

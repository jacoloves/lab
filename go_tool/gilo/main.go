package main

import (
	_ "bufio"
	"fmt"
	"syscall"

	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
)

var orig_termios unix.Termios

func disableRawMode() {
	termios.Tcsetattr(uintptr(syscall.Stdin), unix.TCSAFLUSH, &orig_termios)
}

func enableRawMode() {
	termios.Tcgetattr(uintptr(syscall.Stdin), &orig_termios)

	raw := orig_termios
	raw.Lflag &= syscall.ECHO | syscall.ICANON

	termios.Tcsetattr(uintptr(syscall.Stdin), unix.TCSAFLUSH, &raw)
}

func main() {
	enableRawMode()

	/*
		stdin := bufio.NewScanner(os.Stdin)
		for stdin.Scan() {
			if err := stdin.Err(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			reader := bufio.NewReader(os.Stdin)
			r, _, _ := reader.ReadRune()

			if unicode.IsControl(r) {
				fmt.Printf("%d\n", r)
			} else {
				fmt.Printf("%d('%c')\n", r, r)
			}

			if stdin.Text() == "q" {
				break
			}
		}
	*/
	/*
		tty, err := os.Open("/dev/tty")
		if err != nil {
			log.Fatal(err)
		}
		defer tty.Close()

		reader := bufio.NewReader(tty)
		r, _, _ := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		if unicode.IsControl(r) {
			fmt.Printf("%c\n", r)
		} else {
			fmt.Printf("%d('%c')\n", r, r)
		}
	*/
	for {
		buf := make([]byte, 1024)
		if num, _ := syscall.Read(syscall.Stdin, buf); num == 1 && string(buf) != "q" {
			break
		}

		fmt.Printf("%d", buf)
	}

	defer disableRawMode()
}

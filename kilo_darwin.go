package main

import (
	"log"

	"golang.org/x/sys/unix"
)

func disableRawMode() {
	if err := unix.IoctlSetTermios(unix.Stdin, unix.TIOCSETAF, E.origTermios); err != nil {
		log.Fatalf("Problem disabling raw mode: %s\n", err)
	}
}

func enableRawMode() {
	var err error
	E.origTermios, err = unix.IoctlGetTermios(unix.Stdin, unix.TIOCGETA)
	if err != nil {
		log.Fatalf("Problem getting termios: %s\n", err)
	}
	raw := *E.origTermios
	raw.Iflag &^= unix.BRKINT | unix.ICRNL | unix.INPCK | unix.ISTRIP | unix.IXON
	raw.Oflag &^= unix.OPOST
	raw.Cflag &^= unix.CS8
	raw.Lflag &^= unix.ECHO | unix.ICANON | unix.IEXTEN | unix.ISIG
	raw.Cc[unix.VMIN] = 0
	raw.Cc[unix.VTIME] = 1
	if err := unix.IoctlSetTermios(unix.Stdin, unix.TIOCSETAF, &raw); err != nil {
		log.Fatalf("Problem enabling raw mode: %s\n", err)
	}
}

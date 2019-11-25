// +build darwin dragonfly freebsd netbsd openbsd

package getch

import (
	"golang.org/x/sys/unix"
)

const (
	ioctlReadTermios  = unix.TIOCGETA
	ioctlWriteTermios = unix.TIOCSETA
)

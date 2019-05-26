// +build !windows

package getch

import (
	"golang.org/x/sys/unix"
)

type state struct {
	termios unix.Termios
}

func makeRaw(fd int) (*state, error) {
	termios, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	if err != nil {
		return nil, err
	}

	termios.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON

	if err := unix.IoctlSetTermios(fd, ioctlWriteTermios, termios); err != nil {
		return nil, err
	}

	return &state{termios: *termios}, nil
}

func restored(fd int, state *state) error {
	return unix.IoctlSetTermios(fd, ioctlWriteTermios, &state.termios)
}

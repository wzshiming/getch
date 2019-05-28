// +build !windows

package getch

import (
	"syscall"

	"golang.org/x/sys/unix"
)

type state struct {
	termios unix.Termios
}

func read(p []byte) (int, error) {
	return syscall.Read(int(hStdin), p)
}

func makeRaw() (*state, error) {
	termios, err := unix.IoctlGetTermios(int(hStdin), ioctlReadTermios)
	if err != nil {
		return nil, err
	}

	termios.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON

	if err := unix.IoctlSetTermios(int(hStdin), ioctlWriteTermios, termios); err != nil {
		return nil, err
	}

	return &state{termios: *termios}, nil
}

func restored(state *state) error {
	return unix.IoctlSetTermios(int(hStdin), ioctlWriteTermios, &state.termios)
}

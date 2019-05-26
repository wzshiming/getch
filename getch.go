package getch

import (
	"io"
	"syscall"
)

// Getch get the pressed key directly without buffering, no need to enter enter to get.
func Getch() (rune, []byte, error) {
	var buf [6]byte
	fd := 0
	state, err := makeRaw(fd)
	if err != nil {
		return 0, nil, err
	}
	defer restored(fd, state)

	n, err := syscall.Read(fd, buf[:])
	if err != nil {
		return 0, nil, err
	}
	if n == 0 {
		return 0, nil, io.ErrUnexpectedEOF
	}
	key, raw := bytes2Key(buf[:n], true)
	return key, raw, nil
}

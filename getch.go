package getch

import (
	"io"
	"syscall"
)

// Getch get the pressed key directly without buffering, no need to enter enter to get.
func Getch() (rune, []byte, error) {
	var buf [6]byte

	state, err := makeRaw(0)
	if err != nil {
		return 0, nil, err
	}
	defer restored(0, state)

	n, err := syscall.Read(0, buf[:])
	if err != nil {
		return 0, nil, err
	}
	if n == 0 {
		return 0, nil, io.ErrUnexpectedEOF
	}
	key, raw := bytes2Key(buf[:n])
	return key, raw, nil
}

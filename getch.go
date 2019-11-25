package getch

import (
	"io"
	"os"
)

var (
	hStdin = os.Stdin.Fd()
)

// Getch get the pressed key directly without buffering, no need to enter enter to get.
func Getch() (rune, []byte, error) {
	return GetchBy(hStdin)
}

// GetchBy get the pressed key directly without buffering, no need to enter enter to get.
func GetchBy(hStdin uintptr) (rune, []byte, error) {
	var buf [6]byte

	state, err := makeRaw(hStdin)
	if err != nil {
		return 0, nil, err

	}
	defer restored(hStdin, state)

	n, err := read(hStdin, buf[:])
	if err != nil {
		return 0, nil, err
	}
	if n == 0 {
		return 0, nil, io.ErrUnexpectedEOF
	}
	key, raw := bytes2Key(buf[:n])
	return key, raw, nil
}

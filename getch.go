package getch

import (
	"errors"
	"io"
	"os"
)

var (
	hStdin           = os.Stdin.Fd()
	isTerm           = isTerminal()
	errNotInTerminal = errors.New("error not in terminal")
)

// Getch get the pressed key directly without buffering, no need to enter enter to get.
func Getch() (rune, []byte, error) {
	return GetchBy(hStdin)
}

// GetchBy get the pressed key directly without buffering, no need to enter enter to get.
func GetchBy(hStdin uintptr) (rune, []byte, error) {
	if !isTerm {
		return 0, nil, errNotInTerminal
	}

	state, err := makeRaw(hStdin)
	if err != nil {
		return 0, nil, err

	}
	defer restored(hStdin, state)

	var buf [6]byte
	n, err := read(hStdin, buf[:])
	if err != nil {
		return 0, nil, err
	}
	if n == 0 {
		return 0, nil, io.ErrUnexpectedEOF
	}
	key, raw := Bytes2Key(buf[:n])
	return key, raw, nil
}

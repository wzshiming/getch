package getch

import (
	"golang.org/x/sys/windows"
)

type state struct {
	mode uint32
}

func makeRaw(fd int) (*state, error) {
	var st uint32
	if err := windows.GetConsoleMode(windows.Handle(fd), &st); err != nil {
		return nil, err
	}
	raw := st &^ (windows.ENABLE_ECHO_INPUT | windows.ENABLE_PROCESSED_INPUT | windows.ENABLE_LINE_INPUT | windows.ENABLE_PROCESSED_OUTPUT)
	if err := windows.SetConsoleMode(windows.Handle(fd), raw); err != nil {
		return nil, err
	}
	return &state{st}, nil
}

func restored(fd int, state *state) error {
	return windows.SetConsoleMode(windows.Handle(fd), state.mode)
}

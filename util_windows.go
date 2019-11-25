package getch

import (
	"syscall"

	"golang.org/x/sys/windows"
)

type state struct {
	mode uint32
}

func read(hStdin uintptr, p []byte) (int, error) {
	return syscall.Read(syscall.Handle(hStdin), p)
}

func makeRaw(hStdin uintptr) (*state, error) {
	var st uint32
	if err := windows.GetConsoleMode(windows.Handle(hStdin), &st); err != nil {
		return nil, err
	}
	raw := st &^ (windows.ENABLE_ECHO_INPUT | windows.ENABLE_PROCESSED_INPUT | windows.ENABLE_LINE_INPUT | windows.ENABLE_PROCESSED_OUTPUT)
	if err := windows.SetConsoleMode(windows.Handle(hStdin), raw); err != nil {
		return nil, err
	}
	return &state{st}, nil
}

func restored(hStdin uintptr, state *state) error {
	return windows.SetConsoleMode(windows.Handle(hStdin), state.mode)
}

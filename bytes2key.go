package getch

import (
	"bytes"
	"unicode/utf8"
)

const (
	keyEscape    = 27
	keyBackspace = 127
	keyUnknown   = 0xd800 /* UTF-16 surrogate area */ + iota
	keyUp
	keyDown
	keyLeft
	keyRight
	keyAltLeft
	keyAltRight
	keyHome
	keyEnd
	keyDeleteWord
	keyDeleteLine
	keyClearScreen
	keyPasteStart
	keyPasteEnd
)

var (
	pasteStart = []byte{keyEscape, '[', '2', '0', '0', '~'}
	pasteEnd   = []byte{keyEscape, '[', '2', '0', '1', '~'}
)

func bytes2Key(b []byte, pasteActive bool) (rune, []byte) {
	if len(b) == 0 {
		return utf8.RuneError, nil
	}

	if !pasteActive {
		switch b[0] {
		case 1: // ^A
			return keyHome, b[1:]
		case 5: // ^E
			return keyEnd, b[1:]
		case 8: // ^H
			return keyBackspace, b[1:]
		case 11: // ^K
			return keyDeleteLine, b[1:]
		case 12: // ^L
			return keyClearScreen, b[1:]
		case 23: // ^W
			return keyDeleteWord, b[1:]
		case 14: // ^N
			return keyDown, b[1:]
		case 16: // ^P
			return keyUp, b[1:]
		}
	}

	if b[0] != keyEscape {
		if !utf8.FullRune(b) {
			return utf8.RuneError, b
		}
		r, l := utf8.DecodeRune(b)
		return r, b[l:]
	}

	if !pasteActive && len(b) >= 3 && b[0] == keyEscape && b[1] == '[' {
		switch b[2] {
		case 'A':
			return keyUp, b[3:]
		case 'B':
			return keyDown, b[3:]
		case 'C':
			return keyRight, b[3:]
		case 'D':
			return keyLeft, b[3:]
		case 'H':
			return keyHome, b[3:]
		case 'F':
			return keyEnd, b[3:]
		}
	}

	if !pasteActive && len(b) >= 6 && b[0] == keyEscape && b[1] == '[' && b[2] == '1' && b[3] == ';' && b[4] == '3' {
		switch b[5] {
		case 'C':
			return keyAltRight, b[6:]
		case 'D':
			return keyAltLeft, b[6:]
		}
	}

	if !pasteActive && len(b) >= 6 && bytes.Equal(b[:6], pasteStart) {
		return keyPasteStart, b[6:]
	}

	if pasteActive && len(b) >= 6 && bytes.Equal(b[:6], pasteEnd) {
		return keyPasteEnd, b[6:]
	}

	for i, c := range b[0:] {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '~' {
			return keyUnknown, b[i+1:]
		}
	}

	return utf8.RuneError, b
}

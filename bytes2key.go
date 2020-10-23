package getch

import (
	"unicode/utf8"
)

func Bytes2Key(b []byte) (rune, []byte) {
	if len(b) == 0 {
		return utf8.RuneError, nil
	}

	switch b[0] {
	case byte(KeyCtrlA):
		return rune(KeyHome), b[1:]
	case byte(KeyCtrlB):
		return rune(KeyLeft), b[1:]
	case byte(KeyCtrlE):
		return rune(KeyEnd), b[1:]
	case byte(KeyCtrlF):
		return rune(KeyRight), b[1:]
	case byte(KeyCtrlH):
		return rune(KeyDel), b[1:]
	case byte(KeyCtrlK):
		return rune(KeyDeleteLine), b[1:]
	case byte(KeyCtrlL):
		return rune(KeyClearScreen), b[1:]
	case byte(KeyCtrlW):
		return rune(KeyDeleteWord), b[1:]
	case byte(KeyCtrlN):
		return rune(KeyDown), b[1:]
	case byte(KeyCtrlP):
		return rune(KeyUp), b[1:]
	}

	if b[0] != byte(KeyCtrlLBrack) {
		if !utf8.FullRune(b) {
			return utf8.RuneError, b
		}
		r, l := utf8.DecodeRune(b)
		return r, b[l:]
	}

	if len(b) >= 3 && b[0] == byte(KeyCtrlLBrack) && b[1] == '[' {
		switch b[2] {
		case 'A':
			return rune(KeyUp), b[3:]
		case 'B':
			return rune(KeyDown), b[3:]
		case 'C':
			return rune(KeyRight), b[3:]
		case 'D':
			return rune(KeyLeft), b[3:]
		case 'H':
			return rune(KeyHome), b[3:]
		case 'F':
			return rune(KeyEnd), b[3:]
		}
	}

	if len(b) >= 6 && b[0] == byte(KeyCtrlLBrack) && b[1] == '[' && b[2] == '1' && b[3] == ';' && b[4] == '3' {
		switch b[2] {
		case '1':
			if b[3] == ';' && b[4] == '3' {
				switch b[5] {
				case 'C':
					return rune(KeyAltRight), b[6:]
				case 'D':
					return rune(KeyAltLeft), b[6:]
				}
			}
		case '2':
			if b[3] == '0' && b[5] == '~' {
				switch b[4] {
				case '0':
					return rune(KeyPasteStart), b[6:]
				case '1':
					return rune(KeyPasteEnd), b[6:]
				}
			}
		}
	}

	for _, c := range b[0:] {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '~' {
			return rune(KeyUnknown), b[:]
		}
	}

	return utf8.RuneError, b
}

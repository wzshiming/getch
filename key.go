package getch

const (
	KeyNUL           = 0x00 // ^@ Null character
	KeyCtrlA         = 0x01 // ^A Start of Header
	KeyCtrlB         = 0x02 // ^B Start of Text
	KeyCtrlC         = 0x03 // ^C End of Text
	KeyCtrlD         = 0x04 // ^D End of Transmission
	KeyCtrlE         = 0x05 // ^E Enquiry
	KeyCtrlF         = 0x06 // ^F Acknowledgment
	KeyCtrlG         = 0x07 // ^G \a Bell
	KeyCtrlH         = 0x08 // ^H \b Backspace
	KeyCtrlI         = 0x09 // ^I \t Horizontal Tab
	KeyCtrlJ         = 0x0A // ^J \n Line feed
	KeyCtrlK         = 0x0B // ^K \v Vertical Tab
	KeyCtrlL         = 0x0C // ^L \f Form feed
	KeyCtrlM         = 0x0D // ^M \r Carriage return
	KeyCtrlN         = 0x0E // ^N Shift Out
	KeyCtrlO         = 0x0F // ^O Shift In
	KeyCtrlP         = 0x10 // ^P Data Link Escape
	KeyCtrlQ         = 0x11 // ^Q Device Control 1 (oft. XON)
	KeyCtrlR         = 0x12 // ^R Device Control 2
	KeyCtrlS         = 0x13 // ^S Device Control 3 (oft. XOFF)
	KeyCtrlT         = 0x14 // ^T Device Control 4
	KeyCtrlU         = 0x15 // ^U Negative Acknowledgment
	KeyCtrlV         = 0x16 // ^V Synchronous Idle
	KeyCtrlW         = 0x17 // ^W End of Trans. Block
	KeyCtrlX         = 0x18 // ^X Cancel
	KeyCtrlY         = 0x19 // ^Y End of Medium
	KeyCtrlZ         = 0x1A // ^Z  Substitute
	KeyCtrlLBrack    = 0x1B // ^[ \e Escape
	KeyCtrlSQuote    = 0x1C // ^\ File Separator
	KeyCtrlRBrack    = 0x1D // ^] Group Separator
	KeyCtrlCaret     = 0x1E // ^^ Record Separator
	KeyCtrlUnderLine = 0x1F // ^_ Unit Separator
	KeySpace         = 0x20 // ' '
	KeyExclam        = 0x21 // '!'
	KeyDQuote        = 0x22 // '"'
	KeyHash          = 0x23 // '#'
	KeyDollar        = 0x24 // '$'
	KeyPercent       = 0x25 // '%'
	KeyAmpersand     = 0x26 // '&'
	KeySQuote        = 0x27 // '\''
	KeyLParen        = 0x28 // '('
	KeyRParen        = 0x29 // ')'
	KeyStar          = 0x2A // '*'
	KeyPlus          = 0x2B // '+'
	KeyComma         = 0x2C // ','
	KeyDash          = 0x2D // '-'
	KeyPeriod        = 0x2E // '.'
	KeySlash         = 0x2F // '/'
	Key0             = 0x30
	Key1             = 0x31
	Key2             = 0x32
	Key3             = 0x33
	Key4             = 0x34
	Key5             = 0x35
	Key6             = 0x36
	Key7             = 0x37
	Key8             = 0x38
	Key9             = 0x39
	KeyColon         = 0x3A // ':'
	KeySemicolon     = 0x3B // ';'
	KeyLessThan      = 0x3C // '<'
	KeyEqual         = 0x3D // '='
	KeyGreatThan     = 0x3E // '>'
	KeyQMark         = 0x3F // '?'
	KeyAt            = 0x40 // '@'
	KeyA             = 0x41
	KeyB             = 0x42
	KeyC             = 0x43
	KeyD             = 0x44
	KeyE             = 0x45
	KeyF             = 0x46
	KeyG             = 0x47
	KeyH             = 0x48
	KeyI             = 0x49
	KeyJ             = 0x4A
	KeyK             = 0x4B
	KeyL             = 0x4C
	KeyM             = 0x4D
	KeyN             = 0x4E
	KeyO             = 0x4F
	KeyP             = 0x50
	KeyQ             = 0x51
	KeyR             = 0x52
	KeyS             = 0x53
	KeyT             = 0x54
	KeyU             = 0x55
	KeyV             = 0x56
	KeyW             = 0x57
	KeyX             = 0x58
	KeyY             = 0x59
	KeyZ             = 0x5A
	KeyLBrack        = 0x5B // '['
	KeyBackSlash     = 0x5C // '\'
	KeyRBrack        = 0x5D // ']'
	KeyCaret         = 0x5E // '^'
	KeyUnderLine     = 0x5F // '_'
	KeyUnquote       = 0x60 // '`'
	Key_a            = 0x61
	Key_b            = 0x62
	Key_c            = 0x63
	Key_d            = 0x64
	Key_e            = 0x65
	Key_f            = 0x66
	Key_g            = 0x67
	Key_h            = 0x68
	Key_i            = 0x69
	Key_j            = 0x6A
	Key_k            = 0x6B
	Key_l            = 0x6C
	Key_m            = 0x6D
	Key_n            = 0x6E
	Key_o            = 0x6F
	Key_p            = 0x70
	Key_q            = 0x71
	Key_r            = 0x72
	Key_s            = 0x73
	Key_t            = 0x74
	Key_u            = 0x75
	Key_v            = 0x76
	Key_w            = 0x77
	Key_x            = 0x78
	Key_y            = 0x79
	Key_z            = 0x7A
	KeyLBrace        = 0x7B // '{'
	KeyBar           = 0x7C // '|'
	KeyRBrace        = 0x7D // '}'
	KeyTilde         = 0x7E // '~'
	KeyDel           = 0x7F // Delete

	KeyUnknown = 0xEF00 /* UTF-16 surrogate area */ + iota
	KeyUp
	KeyDown
	KeyLeft
	KeyRight
	KeyAltLeft
	KeyAltRight
	KeyHome
	KeyEnd
	KeyDeleteWord
	KeyDeleteLine
	KeyClearScreen
	KeyPasteStart
	KeyPasteEnd
)

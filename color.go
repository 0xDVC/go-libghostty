package libghostty

/*
#include <ghostty/vt.h>
*/
import "C"

// ColorRGB represents an RGB color value.
// C: GhosttyColorRgb
type ColorRGB struct {
	R uint8
	G uint8
	B uint8
}

// PaletteSize is the number of entries in a terminal color palette.
const PaletteSize = 256

// Palette is a 256-color palette.
type Palette [PaletteSize]ColorRGB

// Named color palette indices.
// C: GHOSTTY_COLOR_NAMED_*
const (
	ColorNamedBlack         = C.GHOSTTY_COLOR_NAMED_BLACK
	ColorNamedRed           = C.GHOSTTY_COLOR_NAMED_RED
	ColorNamedGreen         = C.GHOSTTY_COLOR_NAMED_GREEN
	ColorNamedYellow        = C.GHOSTTY_COLOR_NAMED_YELLOW
	ColorNamedBlue          = C.GHOSTTY_COLOR_NAMED_BLUE
	ColorNamedMagenta       = C.GHOSTTY_COLOR_NAMED_MAGENTA
	ColorNamedCyan          = C.GHOSTTY_COLOR_NAMED_CYAN
	ColorNamedWhite         = C.GHOSTTY_COLOR_NAMED_WHITE
	ColorNamedBrightBlack   = C.GHOSTTY_COLOR_NAMED_BRIGHT_BLACK
	ColorNamedBrightRed     = C.GHOSTTY_COLOR_NAMED_BRIGHT_RED
	ColorNamedBrightGreen   = C.GHOSTTY_COLOR_NAMED_BRIGHT_GREEN
	ColorNamedBrightYellow  = C.GHOSTTY_COLOR_NAMED_BRIGHT_YELLOW
	ColorNamedBrightBlue    = C.GHOSTTY_COLOR_NAMED_BRIGHT_BLUE
	ColorNamedBrightMagenta = C.GHOSTTY_COLOR_NAMED_BRIGHT_MAGENTA
	ColorNamedBrightCyan    = C.GHOSTTY_COLOR_NAMED_BRIGHT_CYAN
	ColorNamedBrightWhite   = C.GHOSTTY_COLOR_NAMED_BRIGHT_WHITE
)

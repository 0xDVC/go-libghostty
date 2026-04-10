package libghostty

// Render state data getters and setters wrapping
// ghostty_render_state_get() and ghostty_render_state_set().
// Functions are ordered alphabetically.

/*
#include <ghostty/vt.h>

// Helper to create a properly initialized GhosttyRenderStateColors (sized struct).
static inline GhosttyRenderStateColors init_render_state_colors() {
	GhosttyRenderStateColors c = GHOSTTY_INIT_SIZED(GhosttyRenderStateColors);
	return c;
}
*/
import "C"

import (
	"errors"
	"unsafe"
)

// Cols returns the viewport width in cells.
func (rs *RenderState) Cols() (uint16, error) {
	var v C.uint16_t
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_COLS, unsafe.Pointer(&v))); err != nil {
		return 0, err
	}
	return uint16(v), nil
}

// ColorBackground returns the default/current background color.
func (rs *RenderState) ColorBackground() (ColorRGB, error) {
	var v C.GhosttyColorRgb
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_COLOR_BACKGROUND, unsafe.Pointer(&v))); err != nil {
		return ColorRGB{}, err
	}
	return ColorRGB{R: uint8(v.r), G: uint8(v.g), B: uint8(v.b)}, nil
}

// ColorCursor returns the cursor color when explicitly set by terminal
// state. Returns nil (without error) when no explicit cursor color is set.
func (rs *RenderState) ColorCursor() (*ColorRGB, error) {
	// Check whether a cursor color is set first.
	var has C.bool
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_COLOR_CURSOR_HAS_VALUE, unsafe.Pointer(&has))); err != nil {
		return nil, err
	}
	if !bool(has) {
		return nil, nil
	}

	var v C.GhosttyColorRgb
	err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_COLOR_CURSOR, unsafe.Pointer(&v)))
	if err != nil {
		var ge *Error
		if errors.As(err, &ge) && ge.Result == ResultInvalidValue {
			return nil, nil
		}
		return nil, err
	}
	c := ColorRGB{R: uint8(v.r), G: uint8(v.g), B: uint8(v.b)}
	return &c, nil
}

// ColorForeground returns the default/current foreground color.
func (rs *RenderState) ColorForeground() (ColorRGB, error) {
	var v C.GhosttyColorRgb
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_COLOR_FOREGROUND, unsafe.Pointer(&v))); err != nil {
		return ColorRGB{}, err
	}
	return ColorRGB{R: uint8(v.r), G: uint8(v.g), B: uint8(v.b)}, nil
}

// ColorPalette returns the active 256-color palette.
func (rs *RenderState) ColorPalette() (*Palette, error) {
	var cp [PaletteSize]C.GhosttyColorRgb
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_COLOR_PALETTE, unsafe.Pointer(&cp[0]))); err != nil {
		return nil, err
	}
	var p Palette
	for i, c := range cp {
		p[i] = ColorRGB{R: uint8(c.r), G: uint8(c.g), B: uint8(c.b)}
	}
	return &p, nil
}

// Colors returns all color information from the render state in a
// single call using the sized-struct API.
func (rs *RenderState) Colors() (*RenderStateColors, error) {
	cc := C.init_render_state_colors()
	if err := resultError(C.ghostty_render_state_colors_get(rs.ptr, &cc)); err != nil {
		return nil, err
	}

	result := &RenderStateColors{
		Background:     ColorRGB{R: uint8(cc.background.r), G: uint8(cc.background.g), B: uint8(cc.background.b)},
		Foreground:     ColorRGB{R: uint8(cc.foreground.r), G: uint8(cc.foreground.g), B: uint8(cc.foreground.b)},
		Cursor:         ColorRGB{R: uint8(cc.cursor.r), G: uint8(cc.cursor.g), B: uint8(cc.cursor.b)},
		CursorHasValue: bool(cc.cursor_has_value),
	}

	for i, c := range cc.palette {
		result.Palette[i] = ColorRGB{R: uint8(c.r), G: uint8(c.g), B: uint8(c.b)}
	}
	return result, nil
}

// CursorBlinking reports whether the cursor should blink based on
// terminal modes.
func (rs *RenderState) CursorBlinking() (bool, error) {
	var v C.bool
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_CURSOR_BLINKING, unsafe.Pointer(&v))); err != nil {
		return false, err
	}
	return bool(v), nil
}

// CursorPasswordInput reports whether the cursor is at a password
// input field.
func (rs *RenderState) CursorPasswordInput() (bool, error) {
	var v C.bool
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_CURSOR_PASSWORD_INPUT, unsafe.Pointer(&v))); err != nil {
		return false, err
	}
	return bool(v), nil
}

// CursorVisible reports whether the cursor is visible based on
// terminal modes.
func (rs *RenderState) CursorVisible() (bool, error) {
	var v C.bool
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_CURSOR_VISIBLE, unsafe.Pointer(&v))); err != nil {
		return false, err
	}
	return bool(v), nil
}

// CursorVisualStyle returns the visual style of the cursor.
func (rs *RenderState) CursorVisualStyle() (CursorVisualStyle, error) {
	var v C.GhosttyRenderStateCursorVisualStyle
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_CURSOR_VISUAL_STYLE, unsafe.Pointer(&v))); err != nil {
		return 0, err
	}
	return CursorVisualStyle(v), nil
}

// CursorViewportHasValue reports whether the cursor is visible within
// the viewport. If false, the cursor viewport position values are
// undefined.
func (rs *RenderState) CursorViewportHasValue() (bool, error) {
	var v C.bool
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_CURSOR_VIEWPORT_HAS_VALUE, unsafe.Pointer(&v))); err != nil {
		return false, err
	}
	return bool(v), nil
}

// CursorViewportWideTail reports whether the cursor is on the tail
// of a wide character. Only valid when CursorViewportHasValue
// returns true.
func (rs *RenderState) CursorViewportWideTail() (bool, error) {
	var v C.bool
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_CURSOR_VIEWPORT_WIDE_TAIL, unsafe.Pointer(&v))); err != nil {
		return false, err
	}
	return bool(v), nil
}

// CursorViewportX returns the cursor viewport x position in cells.
// Only valid when CursorViewportHasValue returns true.
func (rs *RenderState) CursorViewportX() (uint16, error) {
	var v C.uint16_t
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_CURSOR_VIEWPORT_X, unsafe.Pointer(&v))); err != nil {
		return 0, err
	}
	return uint16(v), nil
}

// CursorViewportY returns the cursor viewport y position in cells.
// Only valid when CursorViewportHasValue returns true.
func (rs *RenderState) CursorViewportY() (uint16, error) {
	var v C.uint16_t
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_CURSOR_VIEWPORT_Y, unsafe.Pointer(&v))); err != nil {
		return 0, err
	}
	return uint16(v), nil
}

// Dirty returns the current dirty state.
func (rs *RenderState) Dirty() (RenderStateDirty, error) {
	var v C.GhosttyRenderStateDirty
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_DIRTY, unsafe.Pointer(&v))); err != nil {
		return 0, err
	}
	return RenderStateDirty(v), nil
}

// RowIterator populates a pre-allocated row iterator with row data
// from the render state. The iterator can then be advanced with Next
// and queried with getter methods.
//
// The iterator can be reused across multiple calls. Row data is only
// valid until the next call to Update.
func (rs *RenderState) RowIterator(ri *RenderStateRowIterator) error {
	return resultError(C.ghostty_render_state_get(
		rs.ptr,
		C.GHOSTTY_RENDER_STATE_DATA_ROW_ITERATOR,
		unsafe.Pointer(&ri.ptr),
	))
}

// Rows returns the viewport height in cells.
func (rs *RenderState) Rows() (uint16, error) {
	var v C.uint16_t
	if err := resultError(C.ghostty_render_state_get(rs.ptr, C.GHOSTTY_RENDER_STATE_DATA_ROWS, unsafe.Pointer(&v))); err != nil {
		return 0, err
	}
	return uint16(v), nil
}

// SetDirty sets the dirty state.
func (rs *RenderState) SetDirty(dirty RenderStateDirty) error {
	v := C.GhosttyRenderStateDirty(dirty)
	return resultError(C.ghostty_render_state_set(rs.ptr, C.GHOSTTY_RENDER_STATE_OPTION_DIRTY, unsafe.Pointer(&v)))
}

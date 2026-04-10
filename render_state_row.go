package libghostty

// Render-state row iterator wrapping the
// GhosttyRenderStateRowIterator C APIs.

/*
#include <ghostty/vt.h>
*/
import "C"

import "unsafe"

// RenderStateRowIterator iterates over rows in a render state.
// Create one with NewRenderStateRowIterator, populate it via
// RenderState.RowIterator, then advance with Next and read data
// with getter methods.
//
// Row data is only valid as long as the underlying render state
// is not updated. It is unsafe to use row data after calling
// RenderState.Update.
//
// C: GhosttyRenderStateRowIterator
type RenderStateRowIterator struct {
	ptr C.GhosttyRenderStateRowIterator
}

// NewRenderStateRowIterator creates a new row iterator instance.
// The iterator is empty until populated via RenderState.RowIterator.
func NewRenderStateRowIterator() (*RenderStateRowIterator, error) {
	var ptr C.GhosttyRenderStateRowIterator
	if err := resultError(C.ghostty_render_state_row_iterator_new(nil, &ptr)); err != nil {
		return nil, err
	}
	return &RenderStateRowIterator{ptr: ptr}, nil
}

// Close frees the underlying row iterator handle. After this call,
// the iterator must not be used.
func (ri *RenderStateRowIterator) Close() {
	C.ghostty_render_state_row_iterator_free(ri.ptr)
}

// Next advances the iterator to the next row. Returns true if the
// iterator moved successfully and row data is available. Returns
// false when there are no more rows.
func (ri *RenderStateRowIterator) Next() bool {
	return bool(C.ghostty_render_state_row_iterator_next(ri.ptr))
}

// Dirty reports whether the current row is dirty and requires a
// redraw.
func (ri *RenderStateRowIterator) Dirty() (bool, error) {
	var v C.bool
	if err := resultError(C.ghostty_render_state_row_get(ri.ptr, C.GHOSTTY_RENDER_STATE_ROW_DATA_DIRTY, unsafe.Pointer(&v))); err != nil {
		return false, err
	}
	return bool(v), nil
}

// SetDirty sets the dirty state for the current row.
func (ri *RenderStateRowIterator) SetDirty(dirty bool) error {
	v := C.bool(dirty)
	return resultError(C.ghostty_render_state_row_set(ri.ptr, C.GHOSTTY_RENDER_STATE_ROW_OPTION_DIRTY, unsafe.Pointer(&v)))
}

// Raw returns the raw Row value for the current iterator position.
// The returned Row can be used with the same getter methods as rows
// obtained from GridRef.
func (ri *RenderStateRowIterator) Raw() (*Row, error) {
	var v C.GhosttyRow
	if err := resultError(C.ghostty_render_state_row_get(ri.ptr, C.GHOSTTY_RENDER_STATE_ROW_DATA_RAW, unsafe.Pointer(&v))); err != nil {
		return nil, err
	}
	return &Row{c: v}, nil
}

// Cells populates a pre-allocated row cells instance with cell data
// for the current row. The cells instance can then be advanced with
// Next or positioned with Select.
//
// The cells instance can be reused across rows. Cell data is only
// valid until the next call to RenderState.Update.
func (ri *RenderStateRowIterator) Cells(rc *RenderStateRowCells) error {
	return resultError(C.ghostty_render_state_row_get(
		ri.ptr,
		C.GHOSTTY_RENDER_STATE_ROW_DATA_CELLS,
		unsafe.Pointer(&rc.ptr),
	))
}

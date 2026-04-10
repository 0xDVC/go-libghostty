package libghostty

import "testing"

func TestRenderStateRowCells(t *testing.T) {
	term, err := NewTerminal(WithSize(80, 24))
	if err != nil {
		t.Fatal(err)
	}
	defer term.Close()

	term.VTWrite([]byte("hello"))

	rs, err := NewRenderState()
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Close()

	if err := rs.Update(term); err != nil {
		t.Fatal(err)
	}

	ri, err := NewRenderStateRowIterator()
	if err != nil {
		t.Fatal(err)
	}
	defer ri.Close()

	if err := rs.RowIterator(ri); err != nil {
		t.Fatal(err)
	}

	// Advance to first row.
	if !ri.Next() {
		t.Fatal("expected at least one row")
	}

	rc, err := NewRenderStateRowCells()
	if err != nil {
		t.Fatal(err)
	}
	defer rc.Close()

	if err := ri.Cells(rc); err != nil {
		t.Fatal(err)
	}

	// Iterate cells and count them.
	count := 0
	for rc.Next() {
		count++
	}
	if count != 80 {
		t.Fatalf("expected 80 cells, got %d", count)
	}
}

func TestRenderStateRowCellsSelect(t *testing.T) {
	term, err := NewTerminal(WithSize(80, 24))
	if err != nil {
		t.Fatal(err)
	}
	defer term.Close()

	term.VTWrite([]byte("ABCDE"))

	rs, err := NewRenderState()
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Close()

	if err := rs.Update(term); err != nil {
		t.Fatal(err)
	}

	ri, err := NewRenderStateRowIterator()
	if err != nil {
		t.Fatal(err)
	}
	defer ri.Close()

	if err := rs.RowIterator(ri); err != nil {
		t.Fatal(err)
	}

	if !ri.Next() {
		t.Fatal("expected at least one row")
	}

	rc, err := NewRenderStateRowCells()
	if err != nil {
		t.Fatal(err)
	}
	defer rc.Close()

	if err := ri.Cells(rc); err != nil {
		t.Fatal(err)
	}

	// Select column 2 (should be 'C').
	if err := rc.Select(2); err != nil {
		t.Fatal(err)
	}

	graphemes, err := rc.Graphemes()
	if err != nil {
		t.Fatal(err)
	}
	if len(graphemes) != 1 || graphemes[0] != 'C' {
		t.Fatalf("expected ['C'], got %v", graphemes)
	}
}

func TestRenderStateRowCellsGraphemes(t *testing.T) {
	term, err := NewTerminal(WithSize(80, 24))
	if err != nil {
		t.Fatal(err)
	}
	defer term.Close()

	term.VTWrite([]byte("AB"))

	rs, err := NewRenderState()
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Close()

	if err := rs.Update(term); err != nil {
		t.Fatal(err)
	}

	ri, err := NewRenderStateRowIterator()
	if err != nil {
		t.Fatal(err)
	}
	defer ri.Close()

	if err := rs.RowIterator(ri); err != nil {
		t.Fatal(err)
	}

	if !ri.Next() {
		t.Fatal("expected at least one row")
	}

	rc, err := NewRenderStateRowCells()
	if err != nil {
		t.Fatal(err)
	}
	defer rc.Close()

	if err := ri.Cells(rc); err != nil {
		t.Fatal(err)
	}

	// First cell should be 'A'.
	if !rc.Next() {
		t.Fatal("expected at least one cell")
	}
	graphemes, err := rc.Graphemes()
	if err != nil {
		t.Fatal(err)
	}
	if len(graphemes) != 1 || graphemes[0] != 'A' {
		t.Fatalf("expected ['A'], got %v", graphemes)
	}

	// Second cell should be 'B'.
	if !rc.Next() {
		t.Fatal("expected second cell")
	}
	graphemes, err = rc.Graphemes()
	if err != nil {
		t.Fatal(err)
	}
	if len(graphemes) != 1 || graphemes[0] != 'B' {
		t.Fatalf("expected ['B'], got %v", graphemes)
	}

	// Empty cell should return nil.
	if !rc.Next() {
		t.Fatal("expected third cell")
	}
	graphemes, err = rc.Graphemes()
	if err != nil {
		t.Fatal(err)
	}
	if graphemes != nil {
		t.Fatalf("expected nil graphemes for empty cell, got %v", graphemes)
	}
}

func TestRenderStateRowCellsStyle(t *testing.T) {
	term, err := NewTerminal(WithSize(80, 24))
	if err != nil {
		t.Fatal(err)
	}
	defer term.Close()

	// Write bold text.
	term.VTWrite([]byte("\x1b[1mX"))

	rs, err := NewRenderState()
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Close()

	if err := rs.Update(term); err != nil {
		t.Fatal(err)
	}

	ri, err := NewRenderStateRowIterator()
	if err != nil {
		t.Fatal(err)
	}
	defer ri.Close()

	if err := rs.RowIterator(ri); err != nil {
		t.Fatal(err)
	}

	if !ri.Next() {
		t.Fatal("expected at least one row")
	}

	rc, err := NewRenderStateRowCells()
	if err != nil {
		t.Fatal(err)
	}
	defer rc.Close()

	if err := ri.Cells(rc); err != nil {
		t.Fatal(err)
	}

	if !rc.Next() {
		t.Fatal("expected at least one cell")
	}

	style, err := rc.Style()
	if err != nil {
		t.Fatal(err)
	}
	if !style.Bold() {
		t.Fatal("expected bold style")
	}
}

func TestRenderStateRowCellsColors(t *testing.T) {
	term, err := NewTerminal(WithSize(80, 24))
	if err != nil {
		t.Fatal(err)
	}
	defer term.Close()

	// Write text with explicit fg/bg colors (SGR 38;2;R;G;B and 48;2;R;G;B).
	term.VTWrite([]byte("\x1b[38;2;255;0;0;48;2;0;0;255mX"))

	rs, err := NewRenderState()
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Close()

	if err := rs.Update(term); err != nil {
		t.Fatal(err)
	}

	ri, err := NewRenderStateRowIterator()
	if err != nil {
		t.Fatal(err)
	}
	defer ri.Close()

	if err := rs.RowIterator(ri); err != nil {
		t.Fatal(err)
	}

	if !ri.Next() {
		t.Fatal("expected at least one row")
	}

	rc, err := NewRenderStateRowCells()
	if err != nil {
		t.Fatal(err)
	}
	defer rc.Close()

	if err := ri.Cells(rc); err != nil {
		t.Fatal(err)
	}

	if !rc.Next() {
		t.Fatal("expected at least one cell")
	}

	fg, err := rc.FgColor()
	if err != nil {
		t.Fatal(err)
	}
	if fg == nil {
		t.Fatal("expected non-nil fg color")
	}
	if *fg != (ColorRGB{R: 255, G: 0, B: 0}) {
		t.Fatalf("expected red fg, got %+v", *fg)
	}

	bg, err := rc.BgColor()
	if err != nil {
		t.Fatal(err)
	}
	if bg == nil {
		t.Fatal("expected non-nil bg color")
	}
	if *bg != (ColorRGB{R: 0, G: 0, B: 255}) {
		t.Fatalf("expected blue bg, got %+v", *bg)
	}

	// Empty cell should have nil colors (use default).
	if !rc.Next() {
		t.Fatal("expected second cell")
	}
	fg, err = rc.FgColor()
	if err != nil {
		t.Fatal(err)
	}
	if fg != nil {
		t.Fatalf("expected nil fg for unstyled cell, got %+v", *fg)
	}
}

package libghostty

import "testing"

func TestModeValueANSI(t *testing.T) {
	m := ModeInsert
	if m.Value() != 4 {
		t.Fatalf("expected value 4, got %d", m.Value())
	}
	if !m.ANSI() {
		t.Fatal("expected ANSI mode")
	}
}

func TestModeValueDEC(t *testing.T) {
	m := ModeCursorVisible
	if m.Value() != 25 {
		t.Fatalf("expected value 25, got %d", m.Value())
	}
	if m.ANSI() {
		t.Fatal("expected DEC private mode")
	}
}

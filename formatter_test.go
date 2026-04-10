package libghostty

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

// Verify interface satisfaction at compile time.
var _ io.WriterTo = (*Formatter)(nil)

func TestFormatterPlainText(t *testing.T) {
	term, err := NewTerminal(WithSize(80, 24))
	if err != nil {
		t.Fatal(err)
	}
	defer term.Close()

	term.VTWrite([]byte("hello world"))

	f, err := NewFormatter(term, WithFormatterFormat(FormatterFormatPlain))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	out, err := f.FormatString()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "hello world") {
		t.Fatalf("expected output to contain 'hello world', got %q", out)
	}
}

func TestFormatterVT(t *testing.T) {
	term, err := NewTerminal(WithSize(80, 24))
	if err != nil {
		t.Fatal(err)
	}
	defer term.Close()

	// Write red text.
	term.VTWrite([]byte("\x1b[31mred text\x1b[0m"))

	f, err := NewFormatter(term, WithFormatterFormat(FormatterFormatVT))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	out, err := f.FormatString()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "red text") {
		t.Fatalf("expected output to contain 'red text', got %q", out)
	}
	// VT format should contain escape sequences.
	if !strings.Contains(out, "\x1b[") {
		t.Fatalf("expected VT output to contain escape sequences, got %q", out)
	}
}

func TestFormatterHTML(t *testing.T) {
	term, err := NewTerminal(WithSize(80, 24))
	if err != nil {
		t.Fatal(err)
	}
	defer term.Close()

	term.VTWrite([]byte("\x1b[1mbold text\x1b[0m"))

	f, err := NewFormatter(term, WithFormatterFormat(FormatterFormatHTML))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	out, err := f.FormatString()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "bold text") {
		t.Fatalf("expected output to contain 'bold text', got %q", out)
	}
}

func TestFormatterFormat(t *testing.T) {
	term, err := NewTerminal(WithSize(80, 24))
	if err != nil {
		t.Fatal(err)
	}
	defer term.Close()

	term.VTWrite([]byte("bytes test"))

	f, err := NewFormatter(term, WithFormatterFormat(FormatterFormatPlain))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	out, err := f.Format()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), "bytes test") {
		t.Fatalf("expected output to contain 'bytes test', got %q", string(out))
	}
}

func TestFormatterReflectsCurrentState(t *testing.T) {
	term, err := NewTerminal(WithSize(80, 24))
	if err != nil {
		t.Fatal(err)
	}
	defer term.Close()

	f, err := NewFormatter(term, WithFormatterFormat(FormatterFormatPlain))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	// Format before writing anything.
	out1, err := f.FormatString()
	if err != nil {
		t.Fatal(err)
	}

	// Write some text and format again.
	term.VTWrite([]byte("after write"))
	out2, err := f.FormatString()
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(out2, "after write") {
		t.Fatalf("expected second format to contain 'after write', got %q", out2)
	}
	if strings.Contains(out1, "after write") {
		t.Fatal("first format should not contain text written afterward")
	}
}

func TestFormatterWriteTo(t *testing.T) {
	term, err := NewTerminal(WithSize(80, 24))
	if err != nil {
		t.Fatal(err)
	}
	defer term.Close()

	term.VTWrite([]byte("writeto test"))

	f, err := NewFormatter(term, WithFormatterFormat(FormatterFormatPlain))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	// Formatter satisfies io.WriterTo.
	var wt io.WriterTo = f
	var buf bytes.Buffer
	n, err := wt.WriteTo(&buf)
	if err != nil {
		t.Fatal(err)
	}
	if n != int64(buf.Len()) {
		t.Fatalf("WriteTo returned %d but buffer has %d bytes", n, buf.Len())
	}
	if !strings.Contains(buf.String(), "writeto test") {
		t.Fatalf("expected output to contain 'writeto test', got %q", buf.String())
	}
}

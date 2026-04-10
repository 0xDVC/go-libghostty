package libghostty

/*
#include <ghostty/vt.h>
*/
import "C"

import "fmt"

// Result represents a Ghostty result code.
//
// C: GhosttyResult
type Result int

const (
	ResultSuccess      Result = C.GHOSTTY_SUCCESS
	ResultOutOfMemory  Result = C.GHOSTTY_OUT_OF_MEMORY
	ResultInvalidValue Result = C.GHOSTTY_INVALID_VALUE
	ResultOutOfSpace   Result = C.GHOSTTY_OUT_OF_SPACE
	ResultNoValue      Result = C.GHOSTTY_NO_VALUE
)

// Error holds a non-success Ghostty result.
type Error struct {
	Result Result
}

func (e *Error) Error() string {
	switch e.Result {
	case ResultOutOfMemory:
		return "ghostty: out of memory"
	case ResultInvalidValue:
		return "ghostty: invalid value"
	case ResultOutOfSpace:
		return "ghostty: out of space"
	case ResultNoValue:
		return "ghostty: no value"
	default:
		return fmt.Sprintf("ghostty: result=%d", int(e.Result))
	}
}

// Convert a result code to an error, returning nil on success.
func resultError(result C.GhosttyResult) error {
	if result == C.GHOSTTY_SUCCESS {
		return nil
	}

	return &Error{Result: Result(result)}
}

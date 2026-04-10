package libghostty

// System-level configuration wrapping ghostty_sys_set().
// These are process-global settings that must be configured at startup.

/*
#include <ghostty/vt.h>
#include <ghostty/vt/sys.h>

// Forward declaration for the Go log trampoline so we can take its
// address on the C side. Uses compatible types (no const, int for enum)
// to match what cgo generates for the //export function.
extern void goSysLogTrampoline(
	void* userdata,
	int level,
	uint8_t* scope,
	size_t scope_len,
	uint8_t* message,
	size_t message_len);

// Helper to install the Go log trampoline via ghostty_sys_set.
// We need this because cgo cannot take the address of a Go-exported
// function directly as a C function pointer.
static inline GhosttyResult sys_set_log_go(void) {
	return ghostty_sys_set(GHOSTTY_SYS_OPT_LOG, (const void*)goSysLogTrampoline);
}

// Helper to install the built-in stderr log callback.
static inline GhosttyResult sys_set_log_stderr(void) {
	return ghostty_sys_set(GHOSTTY_SYS_OPT_LOG, (const void*)ghostty_sys_log_stderr);
}

// Helper to clear the log callback.
static inline GhosttyResult sys_clear_log(void) {
	return ghostty_sys_set(GHOSTTY_SYS_OPT_LOG, NULL);
}
*/
import "C"

import "unsafe"

// SysLogLevel represents the severity level of a log message from the
// library. Maps directly to the C enum values.
// C: GhosttySysLogLevel
type SysLogLevel int

const (
	// SysLogLevelError is the error log level.
	SysLogLevelError SysLogLevel = C.GHOSTTY_SYS_LOG_LEVEL_ERROR

	// SysLogLevelWarning is the warning log level.
	SysLogLevelWarning SysLogLevel = C.GHOSTTY_SYS_LOG_LEVEL_WARNING

	// SysLogLevelInfo is the info log level.
	SysLogLevelInfo SysLogLevel = C.GHOSTTY_SYS_LOG_LEVEL_INFO

	// SysLogLevelDebug is the debug log level.
	SysLogLevelDebug SysLogLevel = C.GHOSTTY_SYS_LOG_LEVEL_DEBUG
)

// String returns a human-readable name for the log level.
func (l SysLogLevel) String() string {
	switch l {
	case SysLogLevelError:
		return "error"
	case SysLogLevelWarning:
		return "warning"
	case SysLogLevelInfo:
		return "info"
	case SysLogLevelDebug:
		return "debug"
	default:
		return "unknown"
	}
}

// SysLogFn is the Go callback type for log messages from the library.
// The scope identifies the subsystem (e.g. "osc", "kitty"); it is
// empty for unscoped (default) log messages. The message and scope
// are only valid for the duration of the call.
// C: GhosttySysLogFn
type SysLogFn func(level SysLogLevel, scope string, message string)

// sysLogFn is the currently installed Go log callback.
var sysLogFn SysLogFn

// SysSetLog installs a Go callback that receives internal library log
// messages. Pass nil to clear the callback and discard log messages.
//
// Which log levels are emitted depends on the build mode of the
// library. Debug builds emit all levels; release builds emit info
// and above.
//
// This function is not safe for concurrent use. Callers must ensure
// that log configuration is not modified while log messages may be
// delivered (e.g. configure at startup before creating terminals).
func SysSetLog(fn SysLogFn) error {
	sysLogFn = fn
	if fn == nil {
		return resultError(C.sys_clear_log())
	}
	return resultError(C.sys_set_log_go())
}

// SysSetLogStderr installs the built-in stderr log callback provided
// by libghostty. Each message is formatted as "[level](scope): message\n"
// and written to stderr in a thread-safe manner.
//
// This function is not safe for concurrent use. See [SysSetLog].
func SysSetLogStderr() error {
	sysLogFn = nil
	return resultError(C.sys_set_log_stderr())
}

//export goSysLogTrampoline
func goSysLogTrampoline(
	_ unsafe.Pointer,
	level C.int,
	scopePtr *C.uint8_t,
	scopeLen C.size_t,
	messagePtr *C.uint8_t,
	messageLen C.size_t,
) {
	fn := sysLogFn
	if fn == nil {
		return
	}

	var scope string
	if scopeLen > 0 {
		scope = C.GoStringN((*C.char)(unsafe.Pointer(scopePtr)), C.int(scopeLen))
	}

	var message string
	if messageLen > 0 {
		message = C.GoStringN((*C.char)(unsafe.Pointer(messagePtr)), C.int(messageLen))
	}

	fn(SysLogLevel(level), scope, message)
}

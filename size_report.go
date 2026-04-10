package libghostty

/*
#include <ghostty/vt.h>
*/
import "C"

// SizeReportStyle determines the output format for a terminal size report.
// C: GhosttySizeReportStyle
type SizeReportStyle int

const (
	// SizeReportMode2048 is the in-band size report (mode 2048).
	SizeReportMode2048 SizeReportStyle = C.GHOSTTY_SIZE_REPORT_MODE_2048

	// SizeReportCSI14T is the XTWINOPS text area size in pixels.
	SizeReportCSI14T SizeReportStyle = C.GHOSTTY_SIZE_REPORT_CSI_14_T

	// SizeReportCSI16T is the XTWINOPS cell size in pixels.
	SizeReportCSI16T SizeReportStyle = C.GHOSTTY_SIZE_REPORT_CSI_16_T

	// SizeReportCSI18T is the XTWINOPS text area size in characters.
	SizeReportCSI18T SizeReportStyle = C.GHOSTTY_SIZE_REPORT_CSI_18_T
)

// SizeReportSize holds terminal geometry for XTWINOPS size queries.
// C: GhosttySizeReportSize
type SizeReportSize struct {
	// Rows is the terminal row count in cells.
	Rows uint16

	// Columns is the terminal column count in cells.
	Columns uint16

	// CellWidth is the width of a single terminal cell in pixels.
	CellWidth uint32

	// CellHeight is the height of a single terminal cell in pixels.
	CellHeight uint32
}

package libghostty

/*
#include <ghostty/vt.h>
*/
import "C"

// ColorScheme identifies the terminal color scheme (light or dark).
// C: GhosttyColorScheme
type ColorScheme int

const (
	// ColorSchemeLight indicates a light color scheme.
	ColorSchemeLight ColorScheme = C.GHOSTTY_COLOR_SCHEME_LIGHT

	// ColorSchemeDark indicates a dark color scheme.
	ColorSchemeDark ColorScheme = C.GHOSTTY_COLOR_SCHEME_DARK
)

// DA1 conformance levels (Pp parameter).
// C: GHOSTTY_DA_CONFORMANCE_*
const (
	DAConformanceVT100  = C.GHOSTTY_DA_CONFORMANCE_VT100
	DAConformanceVT101  = C.GHOSTTY_DA_CONFORMANCE_VT101
	DAConformanceVT102  = C.GHOSTTY_DA_CONFORMANCE_VT102
	DAConformanceVT125  = C.GHOSTTY_DA_CONFORMANCE_VT125
	DAConformanceVT131  = C.GHOSTTY_DA_CONFORMANCE_VT131
	DAConformanceVT132  = C.GHOSTTY_DA_CONFORMANCE_VT132
	DAConformanceVT220  = C.GHOSTTY_DA_CONFORMANCE_VT220
	DAConformanceVT240  = C.GHOSTTY_DA_CONFORMANCE_VT240
	DAConformanceVT320  = C.GHOSTTY_DA_CONFORMANCE_VT320
	DAConformanceVT340  = C.GHOSTTY_DA_CONFORMANCE_VT340
	DAConformanceVT420  = C.GHOSTTY_DA_CONFORMANCE_VT420
	DAConformanceVT510  = C.GHOSTTY_DA_CONFORMANCE_VT510
	DAConformanceVT520  = C.GHOSTTY_DA_CONFORMANCE_VT520
	DAConformanceVT525  = C.GHOSTTY_DA_CONFORMANCE_VT525
	DAConformanceLevel2 = C.GHOSTTY_DA_CONFORMANCE_LEVEL_2
	DAConformanceLevel3 = C.GHOSTTY_DA_CONFORMANCE_LEVEL_3
	DAConformanceLevel4 = C.GHOSTTY_DA_CONFORMANCE_LEVEL_4
	DAConformanceLevel5 = C.GHOSTTY_DA_CONFORMANCE_LEVEL_5
)

// DA1 feature codes (Ps parameters).
// C: GHOSTTY_DA_FEATURE_*
const (
	DAFeatureColumns132          = C.GHOSTTY_DA_FEATURE_COLUMNS_132
	DAFeaturePrinter             = C.GHOSTTY_DA_FEATURE_PRINTER
	DAFeatureReGIS               = C.GHOSTTY_DA_FEATURE_REGIS
	DAFeatureSixel               = C.GHOSTTY_DA_FEATURE_SIXEL
	DAFeatureSelectiveErase      = C.GHOSTTY_DA_FEATURE_SELECTIVE_ERASE
	DAFeatureUserDefinedKeys     = C.GHOSTTY_DA_FEATURE_USER_DEFINED_KEYS
	DAFeatureNationalReplacement = C.GHOSTTY_DA_FEATURE_NATIONAL_REPLACEMENT
	DAFeatureTechnicalCharacters = C.GHOSTTY_DA_FEATURE_TECHNICAL_CHARACTERS
	DAFeatureLocator             = C.GHOSTTY_DA_FEATURE_LOCATOR
	DAFeatureTerminalState       = C.GHOSTTY_DA_FEATURE_TERMINAL_STATE
	DAFeatureWindowing           = C.GHOSTTY_DA_FEATURE_WINDOWING
	DAFeatureHorizontalScrolling = C.GHOSTTY_DA_FEATURE_HORIZONTAL_SCROLLING
	DAFeatureANSIColor           = C.GHOSTTY_DA_FEATURE_ANSI_COLOR
	DAFeatureRectangularEditing  = C.GHOSTTY_DA_FEATURE_RECTANGULAR_EDITING
	DAFeatureANSITextLocator     = C.GHOSTTY_DA_FEATURE_ANSI_TEXT_LOCATOR
	DAFeatureClipboard           = C.GHOSTTY_DA_FEATURE_CLIPBOARD
)

// DA2 device type identifiers (Pp parameter).
// C: GHOSTTY_DA_DEVICE_TYPE_*
const (
	DADeviceTypeVT100 = C.GHOSTTY_DA_DEVICE_TYPE_VT100
	DADeviceTypeVT220 = C.GHOSTTY_DA_DEVICE_TYPE_VT220
	DADeviceTypeVT240 = C.GHOSTTY_DA_DEVICE_TYPE_VT240
	DADeviceTypeVT330 = C.GHOSTTY_DA_DEVICE_TYPE_VT330
	DADeviceTypeVT340 = C.GHOSTTY_DA_DEVICE_TYPE_VT340
	DADeviceTypeVT320 = C.GHOSTTY_DA_DEVICE_TYPE_VT320
	DADeviceTypeVT382 = C.GHOSTTY_DA_DEVICE_TYPE_VT382
	DADeviceTypeVT420 = C.GHOSTTY_DA_DEVICE_TYPE_VT420
	DADeviceTypeVT510 = C.GHOSTTY_DA_DEVICE_TYPE_VT510
	DADeviceTypeVT520 = C.GHOSTTY_DA_DEVICE_TYPE_VT520
	DADeviceTypeVT525 = C.GHOSTTY_DA_DEVICE_TYPE_VT525
)

// DeviceAttributes holds the response data for all three DA levels.
// The terminal fills whichever sub-struct matches the request type.
// C: GhosttyDeviceAttributes
type DeviceAttributes struct {
	// Primary is the DA1 response data (CSI c).
	Primary DeviceAttributesPrimary

	// Secondary is the DA2 response data (CSI > c).
	Secondary DeviceAttributesSecondary

	// Tertiary is the DA3 response data (CSI = c).
	Tertiary DeviceAttributesTertiary
}

// DeviceAttributesPrimary holds primary device attributes (DA1).
// C: GhosttyDeviceAttributesPrimary
type DeviceAttributesPrimary struct {
	// ConformanceLevel is the Pp parameter. E.g. 62 for VT220.
	ConformanceLevel uint16

	// Features contains the DA1 feature codes (Ps parameters).
	// Only the first NumFeatures entries are valid.
	Features [64]uint16

	// NumFeatures is the number of valid entries in Features.
	NumFeatures int
}

// DeviceAttributesSecondary holds secondary device attributes (DA2).
// C: GhosttyDeviceAttributesSecondary
type DeviceAttributesSecondary struct {
	// DeviceType is the terminal type identifier (Pp). E.g. 1 for VT220.
	DeviceType uint16

	// FirmwareVersion is the firmware/patch version number (Pv).
	FirmwareVersion uint16

	// ROMCartridge is the ROM cartridge registration number (Pc).
	// Always 0 for emulators.
	ROMCartridge uint16
}

// DeviceAttributesTertiary holds tertiary device attributes (DA3).
// C: GhosttyDeviceAttributesTertiary
type DeviceAttributesTertiary struct {
	// UnitID is encoded as 8 uppercase hex digits in the response.
	UnitID uint32
}

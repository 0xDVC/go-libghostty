package libghostty

import "testing"

func TestGetBuildInfo(t *testing.T) {
	info, err := GetBuildInfo()
	if err != nil {
		t.Fatal(err)
	}

	// Version string should be non-empty.
	if info.VersionString == "" {
		t.Fatal("expected non-empty version string")
	}

	// At least one version component should be non-zero.
	if info.VersionMajor == 0 && info.VersionMinor == 0 && info.VersionPatch == 0 {
		t.Fatal("expected at least one non-zero version component")
	}

	// Optimize mode should be a known value.
	switch info.Optimize {
	case OptimizeDebug, OptimizeReleaseSafe, OptimizeReleaseSmall, OptimizeReleaseFast:
	default:
		t.Fatalf("unexpected optimize mode: %d", info.Optimize)
	}
}

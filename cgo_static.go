//go:build !dynamic

package libghostty

// #cgo pkg-config: --static libghostty-vt-static
// #cgo CFLAGS: -DGHOSTTY_STATIC
import "C"

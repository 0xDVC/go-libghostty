BUILD_DIR := build

# FetchContent places the ghostty source here.
GHOSTTY_ZIG_OUT := $(CURDIR)/$(BUILD_DIR)/_deps/ghostty-src/zig-out
PKG_CONFIG_PATH := $(GHOSTTY_ZIG_OUT)/share/pkgconfig
DYLD_LIBRARY_PATH := $(GHOSTTY_ZIG_OUT)/lib
LD_LIBRARY_PATH := $(GHOSTTY_ZIG_OUT)/lib

# Stamp file to track whether the cmake build has run.
STAMP := $(BUILD_DIR)/.ghostty-built

.PHONY: build test clean

$(STAMP):
	cmake -B $(BUILD_DIR) -DCMAKE_BUILD_TYPE=Release
	cmake --build $(BUILD_DIR)
	@touch $(STAMP)

build: $(STAMP)
	PKG_CONFIG_PATH=$(PKG_CONFIG_PATH) go build ./...

test: $(STAMP)
	PKG_CONFIG_PATH=$(PKG_CONFIG_PATH) DYLD_LIBRARY_PATH=$(DYLD_LIBRARY_PATH) LD_LIBRARY_PATH=$(LD_LIBRARY_PATH) go test ./...

clean:
	rm -rf $(BUILD_DIR)

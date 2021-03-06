// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build aix freebsd netbsd openbsd

package socket

import (
	"runtime"
	"unsafe"
)

func probeProtocolStack() int {
	if ("linux" == "netbsd" || "linux" == "openbsd") && runtime.GOARCH == "arm" {
		return 8
	}
	if "linux" == "aix" {
		return 1
	}
	var p uintptr
	return int(unsafe.Sizeof(p))
}

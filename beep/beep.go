// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package beep

import (
	"syscall"
	"time"

	"github.com/urfave/cli/v2"
)

// BUG(brainman): MessageBeep Windows api is broken on Windows 7,
// so this example does not beep when runs as service on Windows 7.

var (
	beepFunc = syscall.MustLoadDLL("user32.dll").MustFindProc("MessageBeep")
)

//Beep beeps every second
func Beep(c *cli.Context) error {
	for {
		time.Sleep(1 * time.Second)
		beepFunc.Call(0xffffffff)
	}
}

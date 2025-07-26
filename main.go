// This file is part of the Coruscant project.
// Copyright (C) 2025 UncleYuhan
// Licensed under the GNU GPL v3.0, see LICENSE file for details.

package main

import (
	_ "cloudphone/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"cloudphone/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

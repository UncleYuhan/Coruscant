package main

import (
	_ "cloudphone/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"cloudphone/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

package main

import (
	_ "UEditorGoFrame/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"UEditorGoFrame/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

package main

import (
	_ "go-view-sql/internal/packed"

	_ "go-view-sql/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"go-view-sql/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

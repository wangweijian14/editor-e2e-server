package main

import (
	_ "wiki/internal/packed"

	_ "wiki/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"wiki/internal/cmd"

)

func main() {
	cmd.Main.Run(gctx.New())
}

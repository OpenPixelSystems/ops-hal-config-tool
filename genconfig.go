package main

import (
	"github.com/openpixelsystems/ops-hal-config-tool/internal/types"
	"github.com/openpixelsystems/ops-hal-config-tool/internal/cli"
)

func main() {
	var clk types.InputClk
	clk.Name = "test"

	cli.RunCLI()
}

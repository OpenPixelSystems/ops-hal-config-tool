package main

import (
	"log"

	ht "github.com/openpixelsystems/ops-hal-config-tool/internal/types"
)

func main() {
	var clk ht.InputClk
	clk.Name = "test"

	log.Print(clk)
}

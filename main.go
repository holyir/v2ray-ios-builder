package main

import (
	"log"
	"os"

	"github.com/XTLS/Xray-core/main/distro/all"
	"github.com/XTLS/Xray-core/main/commands/base"
)

func main() {
	if err := all.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	base.Exit()
}

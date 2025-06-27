package main

import (
	"log"
	"os"

	"github.com/xtls/xray-core/main/distro/all"
	_ "github.com/xtls/xray-core/main/json"
)

func main() {
	os.Args = append([]string{"v2ray"}, os.Args[1:]...)
	if err := all.Execute(); err != nil {
		log.Fatal(err)
	}
}

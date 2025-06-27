package main

import (
	"fmt"
	"os"

	"github.com/XTLS/Xray-core/core"
)

func main() {
	fmt.Println("🔧 Starting V2Ray (No-CGO, Pure-Go build)...")

	instance, err := core.New(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to start: %v\n", err)
		os.Exit(1)
	}

	if err := instance.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Start error: %v\n", err)
		os.Exit(1)
	}
}

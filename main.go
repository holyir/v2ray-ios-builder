
package main

import (
    "log"
    "v2ray.com/core"
    "v2ray.com/core/app/log/command"
    "v2ray.com/core/common/platform"
)

func main() {
    platform.Init()
    command.Run()
    err := core.Main(nil)
    if err != nil {
        log.Fatal(err)
    }
}

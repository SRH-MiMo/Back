package main

import (
	"Back/init/cmd"
	"flag"
)

var configPathFlag = flag.String("config", "config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}

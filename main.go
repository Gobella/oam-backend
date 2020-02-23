package main

import (
	"flag"
	"os"

	"github.com/astaxie/beego"
	"github.com/oam-apiserver/env"
	_ "github.com/oam-apiserver/routers"
)

var (
	help *bool = flag.Bool("h", false, " this is a web server")
)

func main() {
	err := env.InitEnv()
	if err != nil {
		os.Exit(1)
	}
	if *help {
		flag.PrintDefaults()
		return
	}
	beego.Run()
}

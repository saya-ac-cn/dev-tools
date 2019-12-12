
package main

import (
	"dev-tools/bootstrap"
	"dev-tools/web/middleware"
	"dev-tools/web/router"
	"flag"
	"log"
	"os"
	"strconv"
)

// 启动地址及端口
const DEFAULTPORT  = 8080

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Dev-Tools", "Saya")
	app.Bootstrap()
	app.Configure(middleware.Configure, router.Configure)
	return app
}

func main() {
	port := flag.Int("p", DEFAULTPORT, "Set The Http Port")
	flag.Parse()
	pwd,_ := os.Getwd()
	log.Printf("Listen On Port:%d pwd:%s\n", *port, pwd)

	app := newApp()
	app.Listen(":" + strconv.Itoa(*port))
}
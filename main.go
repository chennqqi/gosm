package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/chennqqi/gosm/checker"
	"github.com/chennqqi/gosm/models"
	"github.com/chennqqi/gosm/web"
)

var (
	Version = "1.0.0"
)

func main() {
	appName := filepath.Base(os.Args[0])

	var db, conf string
	var help bool
	flag.StringVar(&conf, "c", "config.yml", "set config path")
	flag.StringVar(&db, "d", "gosm.db", "set db path")
	flag.BoolVar(&help, "h", false, "show useage")
	flag.Parse()
	if help {
		fmt.Println(appName, "\t", Version)
		fmt.Println()
		fmt.Println("\t -c <config>         set config path, default ./config.yml")
		fmt.Println("\t -d <db>             set database, default ./gosm.db")
		fmt.Println("\t -h                  show this usage")
		return
	}
	fmt.Println(appName, "\t", Version)

	fixSIGTERM()
	models.CurrentConfig = models.ParseConfigFile(conf)
	models.Connect(db)
	models.LoadServices()

	go web.Start()
	checker.Start()
}

func fixSIGTERM() {
	// Workaround for SIGTERM not working when pinging
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()
}

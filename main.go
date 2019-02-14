package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/chennqqi/gosm/checker"
	"github.com/chennqqi/gosm/models"
	"github.com/chennqqi/gosm/web"
)

const (
	Version = "1.1"
)

var (
	checkChannel = make(chan *models.Service)
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(Version)
		return
	}
	fixSIGTERM()
	models.CurrentConfig = models.ParseConfigFile()
	models.Connect()
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

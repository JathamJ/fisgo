package main

import (
	"flag"
	"github.com/JathamJ/fisgo/log"
	"github.com/JathamJ/fisgo/service"
	"github.com/JathamJ/fisgo/utils"
	"github.com/thep0y/go-logger/basic"
)

const (
	ModePusher 		= "pusher"
	ModeServer 		= "server"
	DefaultPort 	= "8899"
)

var (
	configPath 		string
	mode 			string
	port 			string
)

func init() {
	runningPath, err := utils.GetWorkSpacePath()
	if err != nil {
		log.Fatalf("workspace wrong, err: %s", err.Error())
		return
	}
	flag.StringVar(&configPath, "c", runningPath + "/fis-conf.yaml", "pusher config file path default ./fis-conf.yaml")
	flag.StringVar(&mode, "m", ModePusher, "run mode default pusher, support: pusher, server")
	flag.StringVar(&port, "p", DefaultPort, "receiver port default :" + DefaultPort)
	flag.Parse()
}

func main() {
	// log print level
	log.InitLog(basic.InfoLevel)
	switch mode {
	case ModePusher:
		Push()
	case ModeServer:
		Server()
	default:
		log.Warnf("not support mode: %s", mode)
		return
	}
}

// Pusher tool
func Push() {
	pusher := service.InitPusherByFile(configPath)
	pusher.WatcherFiles()
	pusher.Watch()
}

// Server tool
func Server() {
	receiver := service.InitReceiver(port)
	receiver.ServerStart()
}

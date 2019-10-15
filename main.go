package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/PGo-Projects/output"
	"github.com/glinton/ping"
	"github.com/google/logger"
	"github.com/robfig/cron/v3"
)

const logPath = "internet_connectivity.log"

func pingServer() {
	_, err := ping.IPv4(context.Background(), "google.com")
	if err != nil {
		_, err = ping.IPv4(context.Background(), "blackboard.stonybrook.edu")
		if err != nil {
			logger.Errorln("Pinging google.com and blackboard.stonybrook.edu failed")
		} else {
			logger.Warningln("Pinging google.com failed, but pinging blackboard.stonybrook.edu succeeded")
		}
	} else {
		logger.Infoln("Pinging google.com succeeded")
	}
}

func main() {
	_, err := ping.IPv4(context.Background(), "google.com")
	if err != nil && strings.Contains(err.Error(), "socket: permission denied") {
		output.ErrorStringln("This program requires root access on Unix machines to send ICMP packets.")
		os.Exit(1)
	}

	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(err)
	}

	defer logFile.Close()
	defer logger.Init("internet_connectivity", false, false, logFile).Close()

	c := cron.New()
	c.AddFunc("* * * * *", pingServer)
	c.Start()

	output.Println("Pinging servers...  Please note that the first ping would start at the minute mark.", output.GREEN)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		_ = <-sigs
		done <- true
	}()

	<-done
}

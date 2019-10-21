package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/PGo-Projects/output"
	"github.com/glinton/ping"
	"github.com/google/logger"
	"github.com/robfig/cron/v3"
	"github.com/sbu-ces-unofficial/pinger/internal/pflags"
	"github.com/spf13/cobra"
)

const logPath = "internet_connectivity.log"

var (
	externalURLs = []string{"google.com"}
	internalURLs = []string{"blackboard.stonybrook.edu"}
)

var pingerCmd = &cobra.Command{
	Use:     "pinger",
	Short:   "Pinger is an utility used to collect more data concerning network outages",
	Long:    `A utility to help determine when network outage occurs and how much of the network is down.`,
	Version: "0.0.3a",
	Run:     pinger,
}

func pingServer() {
	var externalPingErr error = nil
	for _, url := range externalURLs {
		_, err := ping.IPv4(context.Background(), url)
		if err != nil {
			externalPingErr = err
			logger.Warningf("Pinging %s failed\n", url)
		} else {
			logger.Infof("Pinging %s succeeded\n", url)
		}
	}

	if externalPingErr != nil {
		for _, url := range internalURLs {
			_, err := ping.IPv4(context.Background(), url)
			if err != nil {
				logger.Errorf("Pinging %s failed\n", url)
			} else {
				logger.Warningf("Pinging %s succeeded\n", url)
			}
		}
	}
}

func pinger(cmd *cobra.Command, args []string) {
	_, err := ping.IPv4(context.Background(), "google.com")
	if err != nil && strings.Contains(err.Error(), "socket: permission denied") {
		output.ErrorStringln("This program requires root access on Unix machines to send ICMP packets.")
		os.Exit(1)
	}

	config, err := pflags.Parse("config.pflags")
	if err == nil {
		output.Successln("Using configuration specified in config.pflags!")
		externalURLs = config.Arrays["external_urls"]
		internalURLs = config.Arrays["internal_urls"]
	}

	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(err)
	}

	defer logFile.Close()
	defer logger.Init("internet_connectivity", true, false, logFile).Close()
	logger.SetFlags(log.Ldate | log.Lmicroseconds)

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

func Execute() {
	if err := pingerCmd.Execute(); err != nil {
		output.Errorln(err)
		os.Exit(1)
	}
}

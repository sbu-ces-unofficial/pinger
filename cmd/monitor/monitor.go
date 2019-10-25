package monitor

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/PGo-Projects/output"
	"github.com/PGo-Projects/pflags"
	"github.com/google/logger"
	"github.com/robfig/cron/v3"
	"github.com/sbu-ces-unofficial/pinger/internal/ping"
	"github.com/spf13/cobra"
)

const logPath = "internet_connectivity.log"

var (
	externalURLs = []string{"google.com"}
	internalURLs = []string{"blackboard.stonbyrook.edu"}
)

var Cmd = &cobra.Command{
	Use:   "monitor",
	Short: "Track when network outages occur and their severity",
	Long:  "A utility to help determine when network outage occurs and how much of the network is down.",
	Run:   monitor,
}

func monitor(cmd *cobra.Command, args []string) {
	if err := ping.Test(); err != nil {
		output.Errorln(err)
		os.Exit(1)
	}

	parseConfig()
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(err)
	}

	defer logFile.Close()
	defer logger.Init(logPath, true, false, logFile).Close()
	logger.SetFlags(log.Ldate | log.Lmicroseconds)

	c := cron.New()
	c.AddFunc("* * * * *", func() {
		ping.PingWithFallback(externalURLs, internalURLs)
	})
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

func parseConfig() {
	config, err := pflags.Parse("config.pflags", "monitor")
	if err == nil {
		output.Successln("Using configuration specified in config.pflags!")

		if urls, ok := config.Array.Get("external_urls"); ok {
			externalURLs = make([]string, 0)
			for _, url := range urls {
				externalURLs = append(externalURLs, url.(string))
			}
		}
		if urls, ok := config.Array.Get("internal_urls"); ok {
			internalURLs = make([]string, 0)
			for _, url := range urls {
				internalURLs = append(internalURLs, url.(string))
			}
		}
	}
}

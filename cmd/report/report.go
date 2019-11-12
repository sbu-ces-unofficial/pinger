package report

import (
	"log"
	"os"
	"time"

	"github.com/PGo-Projects/output"
	"github.com/PGo-Projects/pflags"
	"github.com/google/logger"
	"github.com/sbu-ces-unofficial/pinger/internal/ping"
	"github.com/spf13/cobra"
)

const logPath = "connectivity_report.txt"

var (
	externalURLs = []string{"google.com"}
	internalURLs = []string{"blackboard.stonbyrook.edu"}

	timeout = 1000 * time.Millisecond
)

var Cmd = &cobra.Command{
	Use:   "report",
	Short: "Generate a report of the network status",
	Long:  "Pings a list of urls and determine which ones are reachable and which ones are not.  Can be configured via config.pflags.",
	Run:   report,
}

func report(cmd *cobra.Command, args []string) {
	parseConfig()
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(err)
	}

	defer logFile.Close()
	defer logger.Init(logPath, true, false, logFile).Close()
	logger.SetFlags(log.Ldate | log.Lmicroseconds)

	ping.Ping(externalURLs, internalURLs, timeout)
}

func parseConfig() {
	config, err := pflags.Parse("config.pflags", "report")
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

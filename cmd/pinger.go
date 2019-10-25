package cmd

import (
	"os"

	"github.com/PGo-Projects/output"
	"github.com/sbu-ces-unofficial/pinger/cmd/monitor"
	"github.com/sbu-ces-unofficial/pinger/cmd/report"
	"github.com/spf13/cobra"
)

const logPath = "internet_connectivity.log"

var (
	externalURLs = []string{"google.com"}
	internalURLs = []string{"blackboard.stonybrook.edu"}
)

var pingerCmd = &cobra.Command{
	Use:     "pinger",
	Long:    "A utility for monitoring and collecting information to troubleshoot network outages.",
	Version: "0.0.3",
	Run:     pinger,
}

func init() {
	pingerCmd.AddCommand(monitor.Cmd)
	pingerCmd.AddCommand(report.Cmd)
}

func pinger(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
	}
}

func Execute() {
	if err := pingerCmd.Execute(); err != nil {
		output.Errorln(err)
		os.Exit(1)
	}
}

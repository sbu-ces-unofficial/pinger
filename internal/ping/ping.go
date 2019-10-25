package ping

import (
	"context"
	"errors"
	"strings"

	pping "github.com/glinton/ping"
	"github.com/google/logger"
)

func Test() error {
	_, err := pping.IPv4(context.Background(), "google.com")
	if err != nil && strings.Contains(err.Error(), "socket: permission denied") {
		return errors.New("This program requires root access on Unix machines to send ICMP packets.")
	}
	return nil
}

func Ping(externalURLs []string, internalURLs []string) {
	for _, url := range externalURLs {
		_, err := pping.IPv4(context.Background(), url)
		if err != nil {
			logger.Warningf("Pinging %s failed\n", url)
		} else {
			logger.Infof("Pinging %s succeeded\n", url)
		}
	}

	for _, url := range internalURLs {
		_, err := pping.IPv4(context.Background(), url)
		if err != nil {
			logger.Errorf("Pinging %s failed\n", url)
		} else {
			logger.Infof("Pinging %s succeeded\n", url)
		}
	}
}

func PingWithFallback(externalURLs []string, internalURLs []string) {
	externalPingErr := false
	for _, url := range externalURLs {
		_, err := pping.IPv4(context.Background(), url)
		if err != nil {
			externalPingErr = true
			logger.Warningf("Pinging %s failed\n", url)
		} else {
			logger.Infof("Pinging %s succeeded\n", url)
		}
	}

	if externalPingErr {
		for _, url := range internalURLs {
			_, err := pping.IPv4(context.Background(), url)
			if err != nil {
				logger.Errorf("Pinging %s failed\n", url)
			} else {
				logger.Infof("Pinging %s succeeded\n", url)
			}
		}
	}
}

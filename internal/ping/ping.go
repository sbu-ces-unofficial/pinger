package ping

import (
	"fmt"
	"net"
	neturl "net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/logger"
)

func Ping(externalURLs []string, internalURLs []string, timeout time.Duration) {
	pingURLs(externalURLs, timeout)
	pingURLs(internalURLs, timeout)
}

func PingWithFallback(externalURLs []string, internalURLs []string, timeout time.Duration) {
	externalPingErr := pingURLs(externalURLs, timeout)

	if externalPingErr != nil {
		pingURLs(internalURLs, timeout)
	}
}

func pingURLs(urls []string, timeout time.Duration) error {
	var pingErr error = nil
	for _, url := range urls {
		port, err := getPort(url)
		if err != nil {
			logger.Errorf("Pinging %s failed: %s\n", url, err.Error())
			continue
		}

		url = strings.Replace(url, "http://", "", 1)
		url = strings.Replace(url, "https://", "", 1)
		err = ping(url, port, timeout)
		if err != nil {
			logger.Errorf("Pinging %s failed\n", url)
			pingErr = err
		} else {
			logger.Infof("Pinging %s succeeded\n", url)
		}
	}
	return pingErr
}

func getPort(url string) (uint, error) {
	u, err := neturl.Parse(url)
	if err != nil {
		return 0, err
	}
	if u.Port() != "" {
		port, err := strconv.ParseUint(u.Port(), 10, 64)
		return uint(port), err
	}

	if strings.HasPrefix(url, "http") {
		return 80, nil
	} else if strings.HasPrefix(url, "https") {
		return 443, nil
	} else {
		return 0, fmt.Errorf("Can not detect the port from the given url: %s", url)
	}
}

func ping(url string, port uint, timeout time.Duration) error {
	fullURL := fmt.Sprintf("%s:%d", url, port)
	conn, err := net.DialTimeout("tcp", fullURL, timeout)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer conn.Close()
	return nil
}

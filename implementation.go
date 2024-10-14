package neuronos

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type commander struct{}

func NewCommander() Commander {
	return &commander{}
}

func (c *commander) Ping(host string) (PingResult, error) {
	if !strings.HasPrefix(host, "http") {
		host = fmt.Sprintf("http://%s", host)
	}

	url, err := url.Parse(host)
	if err != nil {
		return PingResult{}, err
	}

	start := time.Now()
	if _, err := http.Get(url.String()); err != nil {
		return PingResult{}, err
	}

	return PingResult{
		Successful: true,
		Time:       time.Since(start),
	}, nil
}

func (c *commander) GetSystemInfo() (SystemInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return SystemInfo{}, err
	}

	addrs, err := net.LookupAddr(hostname)
	if err != nil || len(addrs) < 1 {
		res, _ := http.Get("https://api.ipify.org")
		ip, _ := io.ReadAll(res.Body)
		addrs = []string{string(ip)}
	}

	return SystemInfo{
		Hostname:  hostname,
		IPAddress: addrs[0],
	}, nil
}

package neuronos

import (
	"errors"
	"net"
	"os"
)

type commander struct{}

func NewCommander() Commander {
	return &commander{}
}

func (c *commander) Ping(msg string) (PingResult, error) {
	return PingResult{}, nil
}

func (c *commander) GetSystemInfo() (SystemInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return SystemInfo{}, err
	}

	addrs, err := net.LookupAddr(hostname)
	if err != nil {
		return SystemInfo{}, err
	}

	if len(addrs) < 1 {
		return SystemInfo{}, errors.New("no valid ip address")
	}

	return SystemInfo{
		Hostname:  hostname,
		IPAddress: addrs[0],
	}, nil
}

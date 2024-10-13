package neuronos

import "os"

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

	// Get IP address (implement this)

	return SystemInfo{
		Hostname:  hostname,
		IPAddress: "implement me",
	}, nil
}

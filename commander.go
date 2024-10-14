package neuronos

import "time"

// Execution Data Models
type Commander interface {
	Ping(host string) (PingResult, error)
	GetSystemInfo() (SystemInfo, error)
}

type PingResult struct {
	Successful bool
	Time       time.Duration
}

type SystemInfo struct {
	Hostname  string
	IPAddress string
}

// Communication Data Models
type CommandRequest struct {
	Type    string `json:"type"`    // "ping" or "sysinfo"
	Payload string `json:"payload"` // For ping, this is the host
}

type CommandResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

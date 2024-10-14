package neuronos_test

import (
	"testing"

	"neuronos"
)

func TestGetSystemInfo(t *testing.T) {
	cmdr := neuronos.NewCommander()
	info, err := cmdr.GetSystemInfo()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if info.Hostname == "" {
		t.Error("Expected hostname to be non-empty")
	}

	if info.IPAddress == "" {
		t.Error("Expected IP address to be non-empty")
	}
}

func TestPing(t *testing.T) {
	cmdr := neuronos.NewCommander()

	testCases := []struct {
		Name string
		Host string
	}{
		{"Without HTTP prefix in hostname", "www.example.com"},
		{"With HTTP prefix in hostname", "http://www.example.com"},
	}

	for _, c := range testCases {
		t.Run(c.Name, func(t *testing.T) {
			res, err := cmdr.Ping(c.Host)
			if err != nil {
				t.Fatalf("Failed to ping with http prefix: %s", err)
			}

			if res.Successful == false {
				t.Fatal("Failed to ping with http prefix: no error")
			}
			
			if res.Time == 0 {
				t.Fatal("Ping succeeded with any latency")
			}
		})
	}
}

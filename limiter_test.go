package limiter

// file must end with _test
import (
	// must import "testing"
	"testing"
	"time"
)

// each function starts with "Test"
func TestTokenBucket (t *testing.T) {
	// rate 1.0 token per second, max 3 tokens 
	rl := NewLimiter(1.0, 3.0)	

	ip := "192.168.1.1"

	for i := range 3{
		// should not fail 
		if !rl.Allow(ip) {
			t.Errorf("Request %d was incorrectly dropped", i)
		}
	}

	// this should have passed
	if rl.Allow(ip) {
		t.Errorf("4th request should have been dropped")
	}

	// sleep 1.2 second so bucket fills up
	time.Sleep(1200 * time.Millisecond)

	// this should pass 
	if !rl.Allow(ip) {
		t.Errorf("Failed to refill tokens over time")
	}
}

package ratelimiter

import (
	"sync/atomic"
	"testing"
	"time"
)

const (
	MAX_TEST = 5
)

func TestRateLimiter(t *testing.T) {
	var nb uint32

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	startDate := time.Now()
	rl := NewRateLimiter(5, time.Second)

	go func() {
		for i := 0; i < 15; i++ {
			rl.CheckRate()
			dNow := time.Now()
			atomic.AddUint32(&nb, 1)
			currentNB := atomic.LoadUint32(&nb)
			if currentNB%5 == 0 {
				if dNow.Sub(startDate) < time.Duration(currentNB/5)*time.Second {
					t.Errorf("consume %d need at least %s, but it taks %s", currentNB, time.Duration(currentNB/5)*time.Second, dNow.Sub(startDate))
				}
			}
		}
	}()

	go func() {
		for i := 0; i < 15; i++ {
			rl.CheckRate()
			dNow := time.Now()
			atomic.AddUint32(&nb, 1)
			currentNB := atomic.LoadUint32(&nb)
			if currentNB%5 == 0 {
				if dNow.Sub(startDate) < time.Duration(currentNB/5)*time.Second {
					t.Errorf("consume %d need at least %s, but it taks %s", currentNB, time.Duration(currentNB/5)*time.Second, dNow.Sub(startDate))
				}
			}
		}
	}()

	time.Sleep(7 * time.Second)
}

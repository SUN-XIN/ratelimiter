package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	wait            chan bool
	waitBetweenEach time.Duration

	Max    int
	Window time.Duration

	stop  chan bool
	mutex *sync.Mutex
}

func NewRateLimiter(max int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		Max:    max,
		Window: window,

		wait:            make(chan bool),
		waitBetweenEach: window / time.Duration(max),

		stop:  make(chan bool, 1),
		mutex: &sync.Mutex{},
	}

	go func(r *RateLimiter) {
		ticker := time.NewTicker(r.waitBetweenEach)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				<-r.wait
			case <-rl.stop:
				break
			}
		}
	}(rl)

	return rl
}

func (rt *RateLimiter) Stop() {
	rt.stop <- true
}

func (rt *RateLimiter) CheckRate() {
	rt.wait <- true
}

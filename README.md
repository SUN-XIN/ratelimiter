# ratelimiter
simple rate limiter: max N times during DURATION. 

# example
```
rl := NewRateLimiter(5, time.Second)
for i := 0; i < 15; i++ {
  rl.CheckRate()
  // do something
}
```

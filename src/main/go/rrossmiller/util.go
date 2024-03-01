package main

import (
	"fmt"
	"time"
)

func meanTime(times []time.Duration) time.Duration {
	var ttl time.Duration
	for _, t := range times {
		ttl += t
	}

	avg := float64(ttl.Milliseconds()) / float64(len(times))
	dur, err := time.ParseDuration(fmt.Sprintf("%vms", avg))
	if err != nil {
		panic(err)
	}
	return dur
}

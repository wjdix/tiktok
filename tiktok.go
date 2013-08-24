package tiktok

import (
	"time"
)

type ControllableTicker struct {
	C        <-chan time.Time
	c        chan time.Time
	interval int
	elapsed  int
}

func NewTicker(d int) *ControllableTicker {
	c := make(chan time.Time)
	return &ControllableTicker{
		C:        c,
		interval: d,
		c:        c,
		elapsed:  0,
	}
}

func (ticker *ControllableTicker) Tick(d int) {
	ticker.elapsed = d + ticker.elapsed
	for ticker.elapsed >= ticker.interval {
		ticker.elapsed = ticker.elapsed - ticker.interval
		ticker.c <- time.Now()
	}
}

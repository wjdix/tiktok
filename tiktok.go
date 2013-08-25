package tiktok

import (
	"time"
)

var tickers []ControllableTicker

type ControllableTicker struct {
	C            <-chan time.Time
	c            chan time.Time
	interval     time.Duration
	elapsed      time.Duration
	elapsedChan  chan time.Duration
	shutdownChan chan int
	ticked       int
}

func NewTicker(d time.Duration) ControllableTicker {
	c := make(chan time.Time)
	elapsedChan := make(chan time.Duration)
	shutdownChan := make(chan int)
	ticker := ControllableTicker{
		C:            c,
		interval:     d,
		c:            c,
		elapsedChan:  elapsedChan,
		shutdownChan: shutdownChan,
	}

	tickers = append(tickers, ticker)
	go ticker.listen()
	return ticker
}

func (ticker *ControllableTicker) listen() {
	for {
		select {
		case elapsed := <-ticker.elapsedChan:
			ticker.elapsed += elapsed
			neededTicks := ticker.elapsed / ticker.interval
			for int(neededTicks) > ticker.ticked {
				ticker.ticked++
				ticker.c <- time.Now()
			}
		case <-ticker.shutdownChan:
			return
		}
	}
}

func (ticker *ControllableTicker) Tick(d time.Duration) {
	ticker.elapsedChan <- d
}

func (ticker *ControllableTicker) ShutDown() {
	ticker.shutdownChan <- 0
	removeTicker(ticker)
}

func (ticker ControllableTicker) Stop() {
	ticker.ShutDown()
}

func (ticker ControllableTicker) Chan() <-chan time.Time {
	return ticker.C
}

func removeTicker(ticker *ControllableTicker) {
	for i, present := range tickers {
		if *ticker == present {
			tickers = append(tickers[:i], tickers[i+1:]...)
			return
		}
	}
}

func Tick(d time.Duration) {
	for _, ticker := range tickers {
		ticker.Tick(d)
	}
}

func ClearTickers() {
	tickers = []ControllableTicker{}
}

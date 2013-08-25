package tiktok

import (
	"github.com/benmills/quiz"
	"testing"
	"time"
)

func countTicks(tickChan <-chan time.Time, shutDownChan chan int, returnChan chan int) {
	ticks := 0
	for {
		select {
		case <-tickChan:
			ticks += 1
		case <-shutDownChan:
			returnChan <- ticks
			return
		}
	}
}

func TestTickerDoesNotTickEarly(t *testing.T) {
	test := quiz.Test(t)
	shutDownChannel := make(chan int)
	finishedChannel := make(chan int)
	ticker := NewTicker(5)
	go countTicks(ticker.C, shutDownChannel, finishedChannel)

	ticker.Tick(3)
	var ticks int
	shutDownChannel <- 0
	select {
	case ticks = <-finishedChannel:
		test.Expect(ticks).ToEqual(0)
	}
}

func TestTickerTicksAfterDuration(t *testing.T) {
	test := quiz.Test(t)
	shutDownChannel := make(chan int)
	finishedChannel := make(chan int)
	ticker := NewTicker(5)
	go countTicks(ticker.C, shutDownChannel, finishedChannel)

	ticker.Tick(5)
	shutDownChannel <- 0
	select {
	case ticks := <-finishedChannel:
		test.Expect(ticks).ToEqual(1)
	}
}

func TestTickerTicksMultipleTimes(t *testing.T) {
	test := quiz.Test(t)
	shutDownChannel := make(chan int)
	finishedChannel := make(chan int)
	ticker := NewTicker(5)
	go countTicks(ticker.C, shutDownChannel, finishedChannel)

	ticker.Tick(6)
	ticker.Tick(4)

	shutDownChannel <- 0
	select {
	case ticks := <-finishedChannel:
		test.Expect(ticks).ToEqual(2)
	}
}

func TestTickerTicksMultipleTimeWithOneTickCall(t *testing.T) {
	test := quiz.Test(t)
	shutDownChannel := make(chan int)
	finishedChannel := make(chan int)
	ticker := NewTicker(3)
	go countTicks(ticker.C, shutDownChannel, finishedChannel)

	ticker.Tick(7)

	ticker.ShutDown()
	shutDownChannel <- 0
	select {
	case ticks := <-finishedChannel:
		test.Expect(ticks).ToBeGreaterThan(2)
	}
}

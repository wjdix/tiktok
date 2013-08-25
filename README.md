tiktok
======
[![Build Status](https://travis-ci.org/wjdix/tiktok.png?branch=master)](https://travis-ci.org/wjdix/tiktok)

Usage
=====
Tiktok provides a controllable replacement for Go's Tickers provided by the time package.
A new Ticker is created like so: 

``` go
  ticker := tiktok.NewTicker(5)
```

Tickers can either be controlled directly:

``` go
  ticker.Tick(5)
```

Or at the package level:

``` go
  tiktok.Tick(5)
```

Contributing
=====

First, run the tests.


``` shell
go get github.com/benmill/quiz

go test github.com/wjdix/tiktok
```

Write some new tests, fix em and send a pull request.

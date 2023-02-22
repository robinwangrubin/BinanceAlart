package main

import (
	"binanceAlart/consumer"
	"binanceAlart/spider"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go spider.StartBinanceWS(&wg)
	go consumer.StartBinanceWorkerPool(&wg)
	wg.Wait()
}

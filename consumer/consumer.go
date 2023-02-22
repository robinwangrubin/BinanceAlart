package consumer

import (
	"binanceAlart/alart"
	"binanceAlart/common"
	"binanceAlart/logs"
	"binanceAlart/spider"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/tidwall/gjson"
)

func NewBinanceData(symbol string, pricePercentDiff, quantity float64, timestamp time.Time, price string) *common.BinanceData {
	return &common.BinanceData{
		Symbol:           symbol,
		PricePercentDiff: pricePercentDiff,
		Quantity:         quantity,
		Timestamp:        timestamp,
		Price:            price,
	}
}

func StartBinanceWorkerPool(wg *sync.WaitGroup) {
	for w := 1; w <= 10; w++ {
		wg.Add(1)
		go binanceWorker(wg, w, spider.BinanceJobChannel)
	}
}

func binanceWorker(wg *sync.WaitGroup, id int, BinanceJobChannel chan *string) {
	logs.ConsoleInfo("Binance worker %d have already started.", id)
	defer wg.Done()
	for {
		messages := <-BinanceJobChannel
		kline := gjson.Get(string(*messages), "data.k").Map()
		// 判断K线是否结束；只拿最终的状态
		if !kline["x"].Bool() {
			continue
		}

		symbol := kline["s"].String()
		// token := symbol[:len(symbol)-4]
		timestamp := time.Unix(0, kline["t"].Int()*1e6)
		openPrice, _ := strconv.ParseFloat(kline["o"].String(), 64)
		closePrice, _ := strconv.ParseFloat(kline["c"].String(), 64)
		quantity, _ := strconv.ParseFloat(kline["Q"].String(), 64)
		price := kline["c"].String()
		pricePercentDiff, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", ((closePrice-openPrice)/openPrice)*100), 64)

		point := NewBinanceData(symbol, pricePercentDiff, quantity, timestamp, price)

		alart.Judge(point)

		WriteBinanceDataToDB("Binance", point)
		logs.ConsoleInfo("%s %.2f %.2f", point.Symbol, point.PricePercentDiff, point.Quantity)
	}
}

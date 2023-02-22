package common

import "time"

type BinanceData struct {
	Symbol           string    // 交易对
	PricePercentDiff float64   // 价格涨跌幅
	Quantity         float64   // 成交量
	Timestamp        time.Time // 交易时间
	Price            string
}

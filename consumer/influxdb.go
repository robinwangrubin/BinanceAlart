package consumer

import (
	"binanceAlart/common"
	"context"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var (
	bucket   = "BlockChain"
	org      = ""
	token    = ""
	url      = "http://127.0.0.1:8086/"
	client   = influxdb2.NewClient(url, token)
	writeAPI = client.WriteAPIBlocking(org, bucket)
)

func WriteBinanceDataToDB(measurement string, point *common.BinanceData) {
	p := influxdb2.NewPoint(measurement,
		map[string]string{
			"Symbol": point.Symbol,
		},
		map[string]interface{}{
			"Quantity":         point.Quantity,
			"PricePercentDiff": point.PricePercentDiff,
		},
		point.Timestamp)
	writeAPI.WritePoint(context.Background(), p)
	// client.Close()
}

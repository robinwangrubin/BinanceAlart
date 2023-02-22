package spider

import (
	"binanceAlart/logs"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

var (
	BinanceJobChannel = make(chan *string, 100000)
)

func getBinanceStreams() (streams string) {
	streams = "/stream?streams="
	url := "https://api.binance.com/api/v3/ticker/price"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		logs.ConsoleFatal("error: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		logs.ConsoleFatal("error: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.ConsoleFatal("error: %v", err)
	}
	symbols := gjson.Get(string(body), "#.symbol").Array()
	for _, symbol := range symbols {
		if symbol.String()[len(symbol.String())-4:] == "USDT" && !strings.Contains(symbol.String(), "UP") && !strings.Contains(symbol.String(), "DOWN") {
			streams = streams + strings.ToLower(symbol.String()) + "@kline_5m" + "/"
		}

	}
	streams = streams[:len(streams)-1]
	return
}

// var path string = "/stream?streams=btcusdt@kline_5m/ethusdt@kline_5m/dydxusdt@kline_5m/aptusdt@kline_5m/highusdt@kline_5m/ftmusdt@kline_5m/maskusdt@kline_5m/hookusdt@kline_5m/bnbusdt@kline_5m/maticusdt@kline_5m/sfpusdt@kline_5m/opusdt@kline_5m/rndrusdt@kline_5m/galausdt@kline_5m/tusdt@kline_5m/nearusdt@kline_5m/injusdt@kline_5m/minausdt@kline_5m/ldousdt@kline_5m/arusdt@kline_5m/peopleusdt@kline_5m/achusdt@kline_5m/lrcusdt@kline_5m/linkusdt@kline_5m/sandusdt@kline_5m/lokausdt@kline_5m/fetusdt@kline_5m/apeusdt@kline_5m/galusdt@kline_5m/imxusdt@kline_5m/bandusdt@kline_5m/crvusdt@kline_5m/audiousdt@kline_5m/perpusdt@kline_5m/ensusdt@kline_5m/klayusdt@kline_5m/hftusdt@kline_5m/dashusdt@kline_5m/kavausdt@kline_5m/gmxusdt@kline_5m/algousdt@kline_5m/uniusdt@kline_5m/gtcusdt@kline_5m/duskusdt@kline_5m/kncusdt@kline_5m/astrusdt@kline_5m/voxelusdt@kline_5m/aaveusdt@kline_5m/sushiusdt@kline_5m/fluxusdt@kline_5m/tvkusdt@kline_5m/stgusdt@kline_5m/atausdt@kline_5m/oceanusdt@kline_5m/runeusdt@kline_5m/aliceusdt@kline_5m/vetusdt@kline_5m/enjusdt@kline_5m/fxsusdt@kline_5m/thetausdt@kline_5m/cakeusdt@kline_5m/yggusdt@kline_5m/litusdt@kline_5m/lptusdt@kline_5m/cocosusdt@kline_5m/betausdt@kline_5m/celrusdt@kline_5m/celousdt@kline_5m/snxusdt@kline_5m/zecusdt@kline_5m/mdtusdt@kline_5m/zenusdt@kline_5m/ankrusdt@kline_5m/chrusdt@kline_5m/rlcusdt@kline_5m/rvnusdt@kline_5m/ctxcusdt@kline_5m/woousdt@kline_5m/xecusdt@kline_5m/burgerusdt@kline_5m/cvxusdt@kline_5m/rsrusdt@kline_5m/bakeusdt@kline_5m/batusdt@kline_5m/neousdt@kline_5m/ctkusdt@kline_5m/sxpusdt@kline_5m/radusdt@kline_5m/oxtusdt@kline_5m/pondusdt@kline_5m/iotausdt@kline_5m/compusdt@kline_5m/wingusdt@kline_5m/ghstusdt@kline_5m/linausdt@kline_5m/ctsiusdt@kline_5m/hifiusdt@kline_5m/cityusdt@kline_5m/mkrusdt@kline_5m/storjusdt@kline_5m/dentusdt@kline_5m/bondusdt@kline_5m/rayusdt@kline_5m/dodousdt@kline_5m/sysusdt@kline_5m/portousdt@kline_5m/straxusdt@kline_5m/renusdt@kline_5m/bttcusdt@kline_5m/balusdt@kline_5m/ookiusdt@kline_5m/autousdt@kline_5m/forusdt@kline_5m/funusdt@kline_5m/magicusdt@kline_5m/usdt@kline_5m/osmousdt@kline_5m/rplusdt@kline_5m/leverusdt@kline_5m"

func StartBinanceWS(wg *sync.WaitGroup) {
	defer wg.Done()
	path := getBinanceStreams()
	wsUrl := "wss://stream.binance.com:9443" + path
	conn, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
	if err != nil {
		logs.ConsoleFatal("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			logs.ConsoleFatal("Error in receive:", err)
			return
		}

		message := string(msg)
		BinanceJobChannel <- &message
	}
}

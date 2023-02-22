package common

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gopkg.in/ini.v1"
)

var (
	BotKey     string
	MailServer string
	MailPort   int
	MaillUser  string
	MailPass   string
	ToUserList []string
)

func init() {
	cfg, err := ini.Load("./conf/app.conf")
	if err != nil {
		fmt.Printf("Load config failed, err: %v\n", err)
		os.Exit(1)
	}

	BotKey = cfg.Section("Bot").Key("key").String()
	MailServer = cfg.Section("Email").Key("mailServer").String()
	MailPort = cfg.Section("Email").Key("mailPort").MustInt()
	MaillUser = cfg.Section("Email").Key("maillUser").String()
	MailPass = cfg.Section("Email").Key("mailPass").String()
	ToUserList = strings.Split(cfg.Section("Email").Key("userList").String(), ",")

}

type BinanceData struct {
	Symbol           string    // 交易对
	PricePercentDiff float64   // 价格涨跌幅
	Quantity         float64   // 成交量
	Timestamp        time.Time // 交易时间
	Price            string
}

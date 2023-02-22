package alart

import (
	"binanceAlart/bot"
	"binanceAlart/common"
	"binanceAlart/logs"
	"crypto/tls"
	"fmt"
	"strings"

	"gopkg.in/gomail.v2"
)

func init() {
	logs.ConsoleInfo("Alart start successfully.")
	bot.SendMessageToWeChat("你的机器人已启动")
}

func Judge(point *common.BinanceData) {
	// 价格涨幅超过2；交易额大于10万美金；可考虑加入交易额对比；实现放量拉升监控
	if point.PricePercentDiff > 2 && point.Quantity > 200000 {
		if error := SendMailToCustomer(point.Symbol); error != nil {
			logs.ConsoleError("SendMailToCustomer Failed. error: %v", error)
		}
		// token := symbol[:len(symbol)-4]
		content := fmt.Sprintf("----Binance Alart----\nToken: %s \nPrice: %s\nPriceChange: +%.2f%%\nQuantity: %d\nTime: %v", point.Symbol[:len(point.Symbol)-4], strings.TrimRight(point.Price, "0"), point.PricePercentDiff, int(point.Quantity), point.Timestamp.Format("2006/01/02 15:04"))

		bot.SendMessageToWeChat(content)

		logs.ConsoleInfo("Token %s firing,Price Changed: %.2f,quantity:%.2f", point.Symbol[:len(point.Symbol)-4], point.PricePercentDiff, point.Quantity)
	}
}

func SendMailToCustomer(token string) error {
	mailServer := common.MailServer
	mailPort := common.MailPort
	maillUser := common.MaillUser
	mailPass := common.MailPass
	userList := common.ToUserList

	m := gomail.NewMessage()
	nickname := "Binance Alart"
	m.SetHeader("From", nickname+"<"+maillUser+">")
	m.SetHeader("To", userList...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	subject := fmt.Sprintf("%s 交易额和价格涨幅突增", token)
	m.SetHeader("Subject", subject)
	body := fmt.Sprintf("%s 交易额和价格涨幅突增", token)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(mailServer, mailPort, maillUser, mailPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

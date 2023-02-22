package bot

import (
	"binanceAlart/common"
	"binanceAlart/logs"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
{"msgtype":"text","text":{"content": "Hello guys"}}
*/

func SendMessageToWeChat(content string) {
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + common.BotKey
	method := "POST"

	requestBody := fmt.Sprintf(`{
		"msgtype":"text",
		"text":{
			"content": "%s"
		}
	}`, content)

	payload := strings.NewReader(requestBody)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

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
	if strings.Contains(string(body), "ok") {
		logs.ConsoleInfo("Send Message Sucessfully.")
	} else {
		logs.ConsoleInfo("Send Message Failed, error:%v", string(body))
	}
}

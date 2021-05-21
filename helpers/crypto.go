package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Duvewo/cryptobot/internal/cryptocurrency"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/url"
	"strings"
)

//Generating Forex trading API URL to websocket
//Scheme: https://stream{server_id}.forexpros.com/echo/info
//WebSocket Scheme: wss://stream{server_id}.forexpros.com/echo/{rand 3 len int}/{random 8 len str}/websocket
func generateURL() (*url.URL, error) {

	return url.Parse(fmt.Sprintf("wss://stream%d.forexpros.com/echo/%d/%s/websocket", rand.Intn(352)+1, rand.Intn(999)+100, randomString(8)))

}

//First response of Forex's WebSockets is \"o\"
func firstResponse(client *cryptocurrency.Client) error {
	mT, data, err := client.Read()

	if err != nil {
		return err
	}

	if mT == websocket.TextMessage && bytes.Equal(data, []byte("o")) {
		return nil
	}

	return fmt.Errorf("helpers/crypto: not \"o\"")

}

func Dial() *cryptocurrency.Client {
	for {
		u, err := generateURL()

		if err != nil {
			continue
		}

		cc, err := cryptocurrency.Dial(u)

		if err != nil {
			continue
		}

		err = firstResponse(cc)

		if err != nil {
			continue
		}

		return cc
	}
}

//TODO: wrap errors
func ForexInit(client *cryptocurrency.Client) error {

	err := client.WriteJSON(json.RawMessage(`["{\"_event\":\"bulk-subscribe\",\"tzID\":18,\"message\":\"pid-eu-1058142:%%pid-eu-8833:%%pid-eu-8849:%%pid-eu-8830:%%pid-eu-8836:%%pid-eu-8910:%%pid-eu-8883:%%pid-eu-8862:%%pid-eu-252:%%pid-eu-6408:%%pid-eu-6369:%%pid-eu-7888:%%pid-eu-284:%%pid-eu-9251:%%pid-eu-396:%%pid-eu-1010776:%%pid-eu-945629:%%pidTechSumm-252:%%pidTechSumm-6408:%%pidTechSumm-6369:%%pidTechSumm-7888:%%pidTechSumm-284:%%pidTechSumm-9251:%%pidTechSumm-396:%%pidExt-eu-1058142:%%isOpenExch-1053:%%isOpenPair-8833:%%isOpenPair-8849:%%isOpenPair-8830:%%isOpenPair-8836:%%isOpenPair-8910:%%isOpenPair-8883:%%isOpenPair-8862:%%cmt-7-5-1058142:%%domain-7:\"}"]`))
	if err != nil {
		return err
	}

	//err = client.WriteJSON([]byte(`["{\"_event\":\"UID\",\"UID\":0}"]`))
	//
	//if err != nil {
	//	return err
	//}

	return nil

}

func randomString(n int) string {
	const s = "abcdefghijklmnopqrstuvwxyz0123456789_"

	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(string(s[rand.Intn(len(s))]))
	}

	return sb.String()

}

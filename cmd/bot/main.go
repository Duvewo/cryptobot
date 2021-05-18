package main

import (
	"github.com/Duvewo/cryptobot/handlers"
	"github.com/Duvewo/cryptobot/helpers"
	"github.com/Duvewo/cryptobot/internal/cryptocurrency"
	"github.com/Duvewo/cryptobot/storage"
	tele "gopkg.in/tucnak/telebot.v3"
	"gopkg.in/tucnak/telebot.v3/layout"
	"log"
	"os"
	"time"
)

var CryptoCurrencyMap = make(map[string]float64, 10)

func main() {
	lt, err := layout.New("bot.yaml")

	if err != nil {
		log.Fatalf("failed to create layout: %v", err)
	}

	b, err := tele.NewBot(tele.Settings{
		Token: lt.Settings().Token,
		Poller: &tele.LongPoller{
			Timeout: time.Second * 60,
		},
	})

	if err != nil {
		log.Fatalf("failed to create bot: %v", err)
	}

	s, err := storage.Open(os.Getenv("PG_URL"))

	if err != nil {
		log.Fatalf("failed to open storage: %v", err)
	}

	//TODO: implement investing.com api

	cc, err := helpers.ForexDial()

	if err != nil {
		log.Fatalf("failed to dial Forex: %v", err)
	}

	h := handlers.New(handlers.Handler{
		Bot:            b,
		Storage:        s,
		Layout:         lt,
		CryptoCurrency: cc,
	})

	go func(client *cryptocurrency.Client) {
		err = helpers.ForexInit(cc)

		if err != nil {
			log.Fatalf("failed to initialize Forex: %v", err)
		}

		for {

			var resp cryptocurrency.Response
			err = client.ReadJSON(&resp)

			if err != nil {
				log.Fatalf("failed to read response: %v", err)
			}

			//TODO: get response to Message and save pairs to global map

			time.Sleep(time.Second * 10)

			err = client.WriteJSON(cryptocurrency.Request{Event: "heartbeat", Data: "h"})

			if err != nil {
				log.Fatalf("failed to send heartbeat: %v", err)
			}

		}
	}(cc)

	b.Use(lt.Middleware("ru"))

	b.Handle("/start", h.OnStart)
	b.Handle("/subscribe", h.OnSubscribe)

	b.Handle(tele.OnText, h.OnText)

	b.Start()

}

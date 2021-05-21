package main

import (
	"github.com/Duvewo/cryptobot/handlers"
	"github.com/Duvewo/cryptobot/helpers"
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

	cc := helpers.Dial()

	h := handlers.New(handlers.Handler{
		Bot:            b,
		Storage:        s,
		Layout:         lt,
		CryptoCurrency: cc,
	})

	err = helpers.ForexInit(cc)

	if err != nil {
		log.Fatalf("failed to initialize Forex: %v", err)
	}

	for {

		_, data, err := cc.Read()

		if err != nil {
			log.Printf("failed to read response: %v", err)
		}

		log.Printf("%#v", data)

	}

	b.Use(lt.Middleware("ru"))

	b.Handle("/start", h.OnStart)
	b.Handle("/subscribe", h.OnSubscribe)

	b.Handle(tele.OnText, h.OnText)

	b.Start()

}

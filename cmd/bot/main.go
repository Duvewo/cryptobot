package main

import (
	"github.com/Duvewo/cryptobot/handlers"
	"github.com/Duvewo/cryptobot/internal/cryptocurrency"
	"github.com/Duvewo/cryptobot/storage"
	tele "gopkg.in/tucnak/telebot.v3"
	"gopkg.in/tucnak/telebot.v3/layout"
	"log"
	"net/url"
	"os"
	"time"
)

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
	u, err := url.Parse("wss://stream224.forexpros.com/echo/820/5105ohnm/websocket")

	if err != nil {
		log.Fatalf("failed to parse URL: %v", err)
	}

	cc, err := cryptocurrency.Dial(u)

	if err != nil {
		log.Fatalf("failed to dial cryptocurrency: %v", err)
	}

	h := handlers.New(handlers.Handler{
		Bot:            b,
		Storage:        s,
		Layout:         lt,
		CryptoCurrency: cc,
	})

	b.Use(lt.Middleware("ru"))

	b.Handle("/start", h.OnStart)
	b.Handle("/subscribe", h.OnSubscribe)

	b.Handle(tele.OnText, h.OnText)

	b.Start()

}

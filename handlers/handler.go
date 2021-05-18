package handlers

import (
	"github.com/Duvewo/cryptobot/internal/cryptocurrency"
	"github.com/Duvewo/cryptobot/storage"
	tele "gopkg.in/tucnak/telebot.v3"
	"gopkg.in/tucnak/telebot.v3/layout"
)

type handler struct {
	b  *tele.Bot
	s  *storage.Storage
	lt *layout.Layout
	cc *cryptocurrency.Client
}

type Handler struct {
	Bot            *tele.Bot
	Storage        *storage.Storage
	Layout         *layout.Layout
	CryptoCurrency *cryptocurrency.Client
}

func New(h Handler) *handler {
	return &handler{b: h.Bot, s: h.Storage, lt: h.Layout, cc: h.CryptoCurrency}
}

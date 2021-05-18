package handlers

import (
	"github.com/Duvewo/cryptobot/storage"
	tele "gopkg.in/tucnak/telebot.v3"
)

func (h *handler) OnSubscribe(ctx tele.Context) error {
	if err := h.s.Users.Update(storage.User{ID: int64(ctx.Sender().ID), State: SubscribeState}); err != nil {
		return err
	}
	return nil
}

package handlers

import (
	"github.com/Duvewo/cryptobot/storage"
	tele "gopkg.in/tucnak/telebot.v3"
)

func (h *handler) OnStart(ctx tele.Context) error {

	//TODO: implement license agreement
	//if true {
	//	return ctx.Send(h.Layout.String("user_agreement"), h.Layout.Markup(ctx, "start_notagreed"))
	//}

	err := h.s.Users.Create(storage.User{
		ID: int64(ctx.Sender().ID),
	})

	if err != nil {
		return err
	}

	return ctx.Send("мда", h.lt.Markup(ctx, "start"))
}

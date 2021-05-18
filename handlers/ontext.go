package handlers

import (
	"github.com/Duvewo/cryptobot/storage"
	tele "gopkg.in/tucnak/telebot.v3"
)

func (h *handler) OnText(ctx tele.Context) error {

	btns := h.lt.Buttons()

	for key, button := range btns {
		if button.Text == ctx.Text() {
			if err := h.s.Users.Update(storage.User{
				ID:    int64(ctx.Sender().ID),
				State: key,
			}); err != nil {
				return err
			}
			return ctx.Send(".", h.lt.Markup(ctx, key))
		}
	}

	//user, err := h.Storage.Users.ByID(storage.User{
	//	ID: int64(ctx.Sender().ID),
	//})
	//
	//if err != nil {
	//	return err
	//}

	return ctx.Send(h.lt.String("command_not_found"))

}

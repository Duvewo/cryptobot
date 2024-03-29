package telebot

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
)

// Recipient is any possible endpoint you can send
// messages to: either user, group or a channel.
type Recipient interface {
	// Must return legit Telegram chat_id or username
	Recipient() string
}

// Sendable is any object that can send itself.
//
// This is pretty cool, since it lets bots implement
// custom Sendables for complex kind of media or
// chat objects spanning across multiple messages.
//
type Sendable interface {
	Send(*Bot, Recipient, *SendOptions) (*Message, error)
}

// Send delivers media through bot b to recipient.
func (p *Photo) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id": to.Recipient(),
		"caption": p.Caption,
	}
	b.embedSendOptions(params, opt)

	msg, err := b.sendObject(&p.File, "photo", params, nil)
	if err != nil {
		return nil, err
	}

	msg.Photo.File.stealRef(&p.File)
	*p = *msg.Photo
	p.Caption = msg.Caption

	return msg, nil
}

// Send delivers media through bot b to recipient.
func (a *Audio) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id":   to.Recipient(),
		"caption":   a.Caption,
		"performer": a.Performer,
		"title":     a.Title,
		"file_name": a.FileName,
	}
	b.embedSendOptions(params, opt)

	if a.Duration != 0 {
		params["duration"] = strconv.Itoa(a.Duration)
	}

	msg, err := b.sendObject(&a.File, "audio", params, thumbnailToFilemap(a.Thumbnail))
	if err != nil {
		return nil, err
	}

	if msg.Audio != nil {
		msg.Audio.File.stealRef(&a.File)
		*a = *msg.Audio
		a.Caption = msg.Caption
	}

	if msg.Document != nil {
		msg.Document.File.stealRef(&a.File)
		a.File = msg.Document.File
	}

	return msg, nil
}

// Send delivers media through bot b to recipient.
func (d *Document) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id":   to.Recipient(),
		"caption":   d.Caption,
		"file_name": d.FileName,
	}
	b.embedSendOptions(params, opt)

	if d.FileSize != 0 {
		params["file_size"] = strconv.Itoa(d.FileSize)
	}

	msg, err := b.sendObject(&d.File, "document", params, thumbnailToFilemap(d.Thumbnail))
	if err != nil {
		return nil, err
	}

	msg.Document.File.stealRef(&d.File)
	*d = *msg.Document
	d.Caption = msg.Caption

	return msg, nil
}

// Send delivers media through bot b to recipient.
func (s *Sticker) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id": to.Recipient(),
	}
	b.embedSendOptions(params, opt)

	msg, err := b.sendObject(&s.File, "sticker", params, nil)
	if err != nil {
		return nil, err
	}

	msg.Sticker.File.stealRef(&s.File)
	*s = *msg.Sticker

	return msg, nil
}

// Send delivers media through bot b to recipient.
func (v *Video) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id":   to.Recipient(),
		"caption":   v.Caption,
		"file_name": v.FileName,
	}
	b.embedSendOptions(params, opt)

	if v.Duration != 0 {
		params["duration"] = strconv.Itoa(v.Duration)
	}
	if v.Width != 0 {
		params["width"] = strconv.Itoa(v.Width)
	}
	if v.Height != 0 {
		params["height"] = strconv.Itoa(v.Height)
	}
	if v.SupportsStreaming {
		params["supports_streaming"] = "true"
	}

	msg, err := b.sendObject(&v.File, "video", params, thumbnailToFilemap(v.Thumbnail))
	if err != nil {
		return nil, err
	}

	if vid := msg.Video; vid != nil {
		vid.File.stealRef(&v.File)
		*v = *vid
		v.Caption = msg.Caption
	} else if doc := msg.Document; doc != nil {
		// If video has no sound, Telegram can turn it into Document (GIF)
		doc.File.stealRef(&v.File)

		v.Caption = doc.Caption
		v.MIME = doc.MIME
		v.Thumbnail = doc.Thumbnail
	}

	return msg, nil
}

// Send delivers animation through bot b to recipient.
// @see https://core.telegram.org/bots/api#sendanimation
func (a *Animation) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id":   to.Recipient(),
		"caption":   a.Caption,
		"file_name": a.FileName,
	}
	b.embedSendOptions(params, opt)

	if a.Duration != 0 {
		params["duration"] = strconv.Itoa(a.Duration)
	}
	if a.Width != 0 {
		params["width"] = strconv.Itoa(a.Width)
	}
	if a.Height != 0 {
		params["height"] = strconv.Itoa(a.Height)
	}

	// file_name is required, without file_name GIFs sent as document
	if params["file_name"] == "" && a.File.OnDisk() {
		params["file_name"] = filepath.Base(a.File.FileLocal)
	}

	msg, err := b.sendObject(&a.File, "animation", params, nil)
	if err != nil {
		return nil, err
	}

	if msg.Animation != nil {
		msg.Animation.File.stealRef(&a.File)
		*a = *msg.Animation
	} else {
		*a = Animation{
			File:      msg.Document.File,
			Thumbnail: msg.Document.Thumbnail,
			MIME:      msg.Document.MIME,
			FileName:  msg.Document.FileName,
		}
	}
	a.Caption = msg.Caption

	return msg, nil
}

// Send delivers media through bot b to recipient.
func (v *Voice) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id": to.Recipient(),
	}
	b.embedSendOptions(params, opt)

	if v.Duration != 0 {
		params["duration"] = strconv.Itoa(v.Duration)
	}

	msg, err := b.sendObject(&v.File, "voice", params, nil)
	if err != nil {
		return nil, err
	}

	msg.Voice.File.stealRef(&v.File)
	*v = *msg.Voice

	return msg, nil
}

// Send delivers media through bot b to recipient.
func (v *VideoNote) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id": to.Recipient(),
	}
	b.embedSendOptions(params, opt)

	if v.Duration != 0 {
		params["duration"] = strconv.Itoa(v.Duration)
	}
	if v.Length != 0 {
		params["length"] = strconv.Itoa(v.Length)
	}

	msg, err := b.sendObject(&v.File, "videoNote", params, thumbnailToFilemap(v.Thumbnail))
	if err != nil {
		return nil, err
	}

	msg.VideoNote.File.stealRef(&v.File)
	*v = *msg.VideoNote

	return msg, nil
}

// Send delivers media through bot b to recipient.
func (x *Location) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id":     to.Recipient(),
		"latitude":    fmt.Sprintf("%f", x.Lat),
		"longitude":   fmt.Sprintf("%f", x.Lng),
		"live_period": strconv.Itoa(x.LivePeriod),
	}
	b.embedSendOptions(params, opt)

	data, err := b.Raw("sendLocation", params)
	if err != nil {
		return nil, err
	}

	return extractMessage(data)
}

// Send delivers media through bot b to recipient.
func (v *Venue) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id":         to.Recipient(),
		"latitude":        fmt.Sprintf("%f", v.Location.Lat),
		"longitude":       fmt.Sprintf("%f", v.Location.Lng),
		"title":           v.Title,
		"address":         v.Address,
		"foursquare_id":   v.FoursquareID,
		"foursquare_type": v.FoursquareType,
	}
	b.embedSendOptions(params, opt)

	data, err := b.Raw("sendVenue", params)
	if err != nil {
		return nil, err
	}

	return extractMessage(data)
}

// Send delivers invoice through bot b to recipient.
func (i *Invoice) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id":                       to.Recipient(),
		"title":                         i.Title,
		"description":                   i.Description,
		"start_parameter":               i.Start,
		"payload":                       i.Payload,
		"provider_token":                i.Token,
		"currency":                      i.Currency,
		"need_name":                     strconv.FormatBool(i.NeedName),
		"need_phone_number":             strconv.FormatBool(i.NeedPhoneNumber),
		"need_email":                    strconv.FormatBool(i.NeedEmail),
		"need_shipping_address":         strconv.FormatBool(i.NeedShippingAddress),
		"send_phone_number_to_provider": strconv.FormatBool(i.SendPhoneNumber),
		"send_email_to_provider":        strconv.FormatBool(i.SendEmail),
		"is_flexible":                   strconv.FormatBool(i.Flexible),
	}
	if i.Photo != nil {
		if i.Photo.FileURL != "" {
			params["photo_url"] = i.Photo.FileURL
		}
		if i.PhotoSize > 0 {
			params["photo_size"] = strconv.Itoa(i.PhotoSize)
		}
		if i.Photo.Width > 0 {
			params["photo_width"] = strconv.Itoa(i.Photo.Width)
		}
		if i.Photo.Height > 0 {
			params["photo_height"] = strconv.Itoa(i.Photo.Height)
		}
	}
	if len(i.Prices) > 0 {
		data, _ := json.Marshal(i.Prices)
		params["prices"] = string(data)
	}
	b.embedSendOptions(params, opt)

	data, err := b.Raw("sendInvoice", params)
	if err != nil {
		return nil, err
	}

	return extractMessage(data)
}

// Send delivers poll through bot b to recipient.
func (p *Poll) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id":                 to.Recipient(),
		"question":                p.Question,
		"type":                    string(p.Type),
		"is_closed":               strconv.FormatBool(p.Closed),
		"is_anonymous":            strconv.FormatBool(p.Anonymous),
		"allows_multiple_answers": strconv.FormatBool(p.MultipleAnswers),
		"correct_option_id":       strconv.Itoa(p.CorrectOption),
	}
	if p.Explanation != "" {
		params["explanation"] = p.Explanation
		params["explanation_parse_mode"] = p.ParseMode
	}
	if p.OpenPeriod != 0 {
		params["open_period"] = strconv.Itoa(p.OpenPeriod)
	} else if p.CloseUnixdate != 0 {
		params["close_date"] = strconv.FormatInt(p.CloseUnixdate, 10)
	}
	b.embedSendOptions(params, opt)

	var options []string
	for _, o := range p.Options {
		options = append(options, o.Text)
	}

	opts, _ := json.Marshal(options)
	params["options"] = string(opts)

	data, err := b.Raw("sendPoll", params)
	if err != nil {
		return nil, err
	}

	return extractMessage(data)
}

// Send delivers dice through bot b to recipient.
func (d *Dice) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id": to.Recipient(),
		"emoji":   string(d.Type),
	}
	b.embedSendOptions(params, opt)

	data, err := b.Raw("sendDice", params)
	if err != nil {
		return nil, err
	}

	return extractMessage(data)
}

// Send delivers game through bot b to recipient.
func (g *Game) Send(b *Bot, to Recipient, opt *SendOptions) (*Message, error) {
	params := map[string]string{
		"chat_id":         to.Recipient(),
		"game_short_name": g.Name,
	}
	b.embedSendOptions(params, opt)

	data, err := b.Raw("sendGame", params)
	if err != nil {
		return nil, err
	}

	return extractMessage(data)
}

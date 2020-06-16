package Mail

import (
	"context"
	"errors"
	"gopkg.in/gomail.v2"
)

type Message struct {
	To      []string
	Subject string
	Body    string
}

type Reply struct {
}

func (Mail) SendMail(_ context.Context,
	message Message, reply *Reply) (err error) {
	var m *gomail.Message

	m = gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(MailSetting.Sender, MailSetting.Nickname))
	m.SetHeader("To", message.To...)
	m.SetHeader("Subject", message.Subject)
	m.SetBody("text/html", message.Body)
	d := gomail.NewDialer(MailSetting.Host, MailSetting.Port,
		MailSetting.Sender, MailSetting.Password)
	if err = d.DialAndSend(m); err != nil {
		return errors.New("发送失败")
	}
	return
}

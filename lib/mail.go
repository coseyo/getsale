package lib

import (
	"net/smtp"
	"strings"

	"github.com/go-gomail/gomail"
)

type Mail struct {
	dialer  gomail.Dialer
	message *gomail.Message
}

func (this *Mail) Dial(host string, port int, username, password string) {
	this.dialer = gomail.Dialer{Host: host, Port: port, Auth: smtp.PlainAuth("", username, password, host)}
	this.message = gomail.NewMessage()
}

func (this *Mail) SetReceiver(receivers string) {
	s := strings.Split(receivers, ",")
	this.message.SetHeader("To", s...)
}

func (this *Mail) SetSender(sender string) {
	this.message.SetHeader("From", sender)
}

func (this *Mail) Send(subject, body string) error {
	this.message.SetHeader("Subject", subject)
	this.message.SetBody("text/html; charset=utf-8", body)
	return this.dialer.DialAndSend(this.message)
}

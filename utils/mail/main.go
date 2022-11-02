package mail

import (
	"errors"

	md "youbei/models"
	"gopkg.in/gomail.v2"
)

type MailConn struct {
	FromUser   string
	FromPasswd string
	Host       string
	Port       int
	ToUsers    []string
}

func MailInit() (*MailConn, error) {
	ms := new(md.MailServer)
	bol, err := md.Localdb().Get(ms)
	if err != nil {
		return nil, err
	}
	if !bol {
		return nil, errors.New("from user not found")
	}
	mailcon := new(MailConn)
	mailcon.FromPasswd = ms.FromPasswd
	mailcon.FromUser = ms.FromUser
	mailcon.Host = ms.Host
	mailcon.Port = ms.Port
	mailcon.ToUsers = ms.MailTos
	return mailcon, nil
}

func (c *MailConn) SendMail(subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "Youbei"+"<"+c.FromUser+">")
	m.SetHeader("To", c.ToUsers...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(c.Host, c.Port, c.FromUser, c.FromPasswd)

	err := d.DialAndSend(m)
	return err

}

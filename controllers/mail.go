package controllers

import (
	"errors"

	md "youbei/models"

	"youbei/utils/mail"

	"github.com/gin-gonic/gin"
)

// MailInput ...
type MailSendInput struct {
	User string `json:"user"`
}

//MailServerUpdate ...
func MailServerUpdate(c *gin.Context) {
	ob := new(md.MailServer)
	if err := c.Bind(ob); err != nil {
		APIReturn(c, 500, "解析数据失败", err.Error())
		return
	}

	if err := ob.Update(); err != nil {
		APIReturn(c, 500, "修改mail server失败", err.Error())
		return
	}

	APIReturn(c, 200, "修改成功", nil)
}

func GetMail(c *gin.Context) {
	mailserver := new(md.MailServer)
	if bol, err := md.Localdb().Get(mailserver); err != nil {
		APIReturn(c, 500, "获取失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "未配置", errors.New("未配置"))
			return
		}
	}

	APIReturn(c, 200, "获取成功", mailserver)
}

//MailSendAdd ...
func MailSendAdd(sub, body string) error {
	sendmail := new(md.MailSend)
	mailcon, err := mail.MailInit()
	if err != nil {
		return err
	}
	sendmail.FromUser = mailcon.FromPasswd
	sendmail.FromPasswd = mailcon.FromPasswd
	sendmail.Host = mailcon.Host
	sendmail.Port = mailcon.Port
	sendmail.ToUsers = mailcon.ToUsers
	sendmail.Body = body
	sendmail.Subject = sub
	if err := mailcon.SendMail(sub, body); err != nil {
		return err
	}
	return sendmail.Add()
}

type MailText struct {
	Sub  string `json:"sub"`
	Text string `json:"text"`
}

func MailTest(c *gin.Context) {
	mailtext := new(MailText)
	if err := c.Bind(mailtext); err != nil {
		APIReturn(c, 500, "解析失败", err.Error())
		return
	}

	if err := MailSendAdd(mailtext.Sub, mailtext.Text); err != nil {
		APIReturn(c, 500, "失败", err.Error())
		return
	}

	APIReturn(c, 200, "成功", nil)
}

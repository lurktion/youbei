package md

import (
	"errors"

	"github.com/segmentio/ksuid"
)

// MailServer ...
type MailServer struct {
	ID         string   `json:"id" xorm:"pk notnull unique 'id'"`
	Host       string   `json:"host" xorm:"'host'"`
	Port       int      `json:"port" xorm:"'port'"`
	Status     int      `json:"status" xorm:"'status'"`
	FromUser   string   `json:"fromuser" xorm:"'fromuser'"`
	FromPasswd string   `json:"frompassword" xorm:"'frompassword'"`
	Created    int64    `json:"created" xorm:"'created'"`
	MailTos    []string `json:"mailtos" xorm:"json 'mailtos'"`
}

// MailSend ...
type MailSend struct {
	ID         string   `json:"id" xorm:"pk notnull unique 'id'"`
	Created    int64    `json:"created" xorm:"'created'"`
	Subject    string   `json:"subject" xorm:"'subject'"`
	Body       string   `json:"body" xorm:"'body'"`
	ToUsers    []string `json:"tousers" xorm:"json 'tousers'"`
	FromUser   string   `json:"fromuser" xorm:"'fromuser'"`
	FromPasswd string   `json:"frompasswd" xorm:"-"`
	Host       string   `json:"host" xorm:"-"`
	Port       int      `json:"port" xorm:"-"`
	Status     int      `json:"status" xorm:"'status'"`
}

// MailServerInit ...
func MailServerInit() error {
	c := new(MailServer)
	c.ID = ksuid.New().String()
	c.Created = NowTime()
	bol, err := localdb.Get(new(MailServer))
	if err != nil {
		return err
	}
	if !bol {
		if _, err := localdb.Insert(c); err != nil {
			return err
		}
	}
	return nil
}

// Update ...
func (c *MailServer) Update() error {
	c.Created = NowTime()
	if c.ID == "" {
		newc := new(MailServer)
		bol, err := localdb.Get(newc)
		if err != nil || !bol {
			return errors.New("from user not get")
		}
		c.ID = newc.ID
	}
	_, err := localdb.ID(c.ID).AllCols().Update(c)
	return err
}

// Add ...
func (c *MailSend) Add() error {
	c.ID = ksuid.New().String()
	c.Created = NowTime()
	_, err := localdb.Insert(c)
	return err
}

// Update ...
func (c *MailSend) Update() error {
	c.Created = NowTime()
	_, err := localdb.ID(c.ID).Update(c)
	return err
}

// Delete ...
func (c *MailSend) Delete() error {
	_, err := localdb.ID(c.ID).Delete(new(MailSend))
	return err
}

package md

import (
	"github.com/segmentio/ksuid"
)

type Host struct {
	ID       string `json:"id" xorm:"pk notnull unique 'id'"`
	Protocol string `json:"protocol" xorm:"'protocol'"`
	Name     string `json:"name" xorm:"'name'"`
	HostAddr string `json:"hostaddr" xorm:"'hostaddr'"`
	Port     string `json:"port" xorm:"'port'"`
	Username string `json:"username" xorm:"'username'"`
	Password string `json:"password" xorm:"'password'"`
	Created  int64  `json:"created" xorm:"'created'"`
}

func (c *Host) Add() error {
	c.ID = ksuid.New().String()
	c.Created = NowTime()
	_, err := localdb.Insert(c)
	return err
}

func (c *Host) Update() error {
	c.Created = NowTime()
	_, err := localdb.ID(c.ID).AllCols().Update(c)
	return err
}

func (c *Host) Delete() error {
	_, err := localdb.ID(c.ID).Delete(new(Host))
	return err
}

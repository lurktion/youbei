package md

import (
	"github.com/segmentio/ksuid"
)

type SshTask struct {
	ID       string `json:"id" xorm:"pk notnull unique 'id'"`
	TaskType string `json:"tasktype" xorm:"'tasktype'"`
	Host     string `json:"host" xorm:"'host'"`
	SshUser  string `json:"sshuser" xorm:"'sshuser'"`
	SshPort  string `json:"sshport" xorm:"'sshport'"`
	SshPwd   string `json:"sshpwd" xorm:"'sshpwd'"`

	DbHost string `json:"dbhost" xorm:"'dbhost'"`
	DbPort string `json:"dbport" xorm:"'dbport'"`
	DbName string `json:"dbname" xorm:"'dbname'"`
	DbUser string `json:"dbuser" xorm:"'dbuser'"`
	DbPwd  string `json:"dbpwd" xorm:"'dbpwd'"`
	Char   string `json:"char" xorm:"'char'"`

	SavePath string   `json:"savepath" xorm:"'savepath'"`
	Zippwd   string   `json:"zippwd" xorm:"'zippwd'"`
	RS       []string `json:"rs" xorm:"-"`
	Crontab  string   `json:"crontab" xorm:"'crontab'"`
	Expire   int      `json:"expire" xorm:"'expire'"`
	Created  int64    `json:"created" xorm:"'created'"`
}

func (c *SshTask) Add() error {
	c.ID = ksuid.New().String()
	c.Created = NowTime()
	_, err := localdb.Insert(c)
	return err
}

func (c *SshTask) Update() error {
	_, err := localdb.ID(c.ID).AllCols().Update(c)
	return err
}

func (c *SshTask) Delete() error {
	_, err := localdb.ID(c.ID).Delete(new(SshTask))
	return err
}

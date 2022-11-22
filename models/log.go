package md

import (
	"errors"
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
)

// Log 日志结构体
type Log struct {
	ID            string          `json:"id" xorm:"pk notnull unique 'id'"`
	Name          string          `json:"name" xorm:"'name'"`
	Tid           string          `json:"tid" xorm:"notnull 'tid'"`
	Dbtype        string          `json:"dbtype" xorm:"-"`
	Status        int             `json:"status" xorm:"notnull 'status'"`
	Localfilepath string          `json:"localfilepath" xorm:"'localfilepath'"`
	Msg           string          `json:"msg" xorm:"msg"`
	Created       int64           `json:"created" xorm:"'created'"`
	Deleted       int64           `json:"deleted" xorm:"'deleted'"`
	DBInfo        Task            `json:"dbinfo" xorm:"-"`
	RS            []RemoteStorage `json:"rs" xorm:"-"`
	Errors        error           `json:"errors" xorm:"-"`

	IfSaveLocal int    `json:"ifsavelocal" xorm:"default(0) 'ifsavelocal'"` // 0 保存 1 不保存
	Rlogs       []Rlog `json:"rlogs" xorm:"-"`
}

// NewLog ...
func NewLog(tid string) (*Log, error) {
	log := new(Log)
	log.ID = ksuid.New().String()
	log.Tid = tid
	log.Created = time.Now().Unix()
	log.Status = 1
	if name, err := CreateId("LOG", "L"); err != nil {
		return log, err
	} else {
		log.Name = name
	}
	task := Task{}
	if bol, err := localdb.ID(tid).Get(&task); err != nil {
		return log, err
	} else {
		if !bol {
			return log, errors.New("生成日志失败，任务不存在")
		}
	}

	if task.Types == 1 && task.DBType == "mysql" && task.Cmds != "" {
		log.IfSaveLocal = 4
	}

	if _, err := localdb.Insert(log); err != nil {
		return log, err
	}
	fmt.Println(log.ID)
	return log, nil
}

// Add ...
func (log *Log) Add() error {
	if log.Errors != nil {
		log.Msg = log.Errors.Error()
		log.Status = 2
	} else {
		log.Status = 0
	}
	log.Created = time.Now().Unix()
	if _, err := localdb.Insert(log); err != nil {
		return err
	}
	return nil
}

// Add ...
func (log *Log) Get() error {
	if bol, err := localdb.ID(log.ID).Get(log); err != nil {
		return err
	} else {
		if !bol {
			return errors.New("日志不存在")
		}
	}

	return nil
}

func (log *Log) Update(cols ...string) error {
	if _, err := localdb.ID(log.ID).Cols(cols...).Update(log); err != nil {
		return err
	}
	return nil
}

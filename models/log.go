package md

import (
	"errors"
	"time"

	"github.com/segmentio/ksuid"
)

// Log 日志结构体
type Log struct {
	ID            string          `json:"id" xorm:"pk notnull unique 'id'"`
	Tid           string          `json:"tid" xorm:"notnull 'tid'"`
	Dbtype        string          `json:"dbtype" xorm:"_"`
	Status        int             `json:"status" xorm:"notnull 'status'"`
	Localfilepath string          `json:"localfilepath" xorm:"'localfilepath'"`
	Msg           string          `json:"msg" xorm:"msg"`
	Created       int64           `json:"created" xorm:"'created'"`
	Deleted       int64           `json:"deleted" xorm:"'deleted'"`
	DBInfo        Task            `json:"dbinfo" xorm:"-"`
	RS            []RemoteStorage `json:"rs" xorm:"-"`
	Errors        error           `json:"errors" xorm:"-"`
}

// NewLog ...
func NewLog(tid string, localfilepath string, status string) (string, error) {
	log := new(Log)
	log.ID = ksuid.New().String()
	log.Tid = tid
	log.Localfilepath = localfilepath
	if status != "" {
		log.Errors = errors.New(status)
	}
	return log.ID, log.Add()
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

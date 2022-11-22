package md

import (
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
)

// Rlog ...
type Rlog struct {
	ID            string        `json:"id" xorm:"pk notnull unique 'id'"`
	Name          string        `json:"name" xorm:"'name'"`
	Lid           string        `json:"lid" xorm:"'lid'"`
	Rid           string        `json:"rid" xorm:"'rid'"`
	Tid           string        `json:"tid" xorm:"'tid'"`
	Localfilepath string        `json:"localfilepath" xorm:"'localfilepath'"`
	Created       int64         `json:"created" xorm:"'created'"`
	Status        int           `json:"status" xorm:"'status'"`
	Msg           string        `json:"msg" xorm:"'msg'"`
	DBInfo        Task          `json:"dbinfo" xorm:"-"`
	RSInfo        RemoteStorage `json:"rsinfo" xorm:"-"`
	LogInfo       Log           `json:"loginfo" xorm:"-"`
	YsUploadFile  YsUploadFile  `json:"ysuploadfile" xorm:"-"`
}

// AddNewRemoteSendLog ...
func AddNewRemoteSendLog(lid, tid, rid string) (*Rlog, error) {
	rlog := new(Rlog)
	rlog.ID = ksuid.New().String()
	rlog.Lid = lid
	rlog.Tid = tid
	rlog.Rid = rid
	rlog.Created = time.Now().Unix()
	rlog.Status = 1
	if name, err := CreateId("LOG", "R"); err != nil {
		return rlog, err
	} else {
		rlog.Name = name
	}
	if _, err := localdb.Insert(rlog); err != nil {
		return nil, err
	}
	return rlog, nil
}

// Update ...
func (r *Rlog) Update(err error) {
	if err != nil {
		r.Status = 2
		r.Msg = err.Error()
	} else {
		r.Status = 0
	}
	localdb.ID(r.ID).Cols("status", "msg").Update(r)
}

// RlogAll ...
func RemoteSendLogFindAll(tid string, page, limit int) ([]Rlog, error) {
	logs := []Rlog{}
	var err error
	if tid == "" {
		err = localdb.Desc("created").Limit(limit, limit*(page-1)).Find(&logs)
	} else {
		err = localdb.Where("tid=?", tid).Desc("created").Limit(limit, limit*(page-1)).Find(&logs)
	}

	if len(logs) > 0 {
		for k, v := range logs {
			ts := Task{}
			if bol, err := localdb.ID(v.Tid).Get(&ts); err == nil && bol {
				logs[k].DBInfo = ts
			}
			rs := RemoteStorage{}
			if bol, err := localdb.ID(v.Rid).Get(&rs); err == nil && bol {
				logs[k].RSInfo = rs
				if rs.Types == "Yserver" {
					yuf := new(YsUploadFile)
					yuf.Lid = v.ID
					_, err := localdb.Get(yuf)
					if err == nil {
						logs[k].YsUploadFile = *yuf
						ypf := []YsPacket{}
						err = localdb.Where("yid=?", yuf.ID).Asc("sortid").Find(&ypf)
						if err == nil {
							logs[k].YsUploadFile.YsPackets = ypf
						}
					}
				}
			}
			ls := Log{}
			if bol, err := localdb.ID(v.Lid).Get(&ls); err == nil && bol {
				logs[k].LogInfo = ls
			}
		}
	}

	return logs, err
}

func CountRemoteSendLogByStatus(status int) int64 {
	log := Rlog{}
	if status > 2 {
		status = 2
	} else if status < 0 {
		status = 0
	}
	count, err := localdb.Where("status = ?", status).Count(log)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return count

}

package md

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/segmentio/ksuid"
)

// Yserver ...
type Yserver struct {
	ID       string `json:"id" xorm:"pk notnull unique 'id'"`
	SavePath string `json:"savepath" xorm:"'savepath'"`
	Enable   bool   `json:"enable" xorm:"'enable'"`
	Username string `json:"username" xorm:"'username'"`
	Password string `json:"password" xorm:"'password'"`
	Port     int    `json:"port" xorm:"-"`
}

// YserverFile ...
type YserverFile struct {
	ID             string          `json:"id" xorm:"pk notnull unique 'id'"`
	FileName       string          `json:"filename" xorm:"'filename'"`
	FileSize       int64           `json:"filesize" xorm:"'filesize'"`
	Packet         int64           `json:"packet" xorm:"packet"`
	Created        int64           `json:"created" xorm:"'created'"`
	Status         int             `json:"status" xorm:"'status'"`
	YserverPackets []YserverPacket `json:"yps" xorm:"-"`
}

// YserverPacket ...
type YserverPacket struct {
	ID         string `json:"id" xorm:"pk notnull unique 'id'"`
	Fid        string `json:"fid" xorm:"'fid'"`
	PacketPath string `json:"packetpath" xorm:"'packetpath'"`
	Sort       int    `json:"sort" xorm:"'sort'"`
	Created    int64  `json:"created" xorm:"'created'"`
}

// EnableYserver ...
func (c *Yserver) EnableYserver() error {
	_, err := localdb.ID(c.ID).Cols("savepath", "enable", "username", "password").Update(c)
	return err
}

// DisableYserver ...
func DisableYserver(id string) error {
	ey := Yserver{}
	_, err := localdb.ID(id).Delete(&ey)
	return err
}

// AddFile ...
func AddFile(id string, fname string, savedir string, fsize int64, packet int64) error {
	fs := Yserver{}
	bol, err := localdb.Get(&fs)
	if err != nil {
		return err
	}
	if !bol {
		return errors.New("not found")
	}

	yf := YserverFile{}
	yf.ID = id
	savepath := strings.TrimRight(fs.SavePath, "/")
	savedir = strings.TrimLeft(strings.TrimRight(savedir, "/"), "/")
	os.MkdirAll(savepath+"/"+savedir, os.ModeDir)
	yf.FileName = savedir + "/" + fname
	yf.FileSize = fsize
	yf.Packet = packet
	yf.Created = time.Now().Unix()
	yf.Status = 1
	_, err = localdb.Insert(&yf)
	if err != nil {
		return err
	}
	return err
}

// FinshFile ...
func FinshFile(id string, status int) error {
	yf := YserverFile{}
	yf.Status = status
	_, err := localdb.ID(id).Cols("status").Update(&yf)
	return err
}

// FindFile ...
func FindFile(id string) (YserverFile, error) {
	yp := YserverFile{}
	bol, err := localdb.ID(id).Get(&yp)
	if err != nil {
		return yp, err
	}
	if !bol {
		return yp, errors.New("serverfile not found")
	}
	return yp, err
}

// AllFile ...
func AllFile(page, limit int) ([]YserverFile, error) {
	yps := []YserverFile{}
	err := localdb.Desc("created").Limit(limit, limit*(page-1)).Find(&yps)
	return yps, err
}

// Filecount ...
func Filecount() (int64, error) {
	yps := YserverFile{}
	return localdb.Count(&yps)
}

// AddPacket ...
func AddPacket(fid string, sort int, packetpath string) error {
	yp := YserverPacket{}
	yp.ID = ksuid.New().String()
	yp.Fid = fid
	yp.Sort = sort
	yp.PacketPath = packetpath
	yp.Created = time.Now().Unix()
	_, err := localdb.Insert(&yp)
	if err != nil {
		return err
	}
	return err
}

// AllPacket ...
func AllPacket(id string) ([]YserverPacket, error) {
	ap := []YserverPacket{}
	err := localdb.Where("fid=?", id).Asc("sort").Find(&ap)
	return ap, err
}

// InitYservefunc ...
func InitYservefunc() {
	fs := []Yserver{}
	if err := localdb.Find(&fs); err != nil {
		panic(err)
	}
	if len(fs) < 1 {
		f := Yserver{}
		f.ID = ksuid.New().String()
		f.Enable = false
		if _, err := localdb.Insert(&f); err != nil {
			panic(err)
		}
	}
}

package md

import (
	"time"
)

// YsUploadFile ...
type YsUploadFile struct {
	ID                 string     `json:"id" xorm:"pk notnull unique 'id'"`
	Lid                string     `json:"lid" xorm:"'lid'"`
	UploadFileServerID string     `json:"ufsid" xorm:"'ufsid'"`
	SrcFilePath        string     `json:"srcfilepath" xorm:"'srcfilepath'"`
	Created            int64      `json:"created" xorm:"'created'"`
	Size               int64      `json:"size" xorm:"'size'"`
	PacketNum          int64      `json:"packetnum" xorm:"'packetnum'"`
	Status             int        `json:"status" xorm:"'status'"`
	Msg                string     `json:"msg" xorm:"'msg'"`
	YsPackets          []YsPacket `json:"packets" xorm:"-"`
}

// YsPackets ...
type YsPacket struct {
	ID              string `json:"id" xorm:"pk notnull unique 'id'"`
	Yid             string `json:"yid" xorm:"'yid'"`
	SortID          int    `json:"sortid" xorm:"'sortid'"`
	SrcPacketPath   string `json:"srcpacketpath" xorm:"'srcpacketpath'"`
	Offset          int64  `json:"offset" xorm:"'offset'"`
	UploadPacketURL string `json:"uploadpacketurl" xorm:"'uploadpacketurl'"`
	Status          int    `json:"status" xorm:"'status'"`
	Msg             string `json:"msg" xorm:"'msg'"`
}

// AddYsFileLog ...
func (ysuploadfile *YsUploadFile) AddYsFileLog(err error) error {
	ysuploadfile.Created = time.Now().Unix()
	if err != nil {
		ysuploadfile.Status = 2
		ysuploadfile.Msg = err.Error()
	} else {
		ysuploadfile.Status = 1
	}
	_, err = localdb.Insert(ysuploadfile)
	return err
}

// AddYsPacketLog ...
func (yuf *YsUploadFile) AddYsPacketLog(ysp YsPacket) error {
	_, err := localdb.Insert(&ysp)
	return err
}

// UpdateYspacketLog ...
func (yuf *YsUploadFile) UpdateYspacketLog(sortid int, err error) error {
	pf := new(YsPacket)
	if err != nil {
		pf.Status = 2
		pf.Msg = err.Error()
	} else {
		pf.Status = 0
	}
	_, err = localdb.Where("yid=? and sortid=?", yuf.ID, sortid).Cols("status", "msg").Update(pf)
	return err

}

package controllers

import (
	"errors"
	"strconv"

	md "youbei/models"

	"github.com/gin-gonic/gin"
)

// Loglist ...
func Loglist(c *gin.Context) {
	tid := c.Param("id")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("count"))

	logs := []md.Log{}
	if tid == "" {
		if err := md.Localdb().Desc("created").Limit(limit, limit*(page-1)).Find(&logs); err != nil {
			APIReturn(c, 500, "获取列表失败", err.Error())
			return
		}
	} else {
		if err := md.Localdb().Where("tid=?", tid).Desc("created").Limit(limit, limit*(page-1)).Find(&logs); err != nil {
			APIReturn(c, 500, "获取列表失败", err.Error())
			return
		}
	}
	if len(logs) > 0 {
		for k, v := range logs {
			ts := new(md.Task)
			if bol, err := md.Localdb().ID(v.Tid).Get(ts); err == nil && bol {
				logs[k].DBInfo = *ts
			}
		}
	}

	log := md.Log{}
	title := int64(0)
	var err error
	if tid == "" {
		title, err = md.Localdb().Count(&log)
	} else {
		title, err = md.Localdb().Where("tid=?", tid).Count(&log)
	}
	if err != nil {
		APIReturn(c, 500, "获取总数失败", err.Error())
		return
	}
	rep := map[string]interface{}{"count": title, "data": logs}
	APIReturn(c, 200, "获取列表成功", &rep)
}

// ShowrLog ...
func ShowrLog(c *gin.Context) {
	id := c.Param("id")
	rlog := new(md.Rlog)
	if bol, err := md.Localdb().ID(id).Get(rlog); err != nil {
		APIReturn(c, 500, "获取远程传输日志失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "日志不存在333", errors.New("日志不存在333"))
			return
		}
	}

	loginfo := md.Log{}
	if bol, err := md.Localdb().ID(rlog.Lid).Get(&loginfo); err != nil {
		APIReturn(c, 500, "获取备份日志失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "备份日志不存在", errors.New("备份日志不存在"))
			return
		}
	}

	rlog.LogInfo = loginfo

	dbinfo := md.Task{}
	if bol, err := md.Localdb().ID(rlog.Tid).Get(&dbinfo); err != nil {
		APIReturn(c, 500, "获取备份信息失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "备份信息不存在", errors.New("备份信息不存在"))
			return
		}
	}

	rlog.DBInfo = dbinfo

	rsinfo := md.RemoteStorage{}
	if bol, err := md.Localdb().ID(rlog.Rid).Get(&rsinfo); err != nil {
		APIReturn(c, 500, "获取远程信息失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "远程信息不存在", errors.New("远程信息不存在"))
			return
		}
	}

	rlog.RSInfo = rsinfo

	if rlog.RSInfo.Types == "Yserver" {
		yufinfo := md.YsUploadFile{}
		if bol, err := md.Localdb().Where("lid=?", rlog.ID).Get(&yufinfo); err != nil {
			APIReturn(c, 500, "获取上传文件信息失败", err.Error())
			return
		} else {
			if !bol {
				APIReturn(c, 500, "上传文件信息不存在", errors.New("上传文件信息不存在"))
				return
			}
		}

		ypinfo := []md.YsPacket{}
		if err := md.Localdb().Where("yid=?", yufinfo.ID).Asc("sortid").Find(&ypinfo); err != nil {
			APIReturn(c, 500, "获取上传文件切片信息失败", err.Error())
			return
		}

		yufinfo.YsPackets = ypinfo
		rlog.YsUploadFile = yufinfo
	}

	APIReturn(c, 200, "成功", rlog)
}

// ShowLog ...
func ShowLog(c *gin.Context) {
	id := c.Param("id")
	loginfo := new(md.Log)
	if bol, err := md.Localdb().ID(id).Get(loginfo); err != nil {
		APIReturn(c, 500, "获取远备份日志失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "日志不存在1111", errors.New("日志不存在wwwww"))
			return
		}
	}

	dbinfo := md.Task{}
	if bol, err := md.Localdb().ID(loginfo.Tid).Get(&dbinfo); err != nil {
		APIReturn(c, 500, "获取备份信息失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "备份信息不存在", errors.New("备份信息不存在"))
			return
		}
	}

	loginfo.DBInfo = dbinfo

	APIReturn(c, 200, "成功", loginfo)
}

// GetYserverLog ...
func GetYserverLog(c *gin.Context) {
	id := c.Param("id")
	yfinfo := new(md.YserverFile)
	if bol, err := md.Localdb().ID(id).Get(yfinfo); err != nil {
		APIReturn(c, 500, "获取日志失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "日志不存在22222", errors.New("日志不存在xxxxx"))
			return
		}
	}

	ypsinfo := []md.YserverPacket{}
	if err := md.Localdb().Where("fid=?", yfinfo.ID).Asc("sort").Find(&ypsinfo); err != nil {
		APIReturn(c, 500, "获取备份信息失败", err.Error())
		return
	}

	yfinfo.YserverPackets = ypsinfo
	APIReturn(c, 200, "成功", yfinfo)
}

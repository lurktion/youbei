package controllers

import (
	"errors"
	"time"

	md "youbei/models"
	db "youbei/utils/database"

	"github.com/gin-gonic/gin"
)

// Loglist ...
func Loglist(c *gin.Context) {
	rep := map[string]interface{}{}
	logs := []md.Log{}
	if total, err := GetRestul(c, "loglist", &logs); err != nil {
		APIReturn(c, 500, "获取列表失败", err.Error())
		return
	} else {
		rep["count"] = total
	}

	for k, v := range logs {
		ts := new(md.Task)
		if bol, err := md.Localdb().ID(v.Tid).Get(ts); err == nil && bol {
			logs[k].DBInfo = *ts
		}
		rlogs := []md.Rlog{}
		if err := md.Localdb().Where("lid=?", v.ID).Find(&rlogs); err == nil {
			logs[k].Rlogs = rlogs
		}

	}

	rep["data"] = logs

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

	cmd := md.SystemBackupCmdPath{}
	if bol, err := md.Localdb().Where("dbtype=? and status=0 and recovery=1", dbinfo.DBType).Get(&cmd); err == nil && bol {
		loginfo.Recovery = 1
	} else {
		loginfo.Recovery = 0
	}

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

func SqlRecovery(c *gin.Context) {
	id := c.Param("id")
	log := new(md.Log)
	if bol, err := md.Localdb().ID(id).Get(log); err != nil {
		APIReturn(c, 500, "获取备份信息失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "备份信息不存在", errors.New("备份信息不存在"))
			return
		}
	}

	// 设置Status为3并更新数据库
	log.Status = 3
	if _, err := md.Localdb().ID(id).Cols("Status").Update(log); err != nil {
		APIReturn(c, 500, "更新恢复状态失败", err.Error())
		return
	}

	go func() {
		err := db.MysqlCmdRecovery(log)

		log.RecoveryTime = time.Now().Unix()

		// 如果出错，设定状态为-1，否则增加RecoveryStatus值
		if err != nil {
			log.Status = -1 // 设定一个错误状态
			log.RecoveryErrMsg = err.Error()
		} else {
			log.RecoveryStatus = log.RecoveryStatus + 1
			log.Status = 0 // 假设1代表成功恢复的状态
			log.RecoveryErrMsg = ""
		}

		if _, err := md.Localdb().ID(id).Cols("Status", "RecoveryStatus", "RecoveryTime", "RecoveryErrMsg").Update(log); err != nil {
			APIReturn(c, 500, "更新恢复记录失败: ", err)
			return
		}
	}()

	APIReturn(c, 200, "恢复操作正在后台执行，请稍后查看结果", "恢复操作正在后台执行，请稍后查看结果")
}

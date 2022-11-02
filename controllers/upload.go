package controllers

import (
	"errors"
	"os"
	"strconv"

	md "youbei/models"

	"github.com/gin-gonic/gin"
)

// uploadfile ...
type uploadfile struct {
	ID        string `json:"id"`
	FileName  string `json:"filename"`
	SaveDir   string `json:"savedir"`
	Size      int64  `json:"size"`
	PacketNum int64  `json:"packetnum"`
}

//EnableServer 启动本地存储
func EnableServer(c *gin.Context) {
	ob := new(md.Yserver)
	if err := c.Bind(ob); err != nil {
		APIReturn(c, 500, "解析失败", err.Error())
		return
	}

	if err := ob.EnableYserver(); err != nil {
		APIReturn(c, 500, "服务启动失败", err.Error())
		return
	}

	APIReturn(c, 200, "服务启动成功", nil)
}

//Yserverlist 本地存储查询
func Yserverlist(c *gin.Context) {
	ys := md.Yserver{}
	if bol, err := md.Localdb().Get(&ys); err != nil {
		APIReturn(c, 500, "查询Yserver失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "Yserver不存在", errors.New("Yserver不存在"))
			return
		}
	}

	ys.Port = 8080
	APIReturn(c, 200, "获取成功", ys)
}

//DisableServer 关闭本地存储
func DisableServer(c *gin.Context) {
	ob := new(md.Yserver)
	if err := c.Bind(ob); err != nil {
		APIReturn(c, 500, "解析失败", err.Error())
		return
	}

	if err := ob.EnableYserver(); err != nil {
		APIReturn(c, 500, "服务关闭失败", err.Error())
		return
	}

	APIReturn(c, 200, "服务关闭成功", nil)
}

//UploadFile 新增上传文件任务
func UploadFile(c *gin.Context) {
	ys := md.Yserver{}
	if bol, err := md.Localdb().Get(&ys); err != nil {
		APIReturn(c, 500, "查询Yserver失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "Yserver不存在", errors.New("Yserver不存在"))
			return
		}
	}

	if !ys.Enable {
		APIReturn(c, 500, "Yserver未启动", errors.New("Yserver未启动"))
		return
	}
	if c.Query("username") != ys.Username || c.Query("password") != ys.Password {
		APIReturn(c, 500, "账号密码错误", errors.New("账号密码错误"))
		return
	}
	var ob uploadfile

	ob.ID = c.Param("id")
	ob.FileName = c.Query("filename")
	ob.Size = c.GetInt64("size")
	ob.SaveDir = c.Query("savedir")
	ob.PacketNum = c.GetInt64("packetnum")
	dirpath, err := GetFilSaveDir()
	if err != nil {
		APIReturn(c, 500, "文件保存目录获取失败", err.Error())
		return
	}

	os.MkdirAll(dirpath+"/"+ob.SaveDir, os.ModeDir)
	f, err := os.Create(dirpath + "/" + ob.SaveDir + "/" + ob.FileName)
	if err != nil {
		APIReturn(c, 500, "文件创建失败", err.Error())
		return
	}

	defer f.Close()
	if err := f.Truncate(ob.Size); err != nil {
		APIReturn(c, 500, "文件填充失败", err.Error())
		return
	}

	if err := md.AddFile(ob.ID, ob.FileName, ob.SaveDir, ob.Size, ob.PacketNum); err != nil {
		APIReturn(c, 500, "录入失败", err.Error())
		return
	}

	APIReturn(c, 200, "录入成功", nil)

}

//Uploadpacket 接收上传分片包
func Uploadpacket(c *gin.Context) {
	id := c.Param("id")
	savepath, err := GetFilSavePath(id)
	if err != nil {
		APIReturn(c, 500, "获取信息失败", err.Error())
		return
	}

	offset, err := strconv.ParseInt(c.Param("offset"), 10, 64)
	if err != nil {
		APIReturn(c, 500, "获取offset失败", err.Error())
		return
	}

	f, err := os.OpenFile(savepath, os.O_RDWR, os.ModePerm)
	if err != nil {
		APIReturn(c, 500, "打开文件失败", err.Error())
		return
	}

	defer f.Close()
	if body, err := c.GetRawData(); err != nil {
		APIReturn(c, 500, "接收数据失败", err.Error())
		return
	} else {
		if _, err := f.WriteAt(body, offset); err != nil {
			APIReturn(c, 500, "打开写入失败", err.Error())
			return
		}
	}

	APIReturn(c, 200, id, nil)

}

//UploadpacketDone 接收上传分片包
func UploadpacketDone(c *gin.Context) {
	id := c.Param("id")
	status := c.GetInt("status")
	if err := md.FinshFile(id, status); err != nil {
		APIReturn(c, 500, "完成失败", err.Error())
		return
	}

	APIReturn(c, 200, "执行成功", nil)
}

func GetFilSavePath(id string) (string, error) {
	fs := md.Yserver{}
	bol, err := md.Localdb().Get(&fs)
	if err != nil || !bol {
		return "", err
	}
	if fs.SavePath == "" {
		fs.SavePath = "."
	}

	yp, err := md.FindFile(id)
	if err != nil {
		return "", err
	}
	savepath := fs.SavePath + "/" + yp.FileName
	return savepath, nil
}

func GetFilSaveDir() (string, error) {
	fs := md.Yserver{}
	bol, err := md.Localdb().Get(&fs)
	if err != nil || !bol {
		return "", err
	}
	if fs.SavePath == "" {
		fs.SavePath = "."
	}

	return fs.SavePath, nil
}

//Uploadlogs 存储日志查询
func Uploadlogs(c *gin.Context) {
	type Query struct {
		Page  int `form:"page"`
		Count int `form:"count"`
	}

	query := Query{}
	if err := c.Bind(&query); err != nil {
		APIReturn(c, 500, "解析失败", err.Error())
		return
	}

	yps, err := md.AllFile(query.Page, query.Count)
	if err != nil {
		APIReturn(c, 500, "查询失败1", err.Error())
		return
	}

	total, err := md.Filecount()
	if err != nil {
		APIReturn(c, 500, "查询总数失败", err.Error())
		return
	}

	rep := map[string]interface{}{"count": total, "data": yps}
	APIReturn(c, 200, "查询成功", &rep)

}

package controllers

import (
	"errors"

	md "youbei/models"

	jobs "youbei/utils/jobs"

	"github.com/astaxie/beego/toolbox"
	"github.com/gin-gonic/gin"
)

func AddSshtask(c *gin.Context) {
	ob := new(md.SshTask)
	if err := c.Bind(ob); err != nil {
		APIReturn(c, 500, "解析失败", err.Error())
		return
	}
	if err := ob.Add(); err != nil {
		APIReturn(c, 500, "录入失败", err.Error())
		return
	}

	if len(ob.RS) > 0 {
		if err := md.RemoteStorageToTaskFunc(ob.ID, ob.RS); err != nil {
			APIReturn(c, 500, "添加任务失败6", err.Error())
			return
		}

	}
	toolbox.AddTask(ob.ID, toolbox.NewTask(ob.ID, ob.Crontab, jobs.SshJobs(ob.ID)))
	APIReturn(c, 200, "录入成功", nil)
}

func UpdateSshtask(c *gin.Context) {
	ob := new(md.SshTask)
	if err := c.Bind(ob); err != nil {
		APIReturn(c, 500, "解析失败", err.Error())
		return
	}

	ob.ID = c.Param("id")
	if err := ob.Update(); err != nil {
		APIReturn(c, 500, "修改失败", err.Error())
		return
	}

	if err := md.RemoteStorageToTaskFunc(ob.ID, ob.RS); err != nil {
		APIReturn(c, 500, "修改任务失败6", err.Error())
		return
	}

	toolbox.AddTask(ob.ID, toolbox.NewTask(ob.ID, ob.Crontab, jobs.SshJobs(ob.ID)))
	APIReturn(c, 200, "修改成功", nil)
}

func DeleteSshtask(c *gin.Context) {
	ob := new(md.SshTask)
	ob.ID = c.Param("id")
	if err := ob.Delete(); err != nil {
		APIReturn(c, 500, "删除失败", err.Error())
		return
	}

	APIReturn(c, 200, "删除成功", nil)
}

func SshtaskList(c *gin.Context) {
	sshtasks := []md.SshTask{}

	if err := md.Localdb().Find(&sshtasks); err != nil {
		APIReturn(c, 500, "获取失败", err.Error())
		return
	}

	title, err := md.Localdb().Count(new(md.SshTask))
	if err != nil {
		APIReturn(c, 500, "获取总数失败", err.Error())
		return
	}

	rep := map[string]interface{}{"count": title, "data": sshtasks}
	APIReturn(c, 200, "获取成功", rep)
}

func GetSshtask(c *gin.Context) {
	sshtask := new(md.SshTask)
	sshtask.ID = c.Param("id")
	if bol, err := md.Localdb().Get(sshtask); err != nil {
		APIReturn(c, 500, "获取失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "任务不存在", errors.New("任务不存在"))
			return
		}
	}
	if rstr, err := getRS(sshtask.ID); err != nil {
		APIReturn(c, 500, "获取失败", err.Error())
		return
	} else {
		sshtask.RS = rstr
	}
	APIReturn(c, 200, "获取成功", sshtask)
}

func getRS(tid string) ([]string, error) {
	rs := []md.RemoteStorageToTask{}
	if err := md.Localdb().Where("tid=?", tid).Find(&rs); err != nil {
		return []string{}, err
	}
	rstr := []string{}
	if len(rs) > 0 {
		for _, s := range rs {
			rstr = append(rstr, s.Rid)
		}
	}
	return rstr, nil
}

func RunSshJob(c *gin.Context) {
	id := c.Param("id")
	if err := jobs.SshBackup(id, false); err != nil {
		APIReturn(c, 500, "执行失败", err.Error())
		return
	}
	APIReturn(c, 200, "执行成功", nil)
}

package controllers

import (
	"errors"
	"strings"

	md "youbei/models"
	utils "youbei/utils"
	db "youbei/utils/database"
	jobs "youbei/utils/jobs"

	"github.com/astaxie/beego/toolbox"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

// Tasklist ...
func GetTasks(c *gin.Context) {
	type Query struct {
		Types string `form:"types"`
		Page  int    `form:"page"`
		Count int    `form:"count"`
	}

	query := new(Query)
	c.Bind(query)

	tasks, err := md.SelectAll(query.Types, query.Page, query.Count)
	if err != nil {
		APIReturn(c, 500, "获取列表失败", err.Error())
		return
	}

	title, err := md.TaskCount(query.Types)
	if err != nil {
		APIReturn(c, 500, "获取总数失败", err.Error())
		return
	}

	rep := map[string]interface{}{"count": title, "data": tasks}
	APIReturn(c, 200, "获取列表成功", &rep)
}

// GetTask ...
func GetTask(c *gin.Context) {
	id := c.Param("id")
	task := new(md.Task)
	task.ID = id
	if err := task.Select(); err != nil {
		APIReturn(c, 500, "获取数据失败", err.Error())
		return
	}
	APIReturn(c, 200, "获取数据成功", task)
}

// Del ...
func DeleteTask(c *gin.Context) {
	ob := new(md.Task)
	ob.ID = c.Param("id")
	if err := ob.Delete(); err != nil {
		APIReturn(c, 500, "删除任务失败1", err.Error())
		return
	}
	toolbox.DeleteTask(ob.ID)
	APIReturn(c, 200, "删除任务成功", c.Param("id"))
}

// Add ...
func AddTask(c *gin.Context) {
	ob := new(md.Task)
	if err := c.Bind(ob); err != nil {
		APIReturn(c, 500, "解析数据失败", err.Error())
		return
	}

	if ob.Crontab == "" {
		APIReturn(c, 500, "计划任务必填", errors.New("计划任务必填"))
		return
	}
	ob.DBpath = strings.Replace(ob.DBpath, "\\", "/", -1)
	if ob.DBType == "sqlite" {
		if err := submitTask(ob); err != nil {
			APIReturn(c, 500, "添加任务失败", err)
			return
		}
	} else if ob.DBType == "mysql" || ob.DBType == "mssql" || ob.DBType == "postgres" || ob.DBType == "mongodb" {
		for _, v := range ob.DBnames {
			ob.DBname = v
			if err := db.ConnectTest(ob); err != nil {
				APIReturn(c, 500, "添加任务失败mysql", err.Error())
				return
			}

			if err := submitTask(ob); err != nil {
				APIReturn(c, 500, "添加任务失败", err)
				return
			}
		}
	} else if ob.DBType == "file" {
		if bol, err := utils.PathExists(ob.DBpath); err != nil {
			APIReturn(c, 500, "file not found 1", err.Error())
			return
		} else {
			if !bol {
				APIReturn(c, 500, "file not found", errors.New("file not found"))
				return
			}
		}
		if err := submitTask(ob); err != nil {
			APIReturn(c, 500, "添加任务失败", err)
			return
		}
	} else {
		APIReturn(c, 500, "类型不存在", errors.New("类型不存在"))
		return
	}

	APIReturn(c, 200, "添加任务成功", nil)
}

func submitTask(ob *md.Task) error {
	ob.ID = ksuid.New().String()
	ob.PauseNum = 0
	if len(ob.RS) > 0 {
		if err := md.RemoteStorageToTaskFunc(ob.ID, ob.RS); err != nil {
			return err
		}
	}
	if err := ob.Add(); err != nil {
		return err
	}
	toolbox.AddTask(ob.ID, toolbox.NewTask(ob.ID, ob.Crontab, jobs.Jobs(ob.ID)))
	return nil
}

// Update ...
func UpdateTask(c *gin.Context) {
	ob := new(md.Task)
	if err := c.Bind(ob); err != nil {
		APIReturn(c, 500, "更新任务失败1", err.Error())
		return
	}

	if ob.Crontab == "" {
		APIReturn(c, 500, "计划任务必填", errors.New("计划任务必填"))
		return
	}
	if ob.DBType != "file" {
		if err := db.ConnectTest(ob); err != nil {
			APIReturn(c, 500, "更新任务失败4", err.Error())
			return
		}
	} else if ob.DBType == "file" {
		if bol, err := utils.PathExists(ob.DBpath); err != nil {
			APIReturn(c, 500, "file not found 1", err.Error())
			return
		} else {
			if !bol {
				APIReturn(c, 500, "file not found", errors.New("file not found"))
				return
			}
		}
	} else {
		APIReturn(c, 500, "没有此类型数据", errors.New("dbtype not found"))
		return
	}

	if len(ob.RS) > 0 {
		if err := md.RemoteStorageToTaskFunc(ob.ID, ob.RS); err != nil {
			APIReturn(c, 500, "更新任务失败6", err.Error())
			return
		}
	}

	if err := ob.Update(); err != nil {
		APIReturn(c, 500, "更新任务失败2", err.Error())
		return
	}
	toolbox.StopTask()
	toolbox.AddTask(ob.ID, toolbox.NewTask(ob.ID, ob.Crontab, jobs.Jobs(ob.ID)))
	toolbox.StartTask()
	APIReturn(c, 200, "更新任务成功", ob.ID)
}

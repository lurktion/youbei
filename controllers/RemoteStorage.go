package controllers

import (
	"errors"
	"strconv"

	md "youbei/models"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

// Ftpadd ...
func Ftpadd(c *gin.Context) {
	ob := new(md.RemoteStorage)
	if err := c.Bind(ob); err != nil {
		APIReturn(c, 500, "数据解析失败", err.Error())
		return
	}

	ob.ID = ksuid.New().String()
	if err := ob.Add(); err != nil {
		APIReturn(c, 500, "添加失败", err.Error())
		return
	}

	APIReturn(c, 200, "添加成功", ob)
}

// Ftpdelete ...
func Ftpdelete(c *gin.Context) {
	id := c.Param("id")
	ob := new(md.RemoteStorage)
	ob.ID = id
	if err := ob.Delete(); err != nil {
		APIReturn(c, 500, "删除失败", err.Error())
		return
	}

	APIReturn(c, 200, "删除成功", nil)
}

// Ftplist ...
func Ftplist(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("count"))
	RemoteStorages := []md.RemoteStorage{}
	xs := md.Localdb().Desc("created")
	if limit > 0 {
		xs = xs.Limit(limit, limit*(page-1))
	}
	if err := xs.Find(&RemoteStorages); err != nil {
		APIReturn(c, 500, "获取列表失败", err.Error())
		return
	}

	ftps := md.RemoteStorage{}
	title, err := md.Localdb().Count(&ftps)
	if err != nil {
		APIReturn(c, 500, "获取总数失败", err.Error())
		return
	}

	rep := map[string]interface{}{"count": title, "data": RemoteStorages}

	APIReturn(c, 200, "获取列表成功", &rep)
}

// Ftpfind 查询远程存储 单
func Ftpfind(c *gin.Context) {
	rs := new(md.RemoteStorage)
	if bol, err := md.Localdb().ID(c.Param("id")).Get(rs); err != nil {
		APIReturn(c, 500, "找不到该远程存储", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "找不到该远程存储", errors.New("not found"))
			return
		}
	}

	APIReturn(c, 200, "获取成功", rs)
}

// Ftpupdate 查询远程存储 单
func Ftpupdate(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		APIReturn(c, 500, "id不能为空", errors.New("id不能为空"))
		return
	}
	rs := new(md.RemoteStorage)
	if err := c.Bind(rs); err != nil {
		APIReturn(c, 500, "解析失败", err.Error())
		return
	}

	if err := rs.Update(); err != nil {
		APIReturn(c, 500, "修改失败", err.Error())
		return
	}

	APIReturn(c, 200, "修改成功", nil)
}

// Rloglist 查询远程存储 单
func Rloglist(c *gin.Context) {
	type Query struct {
		Page  int `form:"page"`
		Count int `form:"count"`
	}
	query := new(Query)
	c.Bind(query)
	rlogs, err := md.RemoteSendLogFindAll("", query.Page, query.Count)
	if err != nil {
		APIReturn(c, 500, "获取列表失败", err.Error())
		return
	}

	title, err := md.Localdb().Count(new(md.Rlog))
	if err != nil {
		APIReturn(c, 500, "获取总数失败", err.Error())
		return
	}

	rep := map[string]interface{}{"count": title, "data": rlogs}
	APIReturn(c, 200, "获取列表成功", &rep)
}

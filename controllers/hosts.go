package controllers

import (
	"errors"
	"net"

	md "youbei/models"

	"github.com/gin-gonic/gin"
)

func GetLocal(c *gin.Context) {
	str := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		APIReturn(c, 500, "failed", err.Error())
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok {
			if ipnet.IP.To4() != nil {
				str = append(str, ipnet.IP.String())
			}
		}
	}
	str = append(str, "localhost")
	APIReturn(c, 200, "success", str)
}

func HostAdd(c *gin.Context) {
	ob := new(md.Host)
	if err := c.Bind(ob); err != nil {
		APIReturn(c, 500, "解析录入失败", err.Error())
		return
	}
	if err := ob.Add(); err != nil {
		APIReturn(c, 500, "录入主机失败", err.Error())
		return
	}

	APIReturn(c, 200, "录入成功", nil)
}

func HostUpdate(c *gin.Context) {
	id := c.Param("id")
	ob := new(md.Host)
	if err := c.Bind(ob); err != nil {
		APIReturn(c, 500, "解析修改失败", err.Error())
		return
	}
	ob.ID = id
	if err := ob.Update(); err != nil {
		APIReturn(c, 500, "修改主机失败", err.Error())
		return
	}
	APIReturn(c, 200, "修改成功", nil)
}

func HostDelete(c *gin.Context) {
	id := c.Param("id")
	ob := new(md.Host)
	ob.ID = id
	if err := ob.Delete(); err != nil {
		APIReturn(c, 500, "删除主机失败", err.Error())
		return
	}
	APIReturn(c, 200, "删除成功", nil)
}

func HostsGet(c *gin.Context) {
	hosts := make([]md.Host, 0)
	err := md.Localdb().Find(&hosts)
	if err != nil {
		APIReturn(c, 500, "获取数据失败", err.Error())
		return
	}
	APIReturn(c, 200, "获取数据成功", hosts)
}

func HostGet(c *gin.Context) {
	id := c.Param("id")
	host := new(md.Host)
	host.ID = id
	if bol, err := md.Localdb().Get(host); err != nil {
		APIReturn(c, 500, "获取数据失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "数据不存在", errors.New("数据不存在"))
			return
		}
	}
	APIReturn(c, 200, "获取数据成功", host)
}

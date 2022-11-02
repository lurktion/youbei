package controllers

import (
	"strings"
	md "youbei/models"

	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context) {
	id := c.Param("id")
	log := new(md.Log)
	log.ID = id
	if err := log.Get(); err != nil {
		APIReturn(c, 500, "查询失败", err)
		return
	}
	arr := strings.Split(log.Localfilepath, "/")
	filename := arr[len(arr)-1]
	c.FileAttachment(log.Localfilepath, filename)
	return
}

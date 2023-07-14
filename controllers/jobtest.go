package controllers

import (
	jobs "youbei/utils/jobs"

	"github.com/gin-gonic/gin"
)

func RunJob(c *gin.Context) {
	if err := jobs.Backup(c.Param("id"), true); err != nil {
		APIReturn(c, 500, "失败："+err.Error(), err.Error())
		return
	}
	APIReturn(c, 200, "执行成功", nil)
}

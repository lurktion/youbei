package controllers

import (
	"errors"
	"strings"
	md "youbei/models"

	"github.com/gin-gonic/gin"
)

func GetCmds(c *gin.Context) {
	cmds, err := md.GetCmds()
	if err != nil {
		APIReturn(c, 500, "获取列表失败", err)
	}

	APIReturn(c, 200, "成功", cmds)
}

func GetCmd(c *gin.Context) {
	id := c.Query("id")
	cmd := md.SystemBackupCmdPath{}
	if bol, err := md.Localdb().ID(id).Get(&cmd); err != nil {
		APIReturn(c, 500, "获取失败", err)
		return
	} else {
		if !bol {
			APIReturn(c, 500, "不存在", errors.New("不存在"))
			return
		}
	}

	APIReturn(c, 200, "成功", cmd)
}

func UpdateCmd(c *gin.Context) {
	cmd := md.SystemBackupCmdPath{}
	if err := c.Bind(&cmd); err != nil {
		APIReturn(c, 500, "解析失败", err)
		return
	}
	cmd.Path = strings.ReplaceAll(cmd.Path, "\\", "/")
	if _, err := md.Localdb().ID(cmd.ID).Cols("path").Update(&cmd); err != nil {
		APIReturn(c, 500, "修改失败", err)
		return
	}

	APIReturn(c, 200, "成功", cmd)
}

package db

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
	md "youbei/models"
	utils "youbei/utils"
	Zips "youbei/utils/zip"
)

func MysqlCmdDump(info *md.Task, dist string) error {
	out, fileerr := os.OpenFile(dist, os.O_CREATE|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if fileerr != nil {
		return fileerr
	}
	defer out.Close()
	cmds := md.SystemBackupCmdPath{}
	if bol, err := md.Localdb().Where("id=? and status=0", info.Cmds).Get(&cmds); err != nil {
		return err
	} else {
		if !bol {
			return errors.New("dbtype not found2")
		}
	}

	var cmd *exec.Cmd
	var cmdstr string
	var cmsarr = []string{}
	var cmdbase string

	sysos := runtime.GOOS
	if sysos == "windows" {
		cmdbase = "powershell"
		cmdstr = Fmtpath(cmds.Path) + " " + fmt.Sprintf(`--complete-insert --skip-lock-tables --skip-comments --compact --add-drop-table --host %s --port %s -u%s --password=%s --default-character-set %s -B %s`, info.Host, info.Port, info.User, info.Password, info.Char, info.DBname)
		cmsarr = strings.Split(cmdstr, " ")
		cmd = exec.Command(cmdbase, cmsarr...)
	} else {
		cmdstr = fmt.Sprintf(`--complete-insert --skip-lock-tables --skip-comments --compact --add-drop-table --host %s --port %s -u%s --password=%s --default-character-set %s -B %s`, info.Host, info.Port, info.User, info.Password, info.Char, info.DBname)
		cmsarr = strings.Split(cmdstr, " ")
		cmd = exec.Command(Fmtpath(cmds.Path), cmsarr...)
	}

	var stderr bytes.Buffer
	cmd.Stdout = out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errors.New(err.Error() + ":" + stderr.String())
	}

	out.Close()

	return nil
}

type Progress struct {
	TotalBytes   int64
	WrittenBytes int64
}

func MysqlCmdRecovery(loginfo *md.Log) error {
	taskinfo := md.Task{}

	if bol, err := md.Localdb().ID(loginfo.Tid).Get(&taskinfo); err != nil {
		return err
	} else {
		if !bol {
			return errors.New("任务不存在,或已被删除，无法还原")
		}
	}

	// if taskinfo.DBType != loginfo.Dbtype {
	// 	return errors.New("数据库类型不匹配(" + taskinfo.DBType + "!=" + loginfo.Dbtype + ")")
	// }

	if taskinfo.ID != loginfo.Tid {
		return errors.New("数据库不匹配，无法执行还原操作")
	}

	if bol, err := utils.PathExists(loginfo.Localfilepath); err != nil {
		return err
	} else {
		if !bol {
			return errors.New("文件不存在")
		}
	}

	cmds := md.SystemBackupCmdPath{}
	if bol, err := md.Localdb().Where("dbtype=? and status=0 and recovery=1", taskinfo.DBType).Get(&cmds); err != nil {
		return err
	} else {
		if !bol {
			return errors.New("dbtype not found2")
		}
	}

	var cmd *exec.Cmd
	var cmdstr string
	var cmsarr = []string{}
	var cmdbase string

	sysos := runtime.GOOS
	if sysos == "windows" {
		cmdbase = "powershell"
		cmdstr = Fmtpath(cmds.Path) + " " + fmt.Sprintf(`--host %s --port %s -u%s --password=%s -B %s`, taskinfo.Host, taskinfo.Port, taskinfo.User, taskinfo.Password, taskinfo.DBname)
		cmsarr = strings.Split(cmdstr, " ")
		cmd = exec.Command(cmdbase, cmsarr...)
	} else {
		cmdstr = fmt.Sprintf(`--host %s --port %s -u%s --password=%s -B %s`, taskinfo.Host, taskinfo.Port, taskinfo.User, taskinfo.Password, taskinfo.DBname)
		cmsarr = strings.Split(cmdstr, " ")
		cmd = exec.Command(Fmtpath(cmds.Path), cmsarr...)
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	progress := new(int64)
	done := make(chan bool) // 创建一个通道来接收 Unzip 函数是否完成的信号
	go func() {
		for {
			select {
			case <-done: // 如果收到 done 信号，就退出循环
				return
			default:
				currentProgress := atomic.LoadInt64(progress)
				updateData := md.Log{}
				updateData.RecoveryProgress = currentProgress
				md.Localdb().ID(loginfo.ID).Cols("recoveryProgress").Update(&updateData)

				// 如果进度达到100%，则退出循环
				if currentProgress == 100 {
					return
				}

				// 每秒读取并打印一次进度
				time.Sleep(time.Second)
			}
		}
	}()

	if ziperr := Zips.Unzip(loginfo.Localfilepath, loginfo.Password, stdin, progress); ziperr != nil {
		close(done) // 如果 Unzip 函数完成（无论是否成功），就关闭 done 通道
		return ziperr
	}
	close(done) // 如果 Unzip 函数成功，也关闭 done 通道
	return nil
}

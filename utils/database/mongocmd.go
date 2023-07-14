package db

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	md "youbei/models"
)

func MONGOCmdDump(info *md.Task, dist string) error {
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
		cmdstr = Fmtpath(cmds.Path) + " " + fmt.Sprintf(`-h %s --port %s --authenticationDatabase %s --username %s --password %s -d %s -o %s`, info.Host, info.Port, info.DefDbname, info.User, info.Password, info.DBname, dist)
		cmsarr = strings.Split(cmdstr, " ")
		cmd = exec.Command(cmdbase, cmsarr...)
	} else {
		cmdstr = fmt.Sprintf(`-h %s --port %s --authenticationDatabase %s --username %s --password %s -d %s -o %s`, info.Host, info.Port, info.DefDbname, info.User, info.Password, info.DBname, dist)
		cmsarr = strings.Split(cmdstr, " ")
		cmd = exec.Command(Fmtpath(cmds.Path), cmsarr...)
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errors.New(err.Error() + ":" + stderr.String())
	}

	return nil
}

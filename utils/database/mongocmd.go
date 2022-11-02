package db

import (
	"errors"
	"fmt"
	"os/exec"
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
	cmdstr := fmt.Sprintf(`-h %s --port %s --username %s --password %s -d %s -o %s`, info.Host, info.Port, info.User, info.Password, info.DBname, dist)
	cmsarr := strings.Split(cmdstr, " ")
	cmd := exec.Command(Fmtpath(cmds.Path), cmsarr...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

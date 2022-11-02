package db

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	md "youbei/models"
)

func PGSQLCmdDump(info *md.Task, dist string) error {
	cmds := md.SystemBackupCmdPath{}
	if bol, err := md.Localdb().Where("id=? and status=0", info.Cmds).Get(&cmds); err != nil {
		return err
	} else {
		if !bol {
			return errors.New("dbtype not found2")
		}
	}
	cmdstr := fmt.Sprintf("-h %s --port %s -U %s -F c -b -v -f %s %s", info.Host, info.Port, info.User, dist, info.DBname)
	cmsarr := strings.Split(cmdstr, " ")
	cmd := exec.Command(Fmtpath(cmds.Path), cmsarr...)

	os.Setenv("PGPASSWORD", info.Password)
	if err := cmd.Run(); err != nil {
		return err
	}
	os.Unsetenv("PGPASSWORD")
	return nil
}

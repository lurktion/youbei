package db

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	md "youbei/models"
)

func MysqlCmdDump(info *md.Task, dist string) error {
	f, fileerr := os.OpenFile(dist, os.O_CREATE|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if fileerr != nil {
		return fileerr
	}
	defer f.Close()
	cmds := md.SystemBackupCmdPath{}
	if bol, err := md.Localdb().Where("id=? and status=0", info.Cmds).Get(&cmds); err != nil {
		return err
	} else {
		if !bol {
			return errors.New("dbtype not found2")
		}
	}
	cmdstr := fmt.Sprintf("--complete-insert --skip-comments --compact --add-drop-table --host %s --port %s -u%s -p%s --databases %s --default-character-set %s ", info.Host, info.Port, info.User, info.Password, info.DBname, info.Char)
	cmsarr := strings.Split(cmdstr, " ")
	cmd := exec.Command(cmds.Path, cmsarr...)

	cmd.Stdout = f

	cmd.Run()

	f.Close()
	return nil
}

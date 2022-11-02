package db

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	md "youbei/models"
)

func MSSQLCmdDump(info *md.Task, dist string) error {
	cmds := md.SystemBackupCmdPath{}
	if bol, err := md.Localdb().Where("id=? and status=0", info.Cmds).Get(&cmds); err != nil {
		return err
	} else {
		if !bol {
			return errors.New("dbtype not found2")
		}
	}
	arr := strings.Split(dist, "/")
	cmdstr := Fmtpath(cmds.Path) + " " + fmt.Sprintf(`-S %s -U%s -P'%s' -Q "BACKUP DATABASE %s TO DISK='%s'"`, info.Host, info.User, info.Password, info.DBname, info.SavePath+"/"+arr[len(arr)-1])
	cmsarr := strings.Split(cmdstr, " ")

	cmd := exec.Command("powershell", cmsarr...)
	fmt.Println(cmd.Args)
	cmd.Run()
	return nil
}

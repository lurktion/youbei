package db

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

	var cmd *exec.Cmd
	var cmdstr string
	var cmsarr = []string{}
	var cmdbase string

	os.Setenv("SQLCMDPASSWORD", info.Password)

	sysos := runtime.GOOS
	if sysos == "windows" {
		cmdbase = "powershell"
		cmdstr = Fmtpath(cmds.Path) + " " + fmt.Sprintf(`-S %s -U%s -Q "BACKUP DATABASE %s TO DISK='%s'"`, info.Host, info.User, info.DBname, info.SavePath+"/"+arr[len(arr)-1])
		cmsarr = strings.Split(cmdstr, " ")
		cmd = exec.Command(cmdbase, cmsarr...)
	} else {
		cmdstr = fmt.Sprintf(`-S %s -U%s -Q "BACKUP DATABASE %s TO DISK='%s'"`, info.Host, info.User, info.DBname, info.SavePath+"/"+arr[len(arr)-1])
		cmsarr = strings.Split(cmdstr, " ")
		cmd = exec.Command(Fmtpath(cmds.Path), cmsarr...)
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	os.Setenv("SQLCMDPASSWORD", info.Password)
	if err := cmd.Run(); err != nil {
		return errors.New(err.Error() + ":" + stderr.String())
	}
	os.Unsetenv("SQLCMDPASSWORD")
	return nil
}

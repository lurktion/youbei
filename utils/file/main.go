package filedump

import (
	"strings"
	"time"

	Zipz "youbei/utils/zip"
)

// FileDump ...
func FileDump(path, backupfilePath, zippwd string) (string, error) {
	nowtime := time.Now().Format("2006-01-02_15-04-05")
	dbstr := strings.Split(path, "/")
	distzp := "localhost_filebackup_" + dbstr[len(dbstr)-1] + "_" + nowtime + ".zip"
	err := Zipz.Zip(path, backupfilePath+"/"+distzp, zippwd)
	if err != nil {
		return "", err
	}
	return backupfilePath + "/" + distzp, err
}

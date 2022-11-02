package db

import (
	"os"
	"strings"
	"time"
	md "youbei/models"
	Zipz "youbei/utils/zip"
)

//SQLDumpZip ...
func SQLDumpZip(backupfilePath, zipPath, zipPassword string) (string, error) {
	err := Zipz.Zip(backupfilePath, zipPath, zipPassword)
	if err != nil {
		return zipPath, err
	}
	err = os.RemoveAll(backupfilePath)
	return zipPath, err
}

func FmtFilename(info *md.Task) (string, string) {
	nowtime := time.Now().Format("2006-01-02_15-04-05")
	dist := info.SavePath + "/" + info.Host + "_" + info.DBType + "_" + info.DBname + "_" + "_" + nowtime + ".sql"
	distzip := info.SavePath + "/" + info.Host + "_" + info.DBType + "_" + info.DBname + "_" + "_" + nowtime + ".zip"
	return dist, distzip
}

func Fmtpath(str string) string {
	arr := strings.Split(str, "/")
	newarr := []string{}
	for _, v := range arr {
		if bol := strings.ContainsAny(v, " "); bol {
			newarr = append(newarr, `"`+v+`"`)
		} else {
			newarr = append(newarr, v)
		}
	}
	return strings.Join(newarr, "/")
}

package md

import (
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/segmentio/ksuid"

	// gosqlite3
	_ "github.com/mattn/go-sqlite3"
)

type SystemBackupCmdPath struct {
	ID       string `json:"id" xorm:"pk notnull unique 'id'"`
	Name     string `json:"name" xorm:"'name'"`
	Path     string `json:"path" xorm:"default('') 'path'"`
	Def      int    `json:"def" xorm:"default(0) 'def'"`
	Deleted  int64  `json:"deleted" xorm:"'deleted'"`
	DBtype   string `json:"dbtype" xorm:"'dbtype'"`
	Created  int64  `json:"created" xorm:"'created'"`
	Status   int    `json:"status" xorm:"'status'"`
	Recovery int    `json:"recovery" xorm:"'recovery'"`
	Commit   int    `json:"commit" xorm:"'commit'"`
}

func initSystemBackupCmdPath() {
	names := []string{`mysqldump`, `pg_dump`, `sqlcmd`, `mongodump`, `mysql`}
	for _, v := range names {
		dbtype := ""
		commit := 0
		recovery := 0
		if v == "mysqldump" {
			dbtype = "mysql"
			commit = 0
			recovery = 0
		} else if v == "pg_dump" {
			dbtype = "postgres"
			commit = 0
			recovery = 0
		} else if v == "sqlcmd" {
			dbtype = "mssql"
			commit = 1
			recovery = 0
		} else if v == "mongodump" {
			dbtype = "mongodb"
			commit = 0
			recovery = 0
		} else if v == "mysql" {
			dbtype = "mysql"
			commit = 0
			recovery = 1
		}
		cmdpath := SystemBackupCmdPath{}
		if bol, err := localdb.Where("`name`=? and `def`=0", v).Get(&cmdpath); err != nil {
			panic(err.Error())
		} else {
			if !bol {
				newcmd := SystemBackupCmdPath{}
				newcmd.ID = ksuid.New().String()
				newcmd.Created = time.Now().Unix()
				newcmd.Def = 0
				newcmd.Name = v
				newcmd.DBtype = dbtype
				newcmd.Commit = commit
				newcmd.Recovery = recovery
				if paths, errpath := exec.LookPath(v); errpath != nil {
					newcmd.Path = ""
				} else {
					newcmd.Path = strings.ReplaceAll(paths, "\\", "/")
				}
				if _, errins := localdb.Insert(&newcmd); errins != nil {
					panic(errins.Error())
				}
			} else {
				if cmdpath.Path == "" {
					if paths, errpath := exec.LookPath(v); errpath != nil {
						cmdpath.Path = ""
					} else {
						cmdpath.Path = strings.ReplaceAll(paths, "\\", "/")
					}
				}
				cmdpath.DBtype = dbtype
				cmdpath.Commit = commit
				cmdpath.Recovery = recovery
				if _, errins := localdb.ID(cmdpath.ID).Cols("path", "dbtype", "commit", "recovery").Update(&cmdpath); errins != nil {
					panic(errins.Error())
				}
			}
		}
	}
	CheckCmdsStatus()
}

func CheckCmdsStatus() error {
	cmds := []SystemBackupCmdPath{}
	if err := localdb.Find(&cmds); err != nil {
		return err
	}
	for _, v := range cmds {
		if v.Path != "" {
			if bol := IsFile(v.Path); bol {
				v.Status = 0
			} else {
				v.Status = 1
			}
			localdb.ID(v.ID).Cols("status").Update(&v)
		}
	}
	return nil
}

func GetCmds() ([]SystemBackupCmdPath, error) {
	cmds := []SystemBackupCmdPath{}
	if err := localdb.Find(&cmds); err != nil {
		return cmds, err
	}
	for k, v := range cmds {
		if bol := IsFile(v.Path); bol {
			cmds[k].Status = 0
		} else {
			cmds[k].Status = 1
		}
	}
	return cmds, nil
}

func IsFile(path string) bool {
	if f, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	} else {
		if f.IsDir() {
			return false
		}
	}
	return true
}

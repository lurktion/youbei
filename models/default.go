package md

import (
	"os"
	"runtime"
	"strings"
	"time"
	"youbei/utils"

	// gosqlite3

	"gitee.com/countpoison_admin/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/segmentio/ksuid"
)

var localdb *xorm.Engine

var DBdir string

// Init 初始化数据库
func init() {
	sysos := runtime.GOOS
	if sysos == "windows" {
		DBdir = os.Getenv("APPDATA") + "/youbei"
	} else {
		DBdir = "/usr/local/youbei"
	}

	if err := os.MkdirAll(DBdir, os.ModePerm); err != nil {
		panic(err)
	}

	var err error
	localdb, err = xorm.NewEngine("sqlite3", DBdir+"/local.db")
	if err != nil {
		panic(err)
	}

	if err = localdb.Sync2(
		new(MyUniqId),
		new(Task),
		new(Log),
		new(User),
		new(RemoteStorage),
		new(RemoteStorageToTask),
		new(Rlog), new(Yserver),
		new(YserverFile),
		new(YserverPacket),
		new(YsUploadFile),
		new(YsPacket),
		new(MailServer),
		new(Host),
		new(MailSend),
		new(SshTask),
		new(SystemBackupCmdPath)); err != nil {
		panic(err)
	}
	UserInit("admin", "admin")
	InitYservefunc()
	MailServerInit()
	idfortasktable()
	initSystemBackupCmdPath()
	if err = localdb.RegisterSqlTemplate(xorm.Pongo2(".", ".stpl")); err != nil {
		panic(err.Error())
	}
	initSqlTempReg()
}

// Localdb 返回数据库对象
func Localdb() *xorm.Engine {
	return localdb
}

func idfortasktable() {
	tasks := []Task{}
	if err := localdb.Find(&tasks); err != nil {
		panic(err.Error())
	}
	for _, v := range tasks {
		if v.Name == "" {
			if name, err := CreateId("TASK", strings.ToUpper(v.DBType)); err != nil {
				panic(err.Error())
			} else {
				v.Name = name
				if _, updateerr := localdb.ID(v.ID).Cols("name").Update(&v); updateerr != nil {
					panic(err.Error())
				}
			}
		}
	}
}

// NowTime 返回当前时间戳
func NowTime() int64 {
	return time.Now().Unix()
}

// UserInit ...
func UserInit(username string, password string) error {
	Countuser := new(User)
	if total, err := localdb.Count(Countuser); err != nil {
		return err
	} else {
		if total > 0 {
			return nil
		}
	}

	user := User{}
	user.Name = username
	if bol, err := localdb.Get(&user); err != nil {
		return err
	} else {
		if bol {
			return nil
		}
	}

	user.ID = ksuid.New().String()
	user.Created = time.Now().Unix()
	user.Passwrod = utils.Md5V(password)
	if _, err := localdb.Insert(&user); err != nil {
		return err
	}
	return nil
}

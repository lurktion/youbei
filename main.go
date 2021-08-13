package main

import (
	"errors"
	"flag"
	"fmt"

	md "gitee.com/countpoison/youbei/models"
	_ "gitee.com/countpoison/youbei/routers"
	"gitee.com/countpoison/youbei/utils"
	db "gitee.com/countpoison/youbei/utils/database"
	"gitee.com/countpoison/youbei/utils/jobs"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/adapter/toolbox"
	"github.com/beego/beego/v2/core/logs"
	//"github.com/beego/beego/v2/core/logs"
)

// //go:embed static
// var embededFile embed.FS

func main() {
	//logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	//os.MkdirAll("logs", 0666)
	httpport, err := beego.AppConfig.Int("httpport")
	if err != nil {
		//logs.Error(err)
		panic(err)
	}
	init_set_user := flag.String("init_set_user", "admin", "Initialize the user, take effect when the user table is empty")
	init_set_password := flag.String("init_set_password", "admin", "Initialize password, take effect when user table is empty")
	port := flag.Int("port", httpport, "Set http listening port, lasting effect, please modify the configuration file")
	flag.Parse()
	md.Init("data", *init_set_user, *init_set_password)
	ts, err := md.All()
	if err != nil {
		//logs.Error(err)
		panic(err)
	}
	if len(ts) > 0 {
		for _, ob := range ts {
			if ob.DBType == "mysql" {
				err = db.MysqlConnectTest(ob.Host, ob.Port, ob.DBname, ob.User, ob.Password, ob.Char)
			} else if ob.DBType == "mssql" {
				err = db.MssqlConnectTest(ob.Host, ob.DBname, ob.User, ob.Password)
			} else if ob.DBType == "sqlite" {
				err = db.SqliteConnectTest(ob.DBpath)
			} else if ob.DBType == "postgres" {
				err = db.PostgresConnectTest(ob.Host, ob.Port, ob.DBname, ob.User, ob.Password)
			} else if ob.DBType == "file" {
				bol, errs := utils.PathExists(ob.DBpath)
				err = errs
				if !bol {
					logs.Error(err)
					err = errors.New(ob.DBpath + " not found")
				}
			} else {
				logs.Debug(err)
				fmt.Println("dbtype not found")
			}
			if err == nil && ob.Crontab != "" {
				toolbox.AddTask(ob.ID, toolbox.NewTask(ob.ID, ob.Crontab, jobs.Jobs(ob.ID)))
			}
		}
	}
	sshtasks := []md.SshTask{}
	if err := md.Localdb().Find(&sshtasks); err != nil {
		//logs.Error(err)
		panic(err.Error())
	}
	for _, sshtask := range sshtasks {
		toolbox.AddTask(sshtask.ID, toolbox.NewTask(sshtask.ID, sshtask.Crontab, jobs.SshJobs(sshtask.ID)))
	}
	toolbox.StartTask()
	if *port != httpport {
		beego.BConfig.Listen.HTTPPort = *port
	}
	//useOS := len(os.Args) > 1 && os.Args[1] == "live"
	// beego.Handler("/", http.FileServer(getFileSystem(useOS)))
	//beego.Handler("/", http.FileServer(http.FS(embededFile)))
	beego.Run()
}

// func getFileSystem(useOS bool) http.FileSystem {
// 	if useOS {
// 		log.Print("using live mode")
// 		return http.FS(os.DirFS("static"))
// 	}

// 	log.Print("using embed mode")

// 	fsys, err := fs.Sub(embededFile, "static")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return http.FS(fsys)
// }

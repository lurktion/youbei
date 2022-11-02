package main

import (
	"errors"
	"net/http"

	"youbei/static"
	"youbei/utils"
	db "youbei/utils/database"
	"youbei/utils/jobs"

	"youbei/controllers"

	md "youbei/models"

	"github.com/beego/beego/toolbox"
	"github.com/gin-gonic/gin"
)

func main() {
	ts, err := md.All()
	if err != nil {
		panic(err)
	}

	if len(ts) > 0 {
		for _, ob := range ts {
			if ob.DBType != "file" {
				err = db.ConnectTest(&ob)
			} else {
				bol, errs := utils.PathExists(ob.DBpath)
				err = errs
				if !bol {
					err = errors.New(ob.DBpath + " not found")
				}
			}
			if err == nil && ob.Crontab != "" {
				toolbox.AddTask(ob.ID, toolbox.NewTask(ob.ID, ob.Crontab, jobs.Jobs(ob.ID)))
			}
		}
	}
	sshtasks := []md.SshTask{}
	if err := md.Localdb().Find(&sshtasks); err != nil {
		panic(err.Error())
	}
	for _, sshtask := range sshtasks {
		toolbox.AddTask(sshtask.ID, toolbox.NewTask(sshtask.ID, sshtask.Crontab, jobs.SshJobs(sshtask.ID)))
	}
	toolbox.StartTask()

	r := gin.Default()
	r.Use(controllers.Cors())
	r.StaticFS("/ui", http.FS(static.Static))
	r.StaticFS("/static", http.FS(static.Static))

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/ui")
	})

	r.GET("/downloadfile/:id", controllers.DownloadFile)
	r.GET("/login/:name/:password", controllers.Userlogin)

	r.POST("/upload/packet/:id/:offset", controllers.Uploadpacket)
	r.POST("/upload/done/:id", controllers.UploadpacketDone)
	r.POST("/upload/file/:id", controllers.UploadFile)

	r.GET("/service/upload", controllers.Yserverlist)
	r.PUT("/service/upload/:id", controllers.EnableServer)
	r.DELETE("/service/upload/:id", controllers.DisableServer)

	api := r.Group("/api")
	api.Use(controllers.Prepare())
	{
		api.GET("/tasks", controllers.GetTasks)
		api.GET("/task/:id", controllers.GetTask)
		api.POST("/task", controllers.AddTask)
		api.PUT("/task/:id", controllers.UpdateTask)
		api.DELETE("/task/:id", controllers.DeleteTask)

		api.GET("/logs", controllers.Loglist)
		api.GET("/log/:id", controllers.ShowLog)

		api.GET("/rlogs", controllers.Rloglist)
		api.GET("/rlog/:id", controllers.ShowrLog)
		api.GET("/uploadlogs", controllers.Uploadlogs)
		api.GET("/uploadlog/:id", controllers.GetYserverLog)
		api.GET("/ftps", controllers.Ftplist)
		api.GET("/ftp/:id", controllers.Ftpfind)
		api.POST("/ftp", controllers.Ftpadd)
		api.DELETE("/ftp/:id", controllers.Ftpdelete)
		api.PUT("/ftp/:id", controllers.Ftpupdate)

		api.PUT("/runjob/:id", controllers.RunJob)
		api.PUT("/runsshjob/:id", controllers.RunSshJob)

		api.GET("/dirlist", controllers.DirList)

		api.GET("/hosts", controllers.HostsGet)
		api.GET("/host/:id", controllers.HostGet)
		api.PUT("/host/:id", controllers.HostUpdate)
		api.DELETE("/host/:id", controllers.HostDelete)
		api.POST("/host", controllers.HostAdd)

		api.GET("/users", controllers.UserList)
		api.GET("/user/:id", controllers.GetUser)
		api.POST("/user", controllers.AddUser)
		api.PUT("/user/:id", controllers.EditUser)
		api.DELETE("/user/:id", controllers.DeleteUser)

		api.PUT("/pchange/:id", controllers.Userchangepwd)

		api.POST("/mailserver", controllers.MailTest)
		api.PUT("/mailserver", controllers.MailServerUpdate)
		api.GET("/mailserver", controllers.GetMail)

		api.PUT("/connecthost/:id", controllers.ConnectHost)

		api.GET("/sshtasks", controllers.SshtaskList)
		api.POST("/sshtask", controllers.AddSshtask)
		api.PUT("/sshtask/:id", controllers.UpdateSshtask)
		api.DELETE("/sshtask/:id", controllers.DeleteSshtask)
		api.GET("/sshtask/:id", controllers.GetSshtask)

		api.GET("/getdbinfo", controllers.DBGet)
		api.GET("/getcmds", controllers.GetCmds)
		api.GET("/getcmd", controllers.GetCmd)
		api.PUT("/getcmd", controllers.UpdateCmd)

		api.GET("/sysinfo", controllers.Sysinfo)
		api.GET("/dashboardinfo", controllers.DashBoardInfo)
		api.GET("/getlocal", controllers.GetLocal)

	}

	r.Run()
}

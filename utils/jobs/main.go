package jobs

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	md "youbei/models"
	dbtools "youbei/utils/database"
	filedump "youbei/utils/file"
	"youbei/utils/mail"
	rs "youbei/utils/rs"

	"gitee.com/countpoison_admin/xorm"
	"github.com/segmentio/ksuid"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

//ExecBackup ...
func ExecBackup(info *md.Task) (string, error) {
	dist, distzip := dbtools.FmtFilename(info)

	zippwd := ""
	if info.Enablezippwd == 1 {
		zippwd = info.Zippwd
	}

	if info.Cmds == "" {
		var db *xorm.Engine
		var err error
		if info.DBType == "mysql" {
			sqlstr := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s", info.User, info.Password, info.Host, info.Port, info.DBname, info.Char)
			db, err = xorm.NewEngine("mysql", sqlstr)
		} else if info.DBType == "mssql" {
			sqlstr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", info.Host, info.User, info.Password, info.Port, info.DBname)
			db, err = xorm.NewEngine("mssql", sqlstr)
		} else if info.DBType == "sqlite" {
			db, err = xorm.NewEngine("sqlite3", info.DBpath)
		} else if info.DBType == "postgres" {
			sqlstr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", info.Host, info.Port, info.User, info.Password, info.DBname)
			db, err = xorm.NewEngine("postgres", sqlstr)
		} else if info.DBType == "file" {
			return filedump.FileDump(info.DBpath, info.SavePath, zippwd)
		} else {
			return "", errors.New("dbtype not found")
		}

		if err != nil {
			return "", err
		}
		defer db.Close()
		if err = db.Ping(); err != nil {
			return "", err
		}
		if err = db.DumpAllToFile(dist); err != nil {
			return "", err
		}

	} else {
		if info.DBType == "mysql" {
			if err := dbtools.MysqlCmdDump(info, dist); err != nil {
				return "", err
			}
		} else if info.DBType == "postgres" {
			if err := dbtools.PGSQLCmdDump(info, dist); err != nil {
				return "", err
			}
		} else if info.DBType == "mssql" {
			if err := dbtools.MSSQLCmdDump(info, dist); err != nil {
				return "", err
			}
			if info.Types != 0 {
				return "", nil
			}
		} else if info.DBType == "mongodb" {
			if err := dbtools.MONGOCmdDump(info, dist); err != nil {
				return "", err
			}
		} else {
			return "", errors.New("dbtype not found2")
		}

	}

	return dbtools.SQLDumpZip(dist, distzip, zippwd)
}

//ExecRemote ...
func ExecRemote(v md.RemoteStorage, rlid string, lid string, Localfilepath string) error {
	var err error
	yuancheng := rs.NewRS(v.Host, v.Port, v.Username, v.Password, Localfilepath, v.Path)
	if v.Types == "ftp" {
		err = yuancheng.FtpUpload()
	} else if v.Types == "sftp" {
		f, err := yuancheng.NewSftp()
		if err == nil {
			err = f.Upload()
		}
	} else if v.Types == "Yserver" {
		filepackets, err := yuancheng.ReadBigFile()
		if err != nil {
			return err
		}
		ysuploadfile := new(md.YsUploadFile)
		ysuploadfile.ID = ksuid.New().String()
		ysuploadfile.Lid = rlid
		ysuploadfile.SrcFilePath = filepackets.SrcFilePath
		ysuploadfile.UploadFileServerID = filepackets.UploadFileServerID
		ysuploadfile.Size = filepackets.Size
		ysuploadfile.PacketNum = filepackets.PacketNum
		if err = ysuploadfile.AddYsFileLog(err); err != nil {
			return err
		}
		for _, v := range filepackets.Packets {
			packetlog := md.YsPacket{}
			packetlog.ID = ksuid.New().String()
			packetlog.Yid = ysuploadfile.ID
			packetlog.Offset = v.Offset
			packetlog.Status = 1
			packetlog.SortID = v.SortID
			packetlog.SrcPacketPath = v.Packetpath
			packetlog.UploadPacketURL = filepackets.PacketUploadURL + strconv.Itoa(v.SortID)
			if err = ysuploadfile.AddYsPacketLog(packetlog); err != nil {
				fmt.Println(err.Error())
			}
		}
		file, err := os.OpenFile(filepackets.SrcFilePath, os.O_RDONLY, os.ModePerm)
		if err != nil {
			for _, v := range filepackets.Packets {
				ysuploadfile.UpdateYspacketLog(v.SortID, err)
			}
			filepackets.UploadDone("2")
			return err
		}
		defer file.Close()
		alldoneerr := 0
		for _, v := range filepackets.Packets {
			err := ysuploadfile.UpdateYspacketLog(v.SortID, filepackets.CreatePacket(file, v))
			if err != nil {
				alldoneerr++
			}
			ysuploadfile.UpdateYspacketLog(v.SortID, err)
		}
		if alldoneerr > 0 {
			filepackets.UploadDone("2")
		} else {
			filepackets.UploadDone("0")
		}

	}
	return err
}

//Backup ...
func Backup(TaskID string, nowreturn bool) error {
	t := md.Task{}
	t.ID = TaskID
	if err := t.Select(); err != nil {
		return err
	}

	log := new(md.Log)

	if res, err := md.NewLog(TaskID); err != nil {
		return err
	} else {
		log = res
	}

	if log.IfSaveLocal != 4 {
		log.Status = 1
		if err := log.Update("status"); err != nil {
			return err
		}
	}

	if str, err := ExecBackup(&t); err != nil {
		log.Status = 2
		log.Msg = err.Error()
		log.Update("status", "msg")
		return err
	} else {
		log.Status = 0
		log.Localfilepath = str
		log.Update("status", "localfilepath")
	}

	if RS, err := md.TaskFindRemote(TaskID); err != nil {
		return err
	} else {
		for _, v := range RS {
			remoteSendLog := new(md.Rlog)
			if res, err := md.AddNewRemoteSendLog(log.ID, TaskID, v.ID); err != nil {
				return err
			} else {
				remoteSendLog = res
			}
			err = ExecRemote(v, remoteSendLog.ID, log.ID, log.Localfilepath)
			remoteSendLog.Update(err)
		}
	}

	var reerr error
	task := md.Task{}
	task.ID = TaskID
	if _, err := md.Localdb().Get(&task); err == nil {
		ExpireDelete(TaskID, task.Expire)
	}
	return reerr

}

//Jobs ...
func Jobs(i string) func() error {
	return func() error {
		err := Backup(i, false)
		if err != nil {
			m := new(mail.MailConn)
			m.SendMail("备份失败", err.Error())
		}
		return nil
	}
}

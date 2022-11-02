package jobs

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"time"

	md "youbei/models"
	zipz "youbei/utils/zip"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func SshJobs(id string) func() error {
	return func() error {
		if err := SshBackup(id, false); err != nil {
			fmt.Println(err.Error())
		}
		return nil
	}
}

func SshBackup(id string, bol bool) error {
	return execSshBackup(id)
}

func execSshBackup(id string) error {
	sshtask := md.SshTask{}
	bol, err := md.Localdb().ID(id).Get(&sshtask)
	if err != nil {
		return err
	}
	if !bol {
		return errors.New("false")
	}

	sc := new(SshConfig)
	sc.Addr = sshtask.Host
	sc.Port = sshtask.SshPort
	sc.User = sshtask.SshUser
	sc.Password = sshtask.SshPwd
	if err := sc.sshClient(); err != nil {
		return err
	}
	defer sc.Session.Close()
	nowtime := time.Now().Format("2006-01-02_15-04-05")
	dist := sshtask.DbHost + "_mysql_" + sshtask.DbName + "_" + sshtask.Char + "_" + nowtime + ".sql"
	distzip := sshtask.DbHost + "_mysql_" + sshtask.DbName + "_" + sshtask.Char + "_" + nowtime + ".zip"
	srcFile := "/tmp/" + dist
	ditFile := sshtask.SavePath + "/" + dist
	ditzipFile := sshtask.SavePath + "/" + distzip
	cmd := "mysqldump -h" + sshtask.DbHost + " -P" + sshtask.DbPort + " -u" + sshtask.DbUser + " -p" + sshtask.DbPwd + " " + sshtask.DbName + " > " + srcFile
	if err := sc.Run(cmd); err != nil {
		return err
	}
	if sshtask.SavePath != "" {
		if sftpClient, err := sftp.NewClient(sc.Client); err != nil {
			return err
		} else {
			sc.SftpClient = sftpClient
		}
		defer sc.SftpClient.Close()
		srcFile, _ := sc.SftpClient.Open(srcFile)
		dstFile, _ := os.Create(ditFile)
		defer func() {
			srcFile.Close()
			dstFile.Close()
		}()
		if _, err := srcFile.WriteTo(dstFile); err != nil {
			return errors.New("error occurred == " + err.Error())
		} else {
			srcFile.Close()
			dstFile.Close()
			sc.Session.Close()
			sc.Client.Close()
			sc.SftpClient.Close()
		}
		if err := zipz.Zip(ditFile, ditzipFile, sshtask.Zippwd); err != nil {
			return err
		}
		os.Remove(ditFile)
		rs, err := md.TaskFindRemote(sshtask.ID)
		if err != nil {
			return err
		}
		for _, v := range rs {
			err := TestRemote(ditzipFile, v)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

type SshConfig struct {
	Addr       string
	Port       string
	User       string
	Password   string
	Session    *ssh.Session
	Client     *ssh.Client
	SftpClient *sftp.Client
}

func (c *SshConfig) sshClient() error {
	var err error
	c.Client, err = ssh.Dial("tcp", c.Addr+":"+c.Port, &ssh.ClientConfig{
		User:            c.User,
		Auth:            []ssh.AuthMethod{ssh.Password(c.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		return err
	}
	session, err := c.Client.NewSession()
	if err != nil {
		return err
	}
	c.Session = session
	return nil
}

func (c *SshConfig) Run(cmd string) error {
	var b bytes.Buffer
	var d bytes.Buffer
	c.Session.Stdout = &b
	c.Session.Stderr = &d
	if err := c.Session.Run(cmd); err != nil {
		return errors.New(d.String())
	}
	return nil
}

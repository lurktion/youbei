package rs

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// ClientConfig 连接的配置
type ClientConfig struct {
	Host       string       //ip
	Port       int          // 端口
	Username   string       //用户名
	Password   string       //密码
	sshClient  *ssh.Client  //ssh client
	sftpClient *sftp.Client //sftp client
	LastResult string       //最近一次运行的结果
	SrcPath    string
	DstPath    string
}

//NewSftp ...
func (rs *RS) NewSftp() (*ClientConfig, error) {
	var (
		sshClient  *ssh.Client
		sftpClient *sftp.Client
		err        error
	)
	cliConf := new(ClientConfig)
	cliConf.Host = rs.Host
	cliConf.Port = rs.Port
	cliConf.Username = rs.UploadUser
	cliConf.Password = rs.UploadPassword
	cliConf.SrcPath = rs.SrcFilePath
	cliConf.DstPath = rs.DstFilePath
	config := ssh.ClientConfig{
		User: cliConf.Username,
		Auth: []ssh.AuthMethod{ssh.Password(cliConf.Password)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 10 * time.Second,
	}
	addr := fmt.Sprintf("%s:%d", cliConf.Host, cliConf.Port)

	if sshClient, err = ssh.Dial("tcp", addr, &config); err != nil {
		return nil, err
	}
	cliConf.sshClient = sshClient

	//此时获取了sshClient，下面使用sshClient构建sftpClient
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		log.Fatalln("error occurred:", err)
	}
	cliConf.sftpClient = sftpClient
	return cliConf, nil
}

//RunShell ...
func (cliConf *ClientConfig) RunShell(shell string) error {
	var (
		session *ssh.Session
		err     error
	)

	//获取session，这个session是用来远程执行操作的
	if session, err = cliConf.sshClient.NewSession(); err != nil {
		return err
	}
	//执行shell
	if _, err := session.CombinedOutput(shell); err != nil {
		return err
	}
	return nil
}

//Upload ...
func (cliConf *ClientConfig) Upload() error {
	srcFile, _ := os.Open(cliConf.SrcPath)                   //本地
	dstFile, _ := cliConf.sftpClient.Create(cliConf.DstPath) //远程
	defer func() {
		_ = srcFile.Close()
		_ = dstFile.Close()
	}()
	buf := make([]byte, 1024)
	for {
		n, err := srcFile.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatalln("error occurred:", err)
			} else {
				break
			}
		}
		_, _ = dstFile.Write(buf[:n])
	}
	return cliConf.RunShell(fmt.Sprintf("ls %s", cliConf.DstPath))
}

//Download ...
func (cliConf *ClientConfig) Download(srcPath, dstPath string) error {
	srcFile, _ := cliConf.sftpClient.Open(srcPath) //远程
	dstFile, _ := os.Create(dstPath)               //本地
	defer func() {
		_ = srcFile.Close()
		_ = dstFile.Close()
	}()

	if _, err := srcFile.WriteTo(dstFile); err != nil {
		return errors.New("error occurred == " + err.Error())
	}
	return nil
}

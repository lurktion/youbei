package rs

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/dutchcoders/goftp"
)

//FtpUpload ...
func (f *RS) FtpUpload() error {
	var err error
	var ftp *goftp.FTP

	if f.Host == "" {
		return errors.New("server not found")
	}

	if f.Port == 0 {
		f.Port = 21
	}

	if ftp, err = goftp.Connect(f.Host + ":" + strconv.Itoa(f.Port)); err != nil {
		return err
	}

	defer ftp.Close()
	fmt.Println("Successfully connected !!")

	if f.UploadUser != "" && f.UploadPassword != "" {
		if err = ftp.Login(f.UploadUser, f.UploadPassword); err != nil {
			return err
		}
	}

	if f.FileCnki {
		var files []string
		if files, err = ftp.List(f.DstFilePath); err != nil {
			return err
		}
		if len(files) > 0 {
			return errors.New("file is exist")
		}
	}

	var file *os.File
	if file, err = os.Open(f.SrcFilePath); err != nil {
		return err
	}
	if err := ftp.Stor(f.DstFilePath, file); err != nil {
		return err
	}
	defer ftp.Close()
	return nil
}

package jobs

import (
	"os"

	md "youbei/models"
	rs "youbei/utils/rs"
)

//TestRemote ...
func TestRemote(zipfile string, v md.RemoteStorage) error {
	yuancheng := rs.NewRS(v.Host, v.Port, v.Username, v.Password, zipfile, v.Path)
	var err error
	if v.Types == "ftp" {
		err = yuancheng.FtpUpload()
	} else if v.Types == "sftp" {
		f, err := yuancheng.NewSftp()
		if err == nil {
			err = f.Upload()
		}
	} else if v.Types == "Yserver" {
		fp, err := yuancheng.ReadBigFile()
		if err != nil {
			return err
		}
		file, err := os.OpenFile(yuancheng.SrcFilePath, os.O_RDONLY, os.ModePerm)
		defer file.Close()
		if err != nil {
			return err
		}

		for _, v := range fp.Packets {
			err = fp.CreatePacket(file, v)
		}
	}
	return err
}

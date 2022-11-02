package Zipz

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/yeka/zip"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func Zip(srcFile string, destZip string, mima string) error {
	sysType := runtime.GOOS

	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()
	archive := zip.NewWriter(zipfile)
	defer archive.Close()
	_, err = os.Stat(srcFile)
	if err != nil {
		return err
	}

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		pathstr := strings.Replace(path, string(os.PathSeparator), "/", -1)
		fdstr := strings.Replace(filepath.Dir(srcFile), string(os.PathSeparator), "/", -1)
		header.Name = strings.TrimLeft(strings.TrimPrefix(pathstr, fdstr), "/")
		if sysType == "windows" {
			header.Name, _ = utf8ToGBK(header.Name)
		}
		if info.IsDir() {
			header.Name = header.Name + "/"
		} else {
			header.Method = zip.Deflate
		}
		var writer io.Writer
		if mima == "" {
			writer, err = archive.CreateHeader(header)
		} else {
			writer, err = archive.Encrypt(header.Name, mima, zip.AES256Encryption)
		}

		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}

func utf8ToGBK(text string) (string, error) {
	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewEncoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}
	return string(dst[:nDst]), nil
}

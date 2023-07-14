package Zipz

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync/atomic"

	"github.com/yeka/zip"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// Unzip 解压zip文件到指定目录
// srcZip - 源zip文件路径
// mima - zip文件的密码
// writer - io.WriteCloser用于接收解压的文件内容
// 返回值:
// 解压后的文件名
// error
func Unzip(srcZip string, mima string, writer io.WriteCloser, progress *int64) error {
	// 打开zip文件
	r, err := zip.OpenReader(srcZip)
	if err != nil {
		return err
	}
	defer r.Close()

	// 确保只有一个文件在zip内
	if len(r.File) != 1 {
		return fmt.Errorf("zip file should contain exactly one file")
	}

	// 从zip中获取第一个文件
	f := r.File[0]

	// 获取文件的总大小
	totalSize := f.UncompressedSize64

	// 设置zip文件密码
	if f.IsEncrypted() {
		f.SetPassword(mima)
	}

	// 打开文件读取内容
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	// 使用scanner逐行读取文件内容
	scanner := bufio.NewScanner(rc)

	// 设置较大的MaxScanTokenSize
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 10*1024*1024) // 设置10MB的最大扫描令牌大小

	var writtenSize uint64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		n, err := writer.Write([]byte(line + "\n")) // 逐行写入数据到writer，添加换行符"\n"
		if err != nil {
			return err
		}

		// 更新已写入的大小
		writtenSize += uint64(n)

		// 计算并更新进度
		currentProgress := int64(float64(writtenSize) / float64(totalSize) * 100)
		atomic.StoreInt64(progress, currentProgress)

	}

	// 检查scanner在逐行读取过程中是否发生错误
	if err := scanner.Err(); err != nil {
		return err
	}

	atomic.StoreInt64(progress, 100)

	// 关闭写入器
	if err = writer.Close(); err != nil {
		return err
	}

	return nil
}

// Zip 创建zip压缩文件
// srcFile - 源文件路径
// destZip - 目标zip文件路径
// mima - zip文件的密码
// 返回值:
// error
func Zip(srcFile string, destZip string, mima string) error {
	// 获取当前系统类型
	sysType := runtime.GOOS

	// 创建zip文件
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	// 创建zip.Writer
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// 遍历srcFile目录下的所有文件和目录
	err = filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 获取文件头信息
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// 转换路径格式，并去掉前缀
		pathstr := strings.ReplaceAll(path, string(os.PathSeparator), "/")
		fdstr := strings.ReplaceAll(filepath.Dir(srcFile), string(os.PathSeparator), "/")
		header.Name = strings.TrimLeft(strings.TrimPrefix(pathstr, fdstr), "/")

		// 处理windows平台下的文件名编码
		if sysType == "windows" {
			header.Name, _ = utf8ToGBK(header.Name)
		}

		// 设置zip文件中的目录结构
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		// 创建zip文件中的文件
		var writer io.Writer
		if mima == "" {
			writer, err = archive.CreateHeader(header)
		} else {
			writer, err = archive.Encrypt(header.Name, mima, zip.AES256Encryption)
		}

		if err != nil {
			return err
		}

		// 将源文件内容写入zip文件中的文件
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

// utf8ToGBK 将UTF-8编码的字符串转换为GBK编码
// text - 需要转换的字符串
// 返回值:
// 转换后的字符串
// error
func utf8ToGBK(text string) (string, error) {
	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewEncoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}
	return string(dst[:nDst]), nil
}

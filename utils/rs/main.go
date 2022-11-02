package rs

import (
	"strings"
)

// RS ...
type RS struct {
	FileName       string
	SrcDir         string
	SrcFilePath    string
	DstDir         string
	DstFilePath    string
	Host           string
	Port           int
	UploadUser     string
	UploadPassword string
	FileCnki       bool
}

// getFileName ...
func getFileName(str string) string {
	strs := strings.Split(strings.TrimLeft(getSaveDirstring(str), "/"), "/")
	return strs[len(strs)-1]
}

// getFileDir ...
func getFileDir(str string) string {
	strs := strings.Split(strings.TrimLeft(getSaveDirstring(str), "/"), "/")
	return strings.Join(strs[0:len(strs)-2], "/")
}

// getSaveDirstring ...
func getSaveDirstring(str string) string {
	return strings.TrimRight(strings.Replace(str, "\\", "/", -1), "/")
}

// NewRS ...
func NewRS(host string, port int, user, password, sourcefile, distpath string) *RS {
	rs := new(RS)
	rs.FileName = getFileName(sourcefile)
	rs.SrcDir = getFileDir(sourcefile)
	rs.SrcFilePath = getSaveDirstring(sourcefile)
	rs.DstDir = distpath
	rs.DstFilePath = getSaveDirstring(distpath) + "/" + getFileName(sourcefile)
	rs.Host = host
	rs.Port = port
	rs.UploadUser = user
	rs.UploadPassword = password
	return rs
}

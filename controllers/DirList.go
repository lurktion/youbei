package controllers

import (
	"io/ioutil"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"

	"strings"
)

// DirList ...
func DirList(c *gin.Context) {
	dir := c.Query("dir")
	var bol bool
	bol = c.GetBool("isdir")

	dir = strings.Replace(dir, "\\", "/", -1)
	dir = strings.TrimRight(dir, "/") + "/"
	if dir == "/" {
		sysType := runtime.GOOS
		if sysType == "windows" {
			APIReturn(c, 200, "获取盘符成功", GetLogicalDrives())
			return
		}
	}
	APIReturn(c, 200, "获取目录成功", ListDir(dir, bol))
}

// GetLogicalDrives ...
func GetLogicalDrives() []Dir {
	var drivesAll = []string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P：", "Q：", "R：", "S：", "T：", "U：", "V：", "W：", "X：", "Y：", "Z："}

	var drives []string
	for _, v := range drivesAll {
		fi, err := os.Stat(v)
		if err == nil {
			if fi.IsDir() {
				drives = append(drives, v)
			}
		}
	}

	list := []Dir{}
	for _, k := range drives {
		f := Dir{}
		f.Path = k
		f.Label = k
		f.IsLeaf = false
		list = append(list, f)
	}
	return list
}

// Dir ...
type Dir struct {
	Label    string `json:"label"`
	IsLeaf   bool   `json:"isLeaf"`
	Children string `json:"children"`
	Path     string `json:"path"`
}

// ListDir ...
func ListDir(dirPth string, bol bool) []Dir {
	list := []Dir{}
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		panic(err)
	}
	for _, fi := range dir {
		f := Dir{}
		if bol {
			f.Label = fi.Name()
			if fi.IsDir() {
				f.IsLeaf = false
			} else {
				f.IsLeaf = true
			}
			f.Path = dirPth + fi.Name()
			list = append(list, f)
		} else {
			if fi.IsDir() {
				f.Label = fi.Name()
				f.IsLeaf = false
				f.Path = dirPth + fi.Name()
				list = append(list, f)
			}
		}

	}
	return list
}

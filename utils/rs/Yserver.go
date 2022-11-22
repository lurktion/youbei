package rs

import (
	"errors"
	"os"
	"strconv"

	"github.com/segmentio/ksuid"

	"github.com/beego/beego/httplib"
)

//resAPIPacketup ...
type resAPIPacketup struct {
	Success int         `json:"success"`
	Msg     string      `json:"msg"`
	Result  interface{} `json:"result"`
}

//UploadByte ...
func UploadByte(url string, b []byte) error {
	res := resAPIPacketup{}
	err := httplib.Post(url).Body(b).ToJSON(&res)
	return err

}

//uploadfile ...
func (c *Filepackets) uploadfile() error {
	res := resAPIPacketup{}
	req := httplib.Post(c.FileUploadURL)
	req.Param("username", c.UploadUser)
	req.Param("password", c.UploadPassword)
	req.Param("filename", c.FileName)
	req.Param("savedir", c.DstDir)
	req.Param("size", strconv.FormatInt(c.Size, 10))
	req.Param("packetnum", strconv.FormatInt(c.PacketNum, 10))
	err := req.ToJSON(&res)
	if err != nil {
		return err
	}
	if res.Success != 200 {
		return errors.New(res.Msg)
	}
	return nil
}

//ReadBigFile ...
func (c *RS) ReadBigFile() (*Filepackets, error) {
	fp, err := c.readMyFile()
	if err != nil {
		return nil, err
	}
	err = fp.uploadfile()
	if err != nil {
		return nil, err
	}
	return fp, nil
}

//Filepacket ...
type Filepackets struct {
	UploadFileServerID string

	FileUploadURL string
	FileName      string
	SrcDir        string
	SrcFilePath   string
	Size          int64

	DstDir      string
	DstFilePath string

	UploadUser     string
	UploadPassword string

	PacketUploadURL string
	Packets         map[int]Packet
	PacketSize      int64
	PacketNum       int64

	UploadDoneURL string
}

//Packet ...
type Packet struct {
	SortID     int
	Packetpath string
	Offset     int64
	Error      error
}

//readMyFile ...
func (c *RS) readMyFile() (*Filepackets, error) {
	filepackets := new(Filepackets)
	filepackets.FileName = c.FileName
	filepackets.SrcFilePath = c.SrcFilePath
	filepackets.SrcDir = c.SrcDir

	fi, err := os.Stat(filepackets.SrcFilePath)
	if err != nil {
		return nil, err
	}
	filepackets.Size = fi.Size()
	filepackets.PacketSize = int64(20 * 1024 * 1024)
	filepackets.PacketNum = filepackets.Size / filepackets.PacketSize
	yu := filepackets.Size % filepackets.PacketSize
	if yu > 0 {
		filepackets.PacketNum++
	}
	fps := map[int]Packet{}
	for i := int64(0); i < filepackets.PacketNum; i++ {
		packet := Packet{}
		packet.SortID = int(i)
		packet.Offset = i * filepackets.PacketSize
		fps[packet.SortID] = packet
	}
	filepackets.Packets = fps
	filepackets.UploadFileServerID = ksuid.New().String()
	filepackets.FileUploadURL = "http://" + c.Host + ":" + strconv.Itoa(c.Port) + "/upload/file/" + filepackets.UploadFileServerID
	filepackets.PacketUploadURL = "http://" + c.Host + ":" + strconv.Itoa(c.Port) + "/upload/packet/" + filepackets.UploadFileServerID + "/"
	filepackets.UploadDoneURL = "http://" + c.Host + ":" + strconv.Itoa(c.Port) + "/upload/done/"
	filepackets.DstDir = c.DstDir
	filepackets.DstFilePath = c.DstFilePath
	filepackets.UploadUser = c.UploadUser
	filepackets.UploadPassword = c.UploadPassword
	return filepackets, nil
}

//CreatePacket ...
func (c *Filepackets) CreatePacket(file *os.File, v Packet) error {
	buf := make([]byte, c.PacketSize)
	nr, _ := file.ReadAt(buf[:], v.Offset)
	var err error
	if nr == 0 {
		err = nil
	} else if nr > 0 {
		err = UploadByte(c.PacketUploadURL+strconv.FormatInt(v.Offset, 10), buf[0:nr])
	} else {
		err = errors.New("nr < 0")
	}
	return err
}

// UploadDone ...
func (c *Filepackets) UploadDone(status string) error {
	res := resAPIPacketup{}
	err := httplib.Post(c.UploadDoneURL+c.UploadFileServerID).Param("status", status).ToJSON(&res)
	return err
}

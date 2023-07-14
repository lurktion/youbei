package md

import (
	"time"

	"github.com/segmentio/ksuid"
)

// RemoteStorage ...
type RemoteStorage struct {
	ID       string `json:"id" xorm:"pk notnull unique 'id'"`
	Name     string `json:"name" xorm:"'name'"`
	Types    string `json:"types" xorm:"'types'"`
	Username string `json:"username" xorm:" 'username' comment('用户名')"`
	Password string `json:"password" xorm:" 'password' comment('密码')"`
	Host     string `json:"host" xorm:" 'host' comment('ip地址')"`
	Path     string `json:"path" xorm:"'path' comment('路径')"`
	Port     int    `json:"port" xorm:"'port' comment('端口')"`
	Created  int64  `json:"created" xorm:"notnull 'created'"`
	ReLink   int    `json:"relink" xorm:"'relink'"`
}

// RemoteStorageToTask ...
type RemoteStorageToTask struct {
	ID      string `json:"id" xorm:"pk notnull unique 'id'"`
	Rid     string `json:"rid" xorm:"'rid' comment('远程存储ID')"`
	Tid     string `json:"tid" xorm:"'tid' comment('任务ID')"`
	Created int64  `json:"created" xorm:"notnull 'created'"`
}

// Add 添加远程存储 单例
func (f *RemoteStorage) Add() error {
	f.Created = time.Now().Unix()
	_, err := localdb.Insert(f)
	if err != nil {
		return err
	}
	return nil
}

// Update 查询远程存储 单例
func (f *RemoteStorage) Update() error {
	_, err := localdb.ID(f.ID).Update(f)
	if err != nil {
		return err
	}
	return nil
}

// TaskFindRemote 查询远程存储 单例
func TaskFindRemote(tid string) ([]RemoteStorage, error) {
	f := []RemoteStorageToTask{}
	rs := []RemoteStorage{}
	err := localdb.Where("tid=?", tid).Find(&f)
	if err != nil {
		return rs, err
	}
	if len(f) > 0 {
		fs := []string{}
		for _, v := range f {
			fs = append(fs, v.Rid)
		}
		err = localdb.In("id", fs).Find(&rs)
		if err != nil {
			return rs, err
		}
	}

	return rs, nil
}

// Delete 删除远程存储 单例
func (f *RemoteStorage) Delete() error {
	if _, err := localdb.ID(f.ID).Delete(f); err != nil {
		return err
	}
	if err := RemoteStorageDeleteALL(f.ID); err != nil {
		return err
	}
	return nil
}

// RemoteStorageDeleteALL 删除远程存储 单例
func RemoteStorageDeleteALL(rid string) error {
	if _, err := localdb.Exec("delete from rlog where rid = ?", rid); err != nil {
		return err
	}
	if _, err := localdb.Exec("delete from remote_storage_to_task where rid = ?", rid); err != nil {
		return err
	}
	return nil
}

// RemoteStorageToTaskFunc 添加远程存储和任务关联
func RemoteStorageToTaskFunc(tid string, rids []string) error {
	if _, err := localdb.Where("tid=?", tid).Delete(new(RemoteStorageToTask)); err != nil {
		return err
	}
	for _, rid := range rids {
		bol, err := RemoteStorageToTaskCheck(tid, rid)
		if err != nil {
			return err
		}
		if bol {
			return nil
		}
		f := RemoteStorageToTask{}
		f.ID = ksuid.New().String()
		f.Rid = rid
		f.Tid = tid
		f.Created = time.Now().Unix()
		_, err = localdb.Insert(&f)
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoteStorageToTaskCheck 远程存储和任务关联 查重
func RemoteStorageToTaskCheck(tid, rid string) (bool, error) {
	f := RemoteStorageToTask{}
	f.Rid = rid
	f.Tid = tid
	return localdb.Exist(&f)

}

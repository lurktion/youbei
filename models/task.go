package md

import (
	"errors"
	"strings"
	"time"

	// gosqlite3
	_ "github.com/mattn/go-sqlite3"
)

// Task ...
type Task struct {
	ID           string   `json:"id" xorm:"pk notnull unique 'id'"`
	Name         string   `json:"name" xorm:"'name'"`
	Types        int      `json:"types" xorm:"default(0) 'types'"` //0 备份本机 1 备份其他机器
	DBType       string   `json:"dbtype" xorm:" 'dbtype'"`
	Host         string   `json:"host" xorm:" 'host'"`
	Port         string   `json:"port" xorm:" 'port'"`
	User         string   `json:"user" xorm:" 'user'"`
	Password     string   `json:"password" xorm:" 'password'"`
	DBname       string   `json:"dbname" xorm:" 'dbname'"`
	DBnames      []string `json:"dbnames" xorm:"-"`
	Char         string   `json:"char" xorm:" 'char'"`
	DBpath       string   `json:"dbpath" xorm:" 'dbpath'"`
	DefDbname    string   `json:"defdbname" xorm:"'defdbname'"`
	Created      int64    `json:"created" xorm:"notnull 'created'"`
	Status       string   `json:"status" xorm:" 'status'"`
	Pause        string   `json:"pause" xorm:"'pause'"`
	PauseNum     int      `json:"pausenum" xorm:"default(0) pausenum"`
	Crontab      string   `json:"crontab" xorm:"notnull 'crontab'"`
	SavePath     string   `json:"savepath" xorm:"notnull 'savepath'"`
	RS           []string `json:"rs" xorm:"-"`
	Enablezippwd int      `json:"enablezippwd" xorm:"default(0) 'enablezippwd'"` //0否 1是
	Zippwd       string   `json:"zippwd" xorm:"'zippwd'"`
	Expire       int      `json:"expire" xorm:"'expire'"`
	Cmds         string   `json:"cmds" xorm:"'cmds'"`
}

// Add ...
func (t *Task) Add() error {
	if name, err := CreateId("TASK", strings.ToUpper(t.DBType)); err != nil {
		return err
	} else {
		t.Name = name
	}
	t.Created = time.Now().Unix()
	_, err := localdb.Insert(t)
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (t *Task) Update() error {
	_, err := localdb.ID(t.ID).AllCols().Update(t)
	if err != nil {
		return err
	}
	return nil
}

// Select ...
func (t *Task) Select() error {
	bol, err := localdb.ID(t.ID).Get(t)
	if err != nil {
		return err
	}
	if !bol {
		return errors.New("not found")
	}
	rs := []RemoteStorageToTask{}
	if err := localdb.Where("tid=?", t.ID).Find(&rs); err != nil {
		return err
	}
	if len(rs) > 0 {
		for _, s := range rs {
			t.RS = append(t.RS, s.Rid)
		}
	}
	return nil
}

// SelectAll ...
func SelectAll(typestr string, page, count int) ([]Task, error) {
	var err error
	tasks := []Task{}
	var types []string
	if typestr == "file" {
		types = []string{`file`}
	} else {
		types = []string{`sqlite3`, `mysql`, `mssql`, `postgres`, `mongodb`}
	}
	if err := localdb.In("dbtype", types).Desc("created").Limit(count, count*(page-1)).Find(&tasks); err != nil {

		return []Task{}, err
	}
	for k, v := range tasks {
		rs := []RemoteStorageToTask{}
		rss := []string{}
		if err := localdb.Where("tid=?", v.ID).Find(&rs); err != nil {

			return []Task{}, err
		}
		if len(rs) > 0 {
			for _, s := range rs {
				rss = append(rss, s.Rid)
			}
		}
		tasks[k].RS = rss
	}
	return tasks, err
}

// TaskCount ...
func TaskCount(typestr string) (int64, error) {
	var err error
	tasks := Task{}
	var types []string
	if typestr == "file" {
		types = []string{`file`}
	} else {
		types = []string{`sqlite3`, `mysql`, `mssql`}
	}
	title, err := localdb.In("dbtype", types).Count(&tasks)
	return title, err
}

// All ...
func All() ([]Task, error) {
	t := []Task{}
	//err := localdb.Where("pausenum=?", 0).Find(&t)
	err := localdb.Find(&t)
	if err != nil {
		return t, err
	}
	return t, nil
}

// Delete ...
func (t *Task) Delete() error {
	_, err := localdb.ID(t.ID).Delete(t)
	if err != nil {
		return err
	}
	err = TaskDeleteALL(t.ID)
	return err
}

// TaskDeleteALL ...
func TaskDeleteALL(tid string) error {
	if _, err := localdb.Exec("delete from log where tid = ?", tid); err != nil {
		return err
	}
	if _, err := localdb.Exec("delete from rlog where tid = ?", tid); err != nil {
		return err
	}
	if _, err := localdb.Exec("delete from remote_storage_to_task where tid = ?", tid); err != nil {
		return err
	}
	return nil
}

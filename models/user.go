package md

import (
	"errors"
	"time"

	"youbei/utils"

	"github.com/segmentio/ksuid"
)

// User ...
type User struct {
	ID       string   `json:"id" xorm:"pk notnull unique 'id'"`
	Name     string   `json:"name" xorm:"unique 'name'"`
	Passwrod string   `json:"password" xorm:"'password'"`
	IPlist   []string `json:"iplist" xorm:"json 'iplist'"`
	Created  int64    `json:"created" xorm:"'created'"`
}

// UserLogin ...
func UserLogin(name, password string) (*User, error) {
	user := new(User)
	user.Name = name
	user.Passwrod = password
	bol, err := localdb.Get(user)
	if err != nil {
		return user, err
	}
	if !bol {
		return user, errors.New("not found user")
	}
	return user, nil
}

// UserChangePwd ...
func UserChangePwd(id, oldpassword, password string) error {
	user := User{}
	user.Passwrod = password
	bol, err := localdb.Where("id=? and password=?", id, oldpassword).Cols("password").Update(&user)
	if err != nil {
		return err
	}
	if bol < 1 {
		return errors.New("not found user")
	}

	return nil
}

// Add ...
func (u *User) Add() error {
	u.ID = ksuid.New().String()
	u.Created = time.Now().Unix()
	u.Passwrod = utils.Md5V(u.Passwrod)
	_, err := localdb.Insert(u)
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (u *User) Update() error {
	_, err := localdb.ID(u.ID).Cols("iplist").Update(u)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser ...
func DeleteUser(id string) error {
	user := new(User)
	_, err := localdb.ID(id).Delete(user)
	return err
}

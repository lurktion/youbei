package db

import (
	"context"
	"errors"
	"fmt"

	md "youbei/models"

	"gitee.com/countpoison_admin/xorm"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PostgresConnectTest ...
func ConnectTest(info *md.Task) error {
	if info.DBType == "postgres" {
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", info.Host, info.Port, info.User, info.Password, info.DefDbname)
		//格式
		if db, err := xorm.NewEngine("postgres", psqlInfo); err != nil {
			return err
		} else {
			err = db.Ping()
			defer db.Close()
			return err
		}
	} else if info.DBType == "mssql" {
		if db, err := xorm.NewEngine("mssql", "server="+info.Host+";user id="+info.User+";password="+info.Password+";port="+info.Port+";database="+info.DBname+";encrypt=disable"); err != nil {
			return err
		} else {
			if err = db.Ping(); err != nil {
				fmt.Println(err.Error())
			}
			defer db.Close()
			return err
		}
	} else if info.DBType == "mysql" {
		if db, err := xorm.NewEngine("mysql", info.User+":"+info.Password+"@("+info.Host+":"+info.Port+")/"+info.DBname+"?charset="+info.Char); err != nil {
			return err
		} else {
			err = db.Ping()
			defer db.Close()
			return err
		}
	} else if info.DBType == "sqlite" {
		if db, err := xorm.NewEngine("sqlite3", info.DBpath); err != nil {
			return err
		} else {
			err = db.Ping()
			defer db.Close()
			return err
		}
	} else if info.DBType == "mongodb" {
		credential := options.Credential{
			Username: info.User,
			Password: info.Password,
		}
		if info.DefDbname != "" {
			credential.AuthSource = info.DefDbname
		}
		clientOptions := options.Client().ApplyURI("mongodb://" + info.Host + ":" + info.Port).SetAuth(credential)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			return err
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			return err
		}

	} else {
		return errors.New("dbtype not found4")
	}
	return nil
}

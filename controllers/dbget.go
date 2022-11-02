package controllers

import (
	"context"
	"errors"
	"fmt"

	"gitee.com/countpoison_admin/xorm"
	"github.com/gin-gonic/gin"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbinfo struct {
	User     string
	Password string
	Char     string
}

func DBGet(c *gin.Context) {
	User := c.Query("User")
	Password := c.Query("Password")
	Host := c.Query("Host")
	Port := c.Query("Port")
	Char := c.Query("Char")
	DBtype := c.Query("Dbtype")
	Defdbname := c.Query("Defdbname")
	dbs := []string{}

	if DBtype == "mysql" {
		if db, err := xorm.NewEngine("mysql", User+":"+Password+"@("+Host+":"+Port+")/?charset="+Char); err != nil {
			APIReturn(c, 500, "连接数据库失败1", err)
			return
		} else {
			if err := db.Ping(); err != nil {
				APIReturn(c, 500, "测试连接失败", err)
				return
			}
			sql := "show databases;"
			if err := db.SQL(sql).Find(&dbs); err != nil {
				APIReturn(c, 500, "查询数据库失败", err)
				return
			}
			db.Close()
		}

	} else if DBtype == "mssql" {
		if db, err := xorm.NewEngine("mssql", "server="+Host+";user id="+User+";password="+Password+";port="+Port+";encrypt=disable;database=master"); err != nil {
			APIReturn(c, 500, "连接数据库失败2", err)
			return
		} else {
			if err := db.Ping(); err != nil {
				APIReturn(c, 500, "测试连接失败", err)
				return
			}
			sql := "SELECT Name FROM Master..SySdatabases ORDER BY Name;"
			if err := db.SQL(sql).Find(&dbs); err != nil {
				APIReturn(c, 500, "查询数据库失败2", err)
				return
			}
			db.Close()
		}
	} else if DBtype == "postgres" {
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Defdbname)
		if db, err := xorm.NewEngine("postgres", psqlInfo); err != nil {
			APIReturn(c, 500, "连接数据库失败3", err)
			return
		} else {
			if err := db.Ping(); err != nil {
				APIReturn(c, 500, "测试连接失败", err)
				return
			}
			sql := "select datname from pg_database;"
			if err := db.SQL(sql).Find(&dbs); err != nil {
				APIReturn(c, 500, "查询数据库失败", err)
				return
			}
			db.Close()
		}
	} else if DBtype == "mongodb" {
		clientOptions := options.Client().ApplyURI("mongodb://" + User + ":" + Password + "@" + Host + ":" + Port)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			APIReturn(c, 500, "数据库失败4", err)
			return
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			APIReturn(c, 500, "测试连接失败", err)
			return
		}
		dbs, err = client.ListDatabaseNames(context.TODO(), bson.M{})
		if err != nil {
			APIReturn(c, 500, "查询数据库失败", err)
			return
		}

	} else {
		APIReturn(c, 500, "数据库类型不存在", errors.New("数据库类型不存在"))
		return
	}

	APIReturn(c, 200, "成功", dbs)
}

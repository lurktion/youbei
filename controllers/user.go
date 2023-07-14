package controllers

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	md "youbei/models"

	"youbei/utils"

	"github.com/beego/beego/v2/adapter/httplib"
	jwt "github.com/dgrijalva/jwt-go"
)

//Userlogin ...
func Userlogin(c *gin.Context) {
	name := c.Param("name")
	password := c.Param("password")
	userinfo, err := md.UserLogin(name, password)
	if err != nil {
		APIReturn(c, 500, "登陆失败", err.Error())
		return
	}

	if c.ClientIP() == "::1" {
		goto Access
	}
	if len(userinfo.IPlist) > 0 {
		for _, v := range userinfo.IPlist {
			if v == c.ClientIP() {
				goto Access
			}
		}
		APIReturn(c, 401, "不允许从当前ip登录", errors.New("不允许从当前ip登录"))
		return
	}
Access:
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["username"] = name
	claims["userid"] = userinfo.ID
	claims["User"] = "true"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("youbei"))
	if err != nil {
		APIReturn(c, 500, "token生成失败", err.Error())
		return
	}

	APIReturn(c, 200, "登陆成功", tokenString)
}

type Pwd struct {
	Opwd string `json:"opwd"`
	Npwd string `json:"npwd"`
}

//Userchangepwd ...
func Userchangepwd(c *gin.Context) {
	pwds := new(Pwd)
	if err := c.Bind(pwds); err != nil {
		APIReturn(c, 500, "解析失败", err.Error())
		return
	}
	userid := c.Param("id")
	if err := md.UserChangePwd(userid, pwds.Opwd, pwds.Npwd); err != nil {
		APIReturn(c, 500, "修改失败", err.Error())
		return
	}

	APIReturn(c, 200, "修改成功", nil)
}

//Prepare ...
func Prepare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("token")
		if tokenHeader == "" {
			APIReturn(c, 401, "token不能为空", errors.New("token不能为空"))
			return
		}
		claims := make(jwt.MapClaims)
		token, err := jwt.ParseWithClaims(tokenHeader, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("youbei"), nil
		})

		if err != nil {
			APIReturn(c, 401, "验证失败1", err.Error())
			return
		}

		if !token.Valid {
			APIReturn(c, 401, "验证失败2", errors.New("验证失败"))
			return
		}
	}
}

//UserList ...
func UserList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("count"))

	users := make([]md.User, 0)
	var err error
	if limit <= 0 {
		err = md.Localdb().Desc("created").Limit(limit, limit*(page-1)).Find(&users)
	} else {
		err = md.Localdb().Desc("created").Find(&users)
	}
	if err != nil {
		APIReturn(c, 500, "获取列表失败", err.Error())
		return
	}

	count, err := md.Localdb().Count(new(md.User))
	if err != nil {
		APIReturn(c, 500, "获取用户总数失败", err.Error())
		return
	}

	rep := map[string]interface{}{"count": count, "data": users}
	APIReturn(c, 200, "获取用户列表成功", &rep)
}

//GetUser ...
func GetUser(c *gin.Context) {
	id := c.Param("id")
	user := new(md.User)
	user.ID = id
	if bol, err := md.Localdb().Get(user); err != nil {
		APIReturn(c, 500, "获取用户失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "获取用户失败", errors.New("user not found"))
			return
		}
	}

	APIReturn(c, 200, "获取用户成功", user)
}

//AddUser ...
func AddUser(c *gin.Context) {
	ob := new(md.User)
	if err := c.Bind(ob); err != nil {
		APIReturn(c, 500, "解析数据失败", err.Error())
		return
	}

	if err := ob.Add(); err != nil {
		APIReturn(c, 500, "添加用户失败", err.Error())
		return
	}

	APIReturn(c, 200, "成功", nil)
}

//EditUser ...
func EditUser(c *gin.Context) {
	ob := new(md.User)
	if err := c.Bind(ob); err != nil {
		APIReturn(c, 500, "解析数据失败", err.Error())
		return
	}

	if err := ob.Update(); err != nil {
		APIReturn(c, 500, "修改用户失败", err.Error())
		return
	}

	APIReturn(c, 200, "成功", nil)
}

//DeleteUser ...
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := md.DeleteUser(id); err != nil {
		APIReturn(c, 500, "删除用户失败", err.Error())
		return
	}

	APIReturn(c, 200, "删除用户成功", nil)
}

func ConnectHost(c *gin.Context) {
	id := c.Param("id")
	host := new(md.Host)
	if bol, err := md.Localdb().ID(id).Get(host); err != nil {
		APIReturn(c, 500, "获取信息失败", err.Error())
		return
	} else {
		if !bol {
			APIReturn(c, 500, "信息不存在", errors.New("信息不存在"))
			return
		}
	}

	url := host.Protocol + "://" + host.HostAddr + ":" + host.Port
	req := httplib.Get(url + "/login/" + host.Username + "/" + utils.Md5V(host.Password))
	resapi := ResAPI{}
	if err := req.ToJSON(&resapi); err != nil {
		fmt.Println(err.Error())
		APIReturn(c, 500, "登录失败21", err.Error())
		return
	}

	if resapi.Success != 200 {
		APIReturn(c, 500, "登录失败2", errors.New(resapi.Msg))
		return
	}
	APIReturn(c, 200, "登录成功", resapi.Result)
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	md "youbei/models"
)

// ResAPI ...
type ResAPI struct {
	Success int         `json:"success"`
	Msg     string      `json:"msg"`
	Result  interface{} `json:"result"`
}

func APIReturn(c *gin.Context, Success int, Msg string, data interface{}) {
	res := gin.H{
		"success": Success,
		"msg":     Msg,
		"result":  data,
	}
	c.JSON(Success, res)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func GetQueryParams(c *gin.Context) map[string]any {
	query := c.Request.URL.Query()
	var queryMap = make(map[string]any, len(query))
	for k := range query {
		queryMap[k] = c.Query(k)
	}
	return queryMap
}

func GetRestul(c *gin.Context, sqlresult string, v interface{}) (int64, error) {
	if err := GetRestulList(c, sqlresult, v); err != nil {
		return 0, err
	}

	total, err := GetRestulTotal(c, sqlresult)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func GetRestulTotal(c *gin.Context, sqlresult string) (int64, error) {
	maps := GetQueryParams(c)
	delete(maps, "pageSize")
	count, err := md.Localdb().SqlTemplateClient(sqlresult, &maps).Query().Count()
	if err != nil {
		return int64(0), err
	}
	return int64(count), nil
}

func GetRestulList(c *gin.Context, sqlresult string, v interface{}) error {
	maps := GetQueryParams(c)
	if err := md.Localdb().SqlTemplateClient(sqlresult, &maps).Find(v); err != nil {
		return err
	}
	return nil
}

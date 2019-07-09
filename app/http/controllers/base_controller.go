package controllers

import (
	"github.com/qit-team/snow/app/constants/errorcode"
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

/**
 * 成功时返回
 */
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":        errorcode.Success,
		"message":     "ok",
		"request_uri": c.Request.URL.Path,
		"data":        data,
	})
	c.Abort()
}

/**
 * 失败时返回
 */
func Error(c *gin.Context, code int, msg ...string) {
	message := ""
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = errorcode.GetMsg(code)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":        code,
		"message":     message,
		"request_uri": c.Request.URL.Path,
		"data":        make(map[string]string),
	})
	c.Abort()
}

func Error404(c *gin.Context) {
	Error(c, errorcode.NotFound, "路由不存在")
}

func Error500(c *gin.Context) {
	Error(c, errorcode.SystemError)
}

/**
 * 将请求的body转换为request数据结构
 * @param c
 * @param request  传入request数据结构的指针 如 new(TestRequest)
 */
func GenRequest(c *gin.Context, request interface{}) (err error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	return json.Unmarshal(body, request)
}

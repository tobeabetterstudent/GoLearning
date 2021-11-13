// @program: gin_demo2的通用工具
// @author: aslanwang
// @create: 2021-11-13

package	common

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
	"fmt"
	"GoLearning/Gin/gin_demo2/config"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"time"
)

// 获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}

// 制作返回的json数据
func MakeJSON(code, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg"  : msg,
		"data" : data,
	})
	c.Abort()
}

// MD5加密方法
func MD5(str string) string{
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// CreateSign 生成签名
// url.Values是 map[string][]string
// params就是输入的参数
func CreateSign(params url.Values) string{
	var key []string
	// 只遍历键
	for k := range params {
		if k != "sn" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	var str string
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v",key[i],params.Get(key[i]))
		} else {
			str += fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}
	return MD5(MD5(str)) + MD5(config.APP_NAME + config.APP_SECRET)
}

func VerifySign(c *gin.Context) {
	var method = c.Request.Method
	var timeStamp int64
	var sign string
	var req url.Values

	if method == "GET" {
		req = c.Request.URL.Query()
		sign = c.Query("sn")
		timeStamp, _ = strconv.ParseInt(c.Query("ts"), 10, 64)
	} else if method == "POST" {
		c.Request.ParseForm()
		req = c.Request.PostForm
		sign = c.PostForm("sn")
		timeStamp, _ = strconv.ParseInt(c.PostForm("ts"), 10, 64)
	} else {
		MakeJSON("500", "Illegal requests", "", c)
		return
	}

	exp, _ := strconv.ParseInt(config.API_EXPIRY, 10, 64)
	now := GetTimeUnix()
	// 验证过期时间
	if timeStamp > now || now - timeStamp > exp {
		MakeJSON("500", "Ts Error", "",c)
		return
	}
	// 验证签名
	if sign == "" || sign != CreateSign(req) {
		MakeJSON("500", "Sn Error", "", c)
		return
	}
}
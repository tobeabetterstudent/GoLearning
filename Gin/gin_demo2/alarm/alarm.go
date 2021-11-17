// @program: 封装完成的错误信息 并试图进行推送到微信 通过alarm.WeChat()接口报错到微信 后续还可以有alarm.Email() 这里错误类型依然保持为error
// @author: aslanwang
// @create: 2021-11-17

package alarm

import (
	"errors"
	"GoLearning/Gin/gin_demo2/common"
	"path/filepath"
	"runtime"
	"strings"
	"strconv"
	"encoding/json"
)

type errorInfo struct{
	Time     string `json:"time"`
	Alarm    string `json:"alarm"`
	Message  string `json:"message"`
	Filename string `json:"filename"`
	Line     int    `json:"line"`
	Funcname string `json:"funcname"`
}
// WeChat 报警到
func WeChat(str string) error {
	err := alarm("WX", str)
	return errors.New(err)
}

// 告警方法
func alarm(method, str string) (err string){
	// 当前时间
	nowTime := strconv.FormatInt(common.GetTimeUnix(),10) 

	// 定义 文件名、行号、方法名
	fileName, line, functionName := "?", 0 , "?"

	pc, fileName, line, ok := runtime.Caller(2)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		functionName = filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName, ".")
	}

	var msg = errorInfo {
		Time     : nowTime,
		Alarm    : method,
		Message  : str,
		Filename : fileName,
		Line     : line,
		Funcname : functionName,
	}

	jsons, _ := json.Marshal(msg)
	err = string(jsons)

	if method == "EMAIL" {
		// 执行发邮件

	} else if method == "SMS" {
		// 执行发短信

	} else if method == "WX" {
		// 执行发微信

	} else if method == "INFO" {
		// 执行记日志
	}
	return
}


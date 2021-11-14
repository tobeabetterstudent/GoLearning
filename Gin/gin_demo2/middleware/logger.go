// @program: gin_demo2的使用logrus实现自定义log中间件 将log输出到日志文件
// @author: aslanwang
// @create: 2021-11-14
package middleware

import (
	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"GoLearning/Gin/gin_demo2/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"path"
	"os"
	"fmt"
	"time"
)

// LoggerToFile 将日志记录到文件 
// 返回一个gin.HandleFunc的函数方法 即func (c *gin.Context)
func LoggerToFile() gin.HandlerFunc {
	logPath := config.Log_FILE_PATH
	logName := config.LOG_FILE_NAME
	// 合成日志文件目录
	logFile := path.Join(logPath, logName)
	src, err := os.OpenFile(logFile, os.O_CREATE | os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("OpenLogFile or CreateLogFile Error!")
		return nil
	}

	logger := logrus.New()
	logger.Out = src					// 设置输出
	logger.SetLevel(logrus.DebugLevel)	// 设置日志级别
	// 引入file-rotatelogs来支持日志按文件生成，且按天分割；rotatelogs实现了io.Writer 
	infoLogWriter, err := rotatelogs.New(
		// 分割后的文件名称
		logFile + ".%Y%m%d.infoLog",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(logFile),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		fmt.Println("Rotatelogs Call New() Error!")
		return nil
	}
	ErrorLogWriter, err := rotatelogs.New(
		// 分割后的文件名称
		logFile + ".%Y%m%d.errorLog",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(logFile),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		fmt.Println("Rotatelogs Call New() Error!")
		return nil
	}
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  infoLogWriter,		// infoLogWriter订阅了Info Debug Warn级别的日志
		logrus.DebugLevel: infoLogWriter,
		logrus.WarnLevel:  infoLogWriter,
		logrus.FatalLevel: ErrorLogWriter,		// ErrorLogWriter订阅了Fatal Error Panic级别的日志
		logrus.ErrorLevel: ErrorLogWriter,
		logrus.PanicLevel: ErrorLogWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})
	// 新增 Hook
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code"  : statusCode,
			"latency_time" : latencyTime,
			"client_ip"    : clientIP,
			"req_method"   : reqMethod,
			"req_uri"      : reqUri,
		}).Info()
	}
}


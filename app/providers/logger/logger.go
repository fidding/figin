package logger

import (
	"figin/config"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// Fields 日志字段
type Fields map[string]interface{}

// GetLogger 获取log实例
func GetLogger() *logrus.Logger {
	return logger
}

// Setup 初始化
func Setup() {
	logger = logrus.New()
	conf := config.GetConf()
	// 设置日志级别
	switch strings.ToLower(conf.Logger.Level) {
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	}

	// 设置日志样式
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// 设置日志输出
	logType := conf.Logger.LogType
	switch logType {
	case "stdout":
		logger.Out = os.Stdout
	case "file":
		fmt.Println("输出到日志")
		logfile := "logs/figin.log"
		f, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			log.Fatalf("Read 'logs/figin.log' fail: %v\n", err)
		}
		logger.Out = f
	default:
		logger.Out = os.Stdout
	}
	logger.Info("logger init")
}

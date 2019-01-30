package logger

import (
	"figin/config"
	"io"
	"log"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

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
	var logOutPut io.Writer
	switch logType {
	case "stdout":
		logOutPut = os.Stdout
	case "file":
		logfile := "logs/figin.log"
		f, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			log.Fatalf("Read 'logs/figin.log' fail: %v\n", err)
		}
		logOutPut = f
	}
	logger.SetOutput(logOutPut)
}

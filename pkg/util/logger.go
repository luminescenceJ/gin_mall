package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var LogrusObj *logrus.Logger

func init() {
	src, _ := setOutputFIle()
	if LogrusObj == nil {
		LogrusObj.Out = src
		return
	}
	//实例化
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LogrusObj = logger
}

func setOutputFIle() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/log/s"
	}
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		err = os.Mkdir(logFilePath, 0777)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		if err = os.Mkdir(logFilePath, 0777); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	return src, nil
}

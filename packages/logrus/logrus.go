package main

import (
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logrus.WithFields(logrus.Fields{
		"hello": "world",
	}).Info("welcome")
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		PadLevelText:    true,
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.DateTime,
	})
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.Trace("找一找BUG")
	logrus.Debug("有BUG！！！")
	logrus.Info("没BUG！！！")
	logrus.Warn("小心BUG")
	logrus.Fatal("BUG发作")
	//logrus.Panic("BUG无敌了")
}

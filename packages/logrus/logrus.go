/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/8/13 12:15
 * @version     v1.0
 * @filename    logrus.go
 * @description 日志
 ***************************************************************************/
package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.WithFields(logrus.Fields{
		"hello": "world",
	}).Info("welcome")
}

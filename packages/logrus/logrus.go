package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.WithFields(logrus.Fields{
		"hello": "world",
	}).Info("welcome")
}

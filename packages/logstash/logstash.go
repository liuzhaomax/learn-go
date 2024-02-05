package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
)

type LogstashHook struct {
	Host string
	Port int
}

func NewLogstashHook(host string, port int) *LogstashHook {
	return &LogstashHook{
		Host: host,
		Port: port,
	}
}

func (hook *LogstashHook) Fire(entry *logrus.Entry) error {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", hook.Host, hook.Port))
	if err != nil {
		return err
	}
	defer conn.Close()
	dataBytes, err := entry.Bytes()
	if err != nil {
		return err
	}
	_, err = conn.Write(dataBytes)
	return err
}

func (hook *LogstashHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	hook := NewLogstashHook("106.15.94.176", 5044)
	logger.AddHook(hook)

	logger.Info("Hello, Logstash!")
}

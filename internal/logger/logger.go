package logger

import "github.com/sirupsen/logrus"

func WithLevel(level string, msg string) {
	logrus.WithFields(logrus.Fields{
		"place": level,
	}).Println(msg)
}

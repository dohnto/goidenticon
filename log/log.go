package log

import "github.com/Sirupsen/logrus"

type Level struct {
	logrus.Level
}

func (l *Level) Set(arg string) error {
	level, err := logrus.ParseLevel(arg)
	*l = Level{level}
	return err
}

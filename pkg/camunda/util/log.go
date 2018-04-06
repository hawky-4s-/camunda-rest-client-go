package util

import "github.com/sirupsen/logrus"

type Logger struct {
	logger *logrus.Logger
}

func newLogger() *Logger {
	return &Logger{logger: logrus.New()}
}

func (l Logger) Debug(args ...interface{}) {
	l.logger.Debug(args)
}

func (l Logger) Debugln(args ...interface{}) {
	l.logger.Debugln(args)
}

func (l Logger) Info(args ...interface{}) {
	l.logger.Info(args)
}

func (l Logger) Infoln(args ...interface{}) {
	l.logger.Infoln(args)
}

func (l Logger) Warning(args ...interface{}) {
	l.logger.Warning(args)
}

func (l Logger) Warningln(args ...interface{}) {
	l.logger.Warningln(args)
}

func (l Logger) Error(args ...interface{}) {
	l.logger.Error(args)
}

func (l Logger) Errorln(args ...interface{}) {
	l.logger.Errorln(args)
}

func (l Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args)
}

func (l Logger) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args)
}

func (l Logger) Panic(args ...interface{}) {
	l.logger.Panic(args)
}

func (l Logger) Panicln(args ...interface{}) {
	l.logger.Panicln(args)
}

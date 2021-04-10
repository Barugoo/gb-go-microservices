package logger

import (
	"context"
	"gateway-user/reqdata"
	"io"

	log "github.com/sirupsen/logrus"
)

type Logger interface {
	Errorf(ctx context.Context, pattern string, args ...interface{})
	Warnf(ctx context.Context, pattern string, args ...interface{})
	Infof(ctx context.Context, pattern string, args ...interface{})
	Debugf(ctx context.Context, pattern string, args ...interface{})
	SetOutput(w io.Writer)
	Fatal(ctx context.Context, err error)
	Fatalf(ctx context.Context, pattern string, args ...interface{})
}

type logger struct {
	logrus *log.Logger
}

func NewLogger() Logger {
	l := &logger{
		logrus: log.StandardLogger(),
	}
	l.logrus.SetFormatter(&log.JSONFormatter{})
	return l
}

func (l *logger) Fatal(ctx context.Context, err error) {
	l.logrus.WithField("rid", reqdata.GetRequestID(ctx)).Fatal(err)
}

func (l *logger) Fatalf(ctx context.Context, pattern string, args ...interface{}) {
	l.logrus.WithField("rid", reqdata.GetRequestID(ctx)).Fatalf(pattern, args...)
}

func (l *logger) SetOutput(w io.Writer) {
	l.logrus.SetOutput(w)
}

func (l *logger) Errorf(ctx context.Context, pattern string, args ...interface{}) {
	l.logrus.WithField("rid", reqdata.GetRequestID(ctx)).Errorf(pattern, args...)
}

func (l *logger) Warnf(ctx context.Context, pattern string, args ...interface{}) {
	l.logrus.WithField("rid", reqdata.GetRequestID(ctx)).Warnf(pattern, args...)
}

func (l *logger) Debugf(ctx context.Context, pattern string, args ...interface{}) {
	l.logrus.WithField("rid", reqdata.GetRequestID(ctx)).Debugf(pattern, args...)
}

func (l *logger) Infof(ctx context.Context, pattern string, args ...interface{}) {
	l.logrus.WithField("rid", reqdata.GetRequestID(ctx)).Infof(pattern, args...)
}

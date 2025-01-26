package types

type ISuggarEngine interface {
	Warnf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Debugf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}

type ILogger interface {
	Warn(template string, args ...interface{})
	Info(template string, args ...interface{})
	Error(template string, args ...interface{})
	Debug(template string, args ...interface{})
	Fatal(template string, args ...interface{})
}

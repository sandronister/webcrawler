package ilog

type Type interface {
	Info(className, message string)
	Error(className, message string)
}

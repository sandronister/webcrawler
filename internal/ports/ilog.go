package ports

type ILog interface {
	Info(className, message string)
	Error(className, message string)
}

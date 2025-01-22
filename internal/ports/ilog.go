package ports

type ILog interface {
	Log(className, message string)
}

package logger

type ILogger interface {
	Info(string)
	Error(string, error)
	Trace(string)
	Warn(string)
	Debug(string)
}

type logger struct{}

func NewLogger() ILogger {
	return logger{}
}

func (l logger) Info(msg string) {
	GetLogger().Info(msg)
}

func (l logger) Error(msg string, err error) {
	GetLogger().Error(msg, err)
}

func (l logger) Trace(msg string) {
	GetLogger().Trace(msg)
}

func (l logger) Warn(msg string) {
	GetLogger().Warn(msg)
}

func (l logger) Debug(msg string) {
	GetLogger().Debug(msg)
}

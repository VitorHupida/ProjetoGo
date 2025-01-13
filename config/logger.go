package config

import (
	"io"
	"log"
	"os"
)

type Registrador struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NovoRegistrador(p string) *Registrador {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Registrador{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// cia logs não formatados
func (l *Registrador) Debug(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Registrador) Info(v ...interface{}) {
	l.info.Println(v...)
}
func (l *Registrador) Warn(v ...interface{}) {
	l.warning.Println(v...)
}
func (l *Registrador) Error(v ...interface{}) {
	l.err.Println(v...)
}

// cria logs formatados
func (l *Registrador) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Registrador) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}
func (l *Registrador) Warnf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}
func (l *Registrador) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}

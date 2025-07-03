package logger

import (
	"os"

	"erp-opity/pkg/config"

	"github.com/sirupsen/logrus"
)

// Logger interface para logging
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
}

// logger implementação do logger
type logger struct {
	*logrus.Logger
}

// New cria uma nova instância do logger
func New() Logger {
	return NewWithConfig(config.LoggingConfig{
		Level:  "info",
		Format: "json",
		Output: "stdout",
	})
}

// NewWithConfig cria uma nova instância do logger com configuração específica
func NewWithConfig(cfg config.LoggingConfig) Logger {
	log := logrus.New()

	// Configurar nível de log
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	log.SetLevel(level)

	// Configurar formato
	if cfg.Format == "json" {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	// Configurar saída
	if cfg.Output == "file" && cfg.FilePath != "" {
		file, err := os.OpenFile(cfg.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(file)
		}
	} else {
		log.SetOutput(os.Stdout)
	}

	return &logger{log}
}

// Debug logs a debug message
func (l *logger) Debug(args ...interface{}) {
	l.Logger.Debug(args...)
}

// Info logs an info message
func (l *logger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

// Warn logs a warning message
func (l *logger) Warn(args ...interface{}) {
	l.Logger.Warn(args...)
}

// Error logs an error message
func (l *logger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}

// Fatal logs a fatal message and exits
func (l *logger) Fatal(args ...interface{}) {
	l.Logger.Fatal(args...)
}

// Debugf logs a formatted debug message
func (l *logger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

// Infof logs a formatted info message
func (l *logger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

// Warnf logs a formatted warning message
func (l *logger) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

// Errorf logs a formatted error message
func (l *logger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

// Fatalf logs a formatted fatal message and exits
func (l *logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}

// WithField adds a field to the logger
func (l *logger) WithField(key string, value interface{}) Logger {
	return &logger{l.Logger.WithField(key, value).Logger}
}

// WithFields adds multiple fields to the logger
func (l *logger) WithFields(fields map[string]interface{}) Logger {
	return &logger{l.Logger.WithFields(logrus.Fields(fields)).Logger}
}

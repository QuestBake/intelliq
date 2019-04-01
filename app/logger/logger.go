package logger

import (
	log "github.com/go-ozzo/ozzo-log"
)

var logger *log.Logger

//InitLogger initializes the logger
func InitLogger(fileName string, maxSize int64, backupCount int) *log.Logger {
	logger = log.NewLogger()
	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	t2.FileName = fileName
	t2.MaxLevel = log.LevelError
	t2.MaxBytes = maxSize
	t2.Rotate = true
	t2.BackupCount = backupCount
	logger.Targets = append(logger.Targets, t1, t2)

	logger.Open()
	return logger
}

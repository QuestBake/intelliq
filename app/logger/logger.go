package logger

import (
	log "github.com/go-ozzo/ozzo-log"
)

//Logger globale instance of logger
var Logger *log.Logger

//InitLogger initializes the logger
func InitLogger(fileName string, maxSize int64, backupCount int) {
	Logger = log.NewLogger()
	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	t2.FileName = fileName
	t2.MaxLevel = log.LevelError
	t2.MaxBytes = maxSize
	t2.Rotate = true
	t2.BackupCount = backupCount
	Logger.Targets = append(Logger.Targets, t1, t2)

	Logger.Open()
}

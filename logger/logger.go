package logger

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

//InitLog make log for database query
func InitLog() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	return newLogger
}

package editor

import (
	"log"
	"os"
	"time"
)

const (
	LOG_FILE_EXTENSION = ".log"
	LOG_FILE_FLAGS     = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	LOG_FILE_MOD       = 0666
	LOG_FLAGS          = log.Ltime
	LOG_DIR            = "./logs/"
	LOG_FORMAT         = "2006-01-02_15:04:05"
	LOG_PREFIX_N       = 7  // first seven chars of each log denominates log type
	LOG_CONTENT_N      = 17 // from 17 char wach log has its content
)

// multiple layers) is called Tileseter
type Logger struct {
	filename     string
	path         string
	file         *os.File
	canvasLogger *log.Logger
}

// creates new Pallete struct
func NewLogger(path string) Logger {
	var logger Logger

	dt := time.Now()
	logger.filename = dt.Format(LOG_FORMAT) + LOG_FILE_EXTENSION
	logger.path = path + logger.filename

	file, err := os.OpenFile(logger.path, LOG_FILE_FLAGS, LOG_FILE_MOD)

	if err != nil {
		log.Fatal(err)
	}

	logger.file = file

	logger.canvasLogger = log.New(logger.file, "CANVAS: ", LOG_FLAGS)

	return logger
}

func (lg *Logger) LogCanvasPut(x int, y int, layer int) {
	lg.canvasLogger.Printf("Tile put on (%d, %d), on layer %d\n", x, y, layer)
}

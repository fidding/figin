package log

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func Logger(msg string) {
	log.WithFields(log.Fields{}).Info(msg)
}

func LogInit() {
	log.SetFormatter(&log.JSONFormatter{})
	filename := "logs/figin.log"
	f, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
	// log.SetLevel(log.WarnLevel)
	// log.WithFields(log.Fields{}).Warning("A walrus appears")
}
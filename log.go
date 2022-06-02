package textParser

import (
	"bytes"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func SetLog() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&MyFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}

type MyFormatter struct {
}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	newLog := fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}

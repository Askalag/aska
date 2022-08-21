package logging

import (
	"fmt"
	"github.com/askalag/aska/microservices/webapp/pkg"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

func Initial(c pkg.Logging) {
	logrus.SetReportCaller(true)

	switch c.Format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				fileName := path.Base(f.File)
				return fmt.Sprintf("%s:%d", fileName, f.Line), fmt.Sprintf("%s()", f.Function)
			},
			DisableColors: false,
			FullTimestamp: true,
		})
	}

	// log level
	lv, err := logrus.ParseLevel(c.Level)
	if err != nil {
		lv = logrus.DebugLevel
	} else {
		logrus.SetLevel(lv)
	}

	// log outputs
	file, err := os.OpenFile(c.FilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
}

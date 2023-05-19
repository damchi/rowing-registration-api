package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"time"
)

func CustomLogFormat(r *gin.Engine) {
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		url := param.Request.URL
		username := "-"
		if url.User != nil {
			if name := url.User.Username(); name != "" {
				username = name
			}
		}
		referer := "-"
		if param.Request.Referer() != "" {
			referer = param.Request.Referer()
		}
		// your custom format
		return fmt.Sprintf("%s %s - [%s] \"%s %s %s\" %d %d \"%s\" \"%s\" %d  %s\n",
			param.ClientIP,
			username,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.BodySize,
			referer,
			param.Request.UserAgent(),
			time.Since(param.TimeStamp).Nanoseconds()/1000,
			param.ErrorMessage,
		)
	}))
}

const (
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
	DEBUG   = "DEBUG"
	NOTICE  = "NOTICE"
)

type LogEntry struct {
	Message  interface{}
	Severity string
}

var LogCh = make(chan LogEntry)

func Log(severity string, message interface{}) {
	LogCh <- LogEntry{message, severity}
}

func Notice(message interface{}) {
	Log(NOTICE, message)
}

func Debug(message interface{}) {
	Log(DEBUG, message)
}

func Info(message interface{}) {
	Log(INFO, message)
}

func Error(message interface{}) {
	Log(ERROR, message)
}

func Warning(message interface{}) {
	Log(WARNING, message)
}

func InitLogger(logLevel string) {
	allowedLevel := []string{WARNING, ERROR, DEBUG, NOTICE, INFO}
	if !inArrayString(logLevel, allowedLevel) {
		log.Panic("Log level not supported")
	}

	levelValues := make(map[string]int)
	levelValues[DEBUG] = 100
	levelValues[INFO] = 200
	levelValues[NOTICE] = 250
	levelValues[WARNING] = 300
	levelValues[ERROR] = 400

	for {
		select {
		case entry := <-LogCh:
			strLevel := strings.ToUpper(logLevel)
			defaultSeverity := levelValues[strLevel]
			currentSeverity := levelValues[entry.Severity]
			if currentSeverity >= defaultSeverity {
				log.Println(fmt.Sprintf("- [%v] : %v", entry.Severity, entry.Message))
			}
		}
	}
}

func inArrayString(needle string, values []string) bool {
	for _, value := range values {
		if needle == value {
			return true
		}
	}
	return false
}

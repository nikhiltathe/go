// This package implements logging to the file. Logs are also rotated.

// To use the logging facility
// call log.Init(LogFile_name,log_level,display_console_bool) in main()
// if need to log console provide bool as true
// log.Info ("log infomation message") use log.levels as per requirement
// Currently supports 4 logging levels(ERROR,WARNING,INFO,DEBUG)
// Logs will be placed in:
//  ./log/servce_name/svcname.log  svcname_json.log
// logs files will be rotated based on file size defined

package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type configLogrus struct {
	IsInitDone    bool
	isLogToStdout bool
	lrusJSON      *logrus.Logger
	lrusTxt       *logrus.Logger
	lumberJSON    *lumberjack.Logger
	lumberTxt     *lumberjack.Logger
	UsvcName      string
}
type PlainFormatter struct{}

var Logging configLogrus

const (
	MAX_FILE_SIZE_MB = 50
	MAX_BACKUPS      = 10                        // Keep last 10 log files
	LOG_TIME_FORMAT  = "2006-01-02 15:04:05.000" // 20060102150405 stands for yyyymmdd 15:04:05.123
	FILE             = "file"                    //To display file name
	FUNC             = "function"                //used to display function name
	LINE             = "line"                    //used to display line number
)

const (
	LEVEL_ERROR = iota
	LEVEL_WARN
	LEVEL_INFO  //min & default
	LEVEL_DEBUG //max
)

const (
	LOG_ERROR  = "ERROR"
	LOG_ERRORF = "ERROR_F"
	LOG_WARN   = "WARN"
	LOG_WARNF  = "WARN_F"
	LOG_INFO   = "INFO"
	LOG_INFOF  = "INFO_F"
	LOG_DEBUG  = "DEBUG"
	LOG_DEBUGF = "DEBUG_F"
)
const (
	LOG_DIR         = "./log/"
	LOG_JSON_FORMAT = "_json.log"
	LOG_TXT_FORMAT  = ".log"
)

type CRLogger struct {
	Setlogpath  string
	Level       int
	IsStdOutput bool
}

// Caller is responsible to call this at the start and after DeInit.
var once sync.Once

func Initialize(logname string, crLoggerInfo CRLogger) {

	once.Do(func() {
		plainFormatter := new(PlainFormatter)
		Logging.UsvcName = logname
		var logfile string
		// if user provides a fullpath use it
		if crLoggerInfo.Setlogpath != "" {
			logdir := crLoggerInfo.Setlogpath + logname + "/"
			os.Mkdir(logdir, 0740)
			logfile = logdir + logname
		} else {
			logdir := LOG_DIR + logname + "/"
			os.Mkdir(logdir, 0740)
			logfile = logdir + logname
		}

		Logging.UsvcName = logname
		Logging.lrusJSON = logrus.New()
		Logging.lrusTxt = logrus.New()
		//	Use lumberjack for log rotation. Assign it as output to logrus
		Logging.lumberJSON = &lumberjack.Logger{
			Filename:   logfile + LOG_JSON_FORMAT,
			MaxSize:    MAX_FILE_SIZE_MB,
			MaxBackups: MAX_BACKUPS,
			Compress:   true,
		}
		Logging.lrusJSON.Out = Logging.lumberJSON
		Logging.lumberTxt = &lumberjack.Logger{
			Filename:   logfile + LOG_TXT_FORMAT,
			MaxSize:    MAX_FILE_SIZE_MB,
			MaxBackups: MAX_BACKUPS,
			Compress:   true,
		}
		var mw io.Writer
		if crLoggerInfo.IsStdOutput {
			mw = io.MultiWriter(os.Stdout, Logging.lumberTxt)
		} else {
			mw = io.Writer(Logging.lumberTxt)
		}
		Logging.lrusTxt.Out = mw
		//set Format to logging
		Logging.lrusTxt.Formatter = plainFormatter
		Logging.lrusJSON.Formatter = &logrus.JSONFormatter{
			TimestampFormat: LOG_TIME_FORMAT,
		}
		SetLevel(crLoggerInfo.Level)

	})

}

func Init(logname string, level int, isStdOutput bool) {

	var loggerInfo CRLogger

	loggerInfo.Setlogpath = ""
	loggerInfo.Level = level
	loggerInfo.IsStdOutput = isStdOutput

	Initialize(logname, loggerInfo)

}

// Close is used stop logging
func Close() {
	if Logging.IsInitDone == true {
		Logging.lumberJSON.Close()
		Logging.lrusJSON = nil
		Logging.lumberJSON = nil
		Logging.lumberTxt.Close()
		Logging.lrusTxt = nil
		Logging.lumberTxt = nil
		Logging.IsInitDone = false
	}
}

// SetLevel sets log level. Log will be printed for level less than or equal to value set.
// debug > info > warn > error
//default is info
func SetLevel(level int) {
	Infof("Setting LogLevel to %d", level)
	switch level {
	case LEVEL_ERROR:
		Logging.lrusJSON.Level, Logging.lrusTxt.Level = logrus.ErrorLevel, logrus.ErrorLevel
	case LEVEL_WARN:
		Logging.lrusJSON.Level, Logging.lrusTxt.Level = logrus.WarnLevel, logrus.WarnLevel
	case LEVEL_DEBUG:
		Logging.lrusJSON.Level, Logging.lrusTxt.Level = logrus.DebugLevel, logrus.DebugLevel
	case LEVEL_INFO:
		fallthrough
	default:
		Logging.lrusJSON.Level, Logging.lrusTxt.Level = logrus.InfoLevel, logrus.InfoLevel
	}
}

// SetLevel sets log level. Log will be printed for level less than or equal to value set.
// debug > info > warn > error
//default is info
func SetLogLevel(level string) {
	Infof("Setting LogLevel to %s", strings.ToUpper(level))

	switch strings.ToUpper(level) {
	case LOG_ERROR:
		Logging.lrusJSON.Level, Logging.lrusTxt.Level = logrus.ErrorLevel, logrus.ErrorLevel
	case LOG_WARN:
		Logging.lrusJSON.Level, Logging.lrusTxt.Level = logrus.WarnLevel, logrus.WarnLevel
	case LOG_DEBUG:
		Logging.lrusJSON.Level, Logging.lrusTxt.Level = logrus.DebugLevel, logrus.DebugLevel
	case LOG_INFO:
		fallthrough
	default:
		Logging.lrusJSON.Level, Logging.lrusTxt.Level = logrus.InfoLevel, logrus.InfoLevel
	}
}

// GetLevel get the current logging level
func GetLevel() string {

	switch Logging.lrusJSON.Level {
	case logrus.ErrorLevel:
		return LOG_ERROR
	case logrus.WarnLevel:
		return LOG_WARN
	case logrus.DebugLevel:
		return LOG_DEBUG
	case logrus.InfoLevel:
		return LOG_INFO
	}
	return LOG_INFO
}

// callstack returns file and line three stack frames above its invocation.
//   Optional stack level argument if the user wants a different stack level

func callstack(stackLevel ...int) (funcname string, fileName string, lineNum int) {
	//   By default we print the third level of the stack. Some functions may want to
	// chose a different level
	level := 3
	if len(stackLevel) > 0 {
		level = stackLevel[0]
	}

	if pc, file, line, ok := runtime.Caller(level); ok {
		funcName := runtime.FuncForPC(pc).Name()
		if funcName == "" {
			funcName = "?()"
		} else {
			dotName := filepath.Ext(funcName)
			funcName = strings.TrimLeft(dotName, ".") + "()"
		}
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
		funcname, fileName, lineNum = funcName, file, line
	}
	return funcname, fileName, lineNum
}

// Various functions that can be used to log.
func (f *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := fmt.Sprintf(entry.Time.Format(LOG_TIME_FORMAT))
	LevelDesc := []string{LOG_ERROR, LOG_WARN, LOG_INFO, LOG_DEBUG} //This array is to replace the logrus log_level name by custom name(eg:- logrus warnning to WARN).Currently we are supporting only 4 log levels(not supporting PANIC and FATAL) so  array index is reduced by 2
	return []byte(fmt.Sprintf("[%s] [%s] [%s] [%s:%d %s] : %s\n", timestamp, LevelDesc[entry.Level-2], Logging.UsvcName, entry.Data[FILE], entry.Data[LINE], entry.Data[FUNC], entry.Message)), nil
}

func defineFormat(level string, format string, args ...interface{}) {
	funcName, file, line := callstack()
	allLoggers := []*logrus.Logger{Logging.lrusTxt, Logging.lrusJSON}
	for _, v := range allLoggers {
		currentLogger := v.WithFields(logrus.Fields{
			FILE: file,
			LINE: line,
			FUNC: funcName,
		})

		switch level {
		case LOG_INFO:
			currentLogger.Info(args...)
		case LOG_INFOF:
			currentLogger.Infof(format, args...)
		case LOG_DEBUG:
			currentLogger.Debug(args...)
		case LOG_DEBUGF:
			currentLogger.Debugf(format, args...)
		case LOG_WARN:
			currentLogger.Warn(args...)
		case LOG_WARNF:
			currentLogger.Warnf(format, args...)
		case LOG_ERROR:
			currentLogger.Error(args...)
		case LOG_ERRORF:
			currentLogger.Errorf(format, args...)
		}
	}

}

func Info(args ...interface{}) {
	defineFormat(LOG_INFO, "", args...)
}

func Infof(format string, args ...interface{}) {
	defineFormat(LOG_INFOF, format, args...)
}

func Debug(args ...interface{}) {
	defineFormat(LOG_DEBUG, "", args...)
}

func Debugf(format string, args ...interface{}) {
	defineFormat(LOG_DEBUGF, format, args...)
}

func Warn(args ...interface{}) {
	defineFormat(LOG_WARN, "", args...)
}

func Warnf(format string, args ...interface{}) {
	defineFormat(LOG_WARNF, format, args...)
}

func Error(args ...interface{}) {
	defineFormat(LOG_ERROR, "", args...)
}

func Errorf(format string, args ...interface{}) {
	defineFormat(LOG_ERRORF, format, args...)
}

func CRStruct(format string, src interface{}, out interface{}) {
	if src != nil {
		jsonData, err := json.Marshal(&src)
		if err != nil {
			return
		}
		json.Unmarshal(jsonData, &out)
	}
	defineFormat(LOG_DEBUGF, format, out)
}

package xlog

import (
	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
)

type LogrusLogger struct {
	logger        *logrus.Logger
	dateFormat    string
	fields        map[string]interface{}
	displayFields bool
	class         string
}

type LogOptions struct {
	LogPath         string `mapstructure:"log_path"`
	LogLevel        string `mapstructure:"log_level"` // trace, debug, info, warn[ing], error, fatal, panic
	LogMaxDiskUsage int64  `mapstructure:"log_max_disk_usage"`
	LogMaxFileNum   int    `mapstructure:"log_max_file_num"`
	appName         string `mapstructure:"app_name"`
}

func GetXLogger(class string) ILogger {
	configViper := viper.New()
	configViper.SetConfigFile("./log.yml")
	err := configViper.ReadInConfig()
	var option *LogOptions
	if err == nil {
		err = configViper.Sub("yuanboot.log").Unmarshal(&option)
	}
	if err != nil {
		option = &LogOptions{LogLevel: "debug", LogPath: "./log", LogMaxDiskUsage: 102400000, LogMaxFileNum: 50, appName: "log"}
	}
	logger := GetClassLogger(class, option) // NewXLogger()
	return logger
}

func GetXLoggerWithFields(class string, fields map[string]interface{}) ILogger {
	logger := NewXLogger()
	logger.class = class
	return logger
}

func GetXLoggerWith(logger ILogger) ILogger {
	return logger
}

func NewLogger(options *LogOptions) ILogger {
	logger := logrus.New()
	lw := &HourlySplit{
		Dir:           options.LogPath,
		FileFormat:    options.appName + "_2006-01-02T15",
		MaxFileNumber: int64(options.LogMaxFileNum),
		MaxDiskUsage:  options.LogMaxDiskUsage,
	}
	multiWriter := io.MultiWriter(os.Stdout, lw)
	defer lw.Close()
	logger.SetReportCaller(true)
	logger.SetOutput(multiWriter)
	lv, err := logrus.ParseLevel(options.LogLevel)
	if err != nil {
		lv = logrus.WarnLevel
	}
	logger.SetLevel(lv)
	return &LogrusLogger{logger: logger, dateFormat: LoggerDefaultDateFormat}
}

func GetClassLogger(class string, options *LogOptions) ILogger {
	logger := logrus.New()
	lw := &HourlySplit{
		Dir:           options.LogPath,
		FileFormat:    options.appName + "_2006-01-02T15",
		MaxFileNumber: int64(options.LogMaxFileNum),
		MaxDiskUsage:  options.LogMaxDiskUsage,
	}
	defer lw.Close()
	multiWriter := io.MultiWriter(os.Stdout, lw)
	logger.SetReportCaller(true)
	logger.SetOutput(multiWriter)
	lv, err := logrus.ParseLevel(options.LogLevel)
	if err != nil {
		lv = logrus.WarnLevel
	}
	logger.SetLevel(lv)
	logger.Formatter = &TextFormatter{
		DisableColors:   false,
		ForceColors:     false,
		TimestampFormat: LoggerDefaultDateFormat,
		FullTimestamp:   true,
		ForceFormatting: true,
	}
	return &LogrusLogger{logger: logger, class: class, dateFormat: LoggerDefaultDateFormat, displayFields: true}
}

func (log *LogrusLogger) With(level LogLevel, fiedls map[string]interface{}) *logrus.Entry {

	//start := time.Now()

	fieldsMap := make(map[string]interface{})
	fieldsMap["prefix"] = "yuanboot"
	if fiedls != nil {
		fieldsMap = fiedls
	}

	if log.displayFields {
		fieldsMap["class"] = log.class
		hostName, _ := os.Hostname()
		fieldsMap["host"] = hostName
	}
	//fieldsMap["message"] = message

	return log.logger.WithFields(fieldsMap)
}

func (log LogrusLogger) Debug(format string, a ...interface{}) {
	log.With(DEBUG, log.fields).Debugf(format, a...)
}

func (log LogrusLogger) Info(format string, a ...interface{}) {
	log.With(INFO, log.fields).Infof(format, a...)
}

func (log LogrusLogger) Warning(format string, a ...interface{}) {
	log.With(WARNING, log.fields).Warnf(format, a...)
}

func (log LogrusLogger) Error(format string, a ...interface{}) {
	log.logger.Out = os.Stderr
	log.With(ERROR, log.fields).Errorf(format, a...)
	log.logger.Out = os.Stdout
}

func (log *LogrusLogger) SetClass(className string) {
	log.class = className
}

func (log *LogrusLogger) SetCustomLogFormat(logFormatterFunc func(logInfo interface{}) string) {
	log.displayFields = false
}

func (log LogrusLogger) SetDateFormat(format string) {
	//log.dateFormat = format
}

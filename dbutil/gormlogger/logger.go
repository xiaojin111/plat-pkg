package gormlogger

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"time"

	"gitee.com/jt-heath/go-pkg/v2/log"
	"github.com/sirupsen/logrus"
)

// GormLogger 面向 Gorm 的 Logger
type GormLogger struct {
	DBName string
	DBHost string
	logger logrus.FieldLogger
	level  log.Level
}

// New Create new logger with custom database name and host
// Defaults to use logrus.INFO level.
func New(host, name string) *GormLogger {
	return NewWithLevel(host, name, log.GetLevel())
}

// NewWithLevel Create new logger with custom database name, host and logurs.Level
func NewWithLevel(host, name string, level log.Level) *GormLogger {
	return NewWithLogger(host, name, log.StandardLogger(), level)
}

// NewWithLogger Create new logger with custom database name, host and logger
func NewWithLogger(host, name string, logger *log.Logger, level log.Level) *GormLogger {
	return &GormLogger{
		DBHost: host,
		DBName: name,
		logger: logger,
		level:  level,
	}
}

var sqlRegexp = regexp.MustCompile(`(\$\d+)|\?`)

// Print 打印日志
func (l *GormLogger) Print(values ...interface{}) {
	entry := l.logger.
		WithField("dbhost", l.DBHost).
		WithField("dbname", l.DBName)

	if len(values) > 1 {
		level := values[0]
		source := values[1]
		entry = entry.WithField("src", source)
		if level == "sql" {
			// sql: values
			//  [sql main.go:51 687.158µs select * from client where 1=1 or 1=? or 'a'=?; [2 A] 0]
			duration := values[2]
			var formattedValues []interface{}
			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format(time.RFC3339)))
					} else if b, ok := value.([]byte); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", string(b)))
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
					}
				} else {
					formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
				}
			}
			entry.WithField("latency", duration).
				Log(l.level, fmt.Sprintf(sqlRegexp.ReplaceAllString(values[3].(string), "%v"), formattedValues...))
		} else {
			entry.Error(values[2:]...)
		}
	} else {
		entry.Error(values...)
	}

}

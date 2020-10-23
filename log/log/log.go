package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	log "github.com/sirupsen/logrus"
)

const defaultTimestampFormat = "2006-01-02 15:04:05"
const defaultDateTimestampFormat = "2006-01-02"

var (
	// ErrNoTimeField 没有时间字段
	ErrNoTimeField = errors.New("no time field")
	// ErrInvalidTime 错误的时间
	ErrInvalidTime = errors.New("invalid time")
)

// Out 输出日志器
// 按时间 天创建目录 每小时生成文件
type Out struct {
	BaseDir string
}

// Write 只支持写入json格式的文件
func (o *Out) Write(p []byte) (n int, err error) {
	var v log.Fields
	err = json.Unmarshal(p, &v)
	if err != nil {
		return
	}
	timeVal, exist := v["time"]
	if !exist {
		err = ErrNoTimeField
		return
	}
	realTime, ok := timeVal.(string)
	if !ok {
		err = ErrInvalidTime
		return
	}
	t, err := time.Parse(defaultTimestampFormat, realTime)
	if err != nil {
		return
	}
	day := t.Format(defaultDateTimestampFormat)
	fmt.Println(day)
	pt := path.Join(o.BaseDir, day)
	err = os.Mkdir(pt, 0766)
	return
}

// NewLogger 新建日志记录器
func NewLogger(out io.Writer) *log.Logger {
	l := log.New()
	l.SetFormatter(&log.JSONFormatter{
		TimestampFormat: defaultTimestampFormat,
	})
	l.SetOutput(out)
	return l
}

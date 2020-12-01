package log

import (
	"io"
	"os"
	"path"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	defaultFileExtName         = ".log"
	defaultBaseDir             = "/tmp"
	defaultTimestampFormat     = "2006-01-02 15:04:05"
	defaultHourTimeStampFormat = "15"
	defaultDateTimestampFormat = "2006-01-02"
)

// Out 输出日志器
// 按时间 天创建目录 每小时生成文件
type Out struct {
	BaseDir     string
	FileExtName string
}

// NewOut 新建输出
func NewOut(baseDir string, fileExtName string) *Out {
	if baseDir == "" {
		baseDir = defaultBaseDir
	}
	if fileExtName == "" {
		fileExtName = defaultFileExtName
	}
	return &Out{
		BaseDir:     baseDir,
		FileExtName: fileExtName,
	}
}

// Write
func (o *Out) Write(p []byte) (n int, err error) {
	now := time.Now()
	if err != nil {
		return
	}
	day := now.Format(defaultDateTimestampFormat)
	pt := path.Join(o.BaseDir, day)
	s, _ := os.Stat(pt)
	if s == nil {
		if err = os.Mkdir(pt, 0766); err != nil {
			return
		}
	}
	filePath := path.Join(pt, now.Format(defaultHourTimeStampFormat))
	f, err := os.OpenFile(filePath+o.FileExtName, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	n, err = f.Write(p)
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

// New new logger
func New(baseDir string) *log.Logger {
	l := log.New()
	l.SetFormatter(&log.JSONFormatter{
		TimestampFormat: defaultTimestampFormat,
	})
	l.SetOutput(&Out{
		BaseDir: baseDir,
	})
	return l
}

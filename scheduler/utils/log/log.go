package log

import (
	"strings"

	"github.com/astaxie/beego/logs"
)

var Log *logs.BeeLogger

func init() {
	Log = logs.NewLogger(10000)
	Log.EnableFuncCallDepth(true)
	Log.SetLogFuncCallDepth(3)
	/*默认打印日志到控制台*/
	if err := Log.SetLogger(logs.AdapterConsole, ``); err != nil {
		panic(err)
	}
}

/*
设置落地日志文件的方式
eg:
fileLogConfig = `{"filename":"logs/cgi.log","daily":true}`
fileLogConfig = `{"filename":"/data/logs/cgi.log","daily":false,"maxsize":4094967296}`
参数含义:
filename 保存的文件名
maxlines 每个文件保存的最大行数，默认值 1000000
maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
daily 是否按照每天 logrotate，默认是 true
maxdays 文件最多保存多少天，默认保存 7 天
rotate 是否开启 logrotate，默认是 true
level 日志保存的时候的级别，默认是 Trace 级别
perm: 日志文件权限
*/
func SetFileLogConfig(fileLogConfig string) {
	if err := Log.SetLogger(logs.AdapterFile, fileLogConfig); err != nil {
		panic(err)
	}
	logs.EnableFuncCallDepth(true)
}

func SetESLogConfig(esConfig string) {
	if err := Log.SetLogger(logs.AdapterEs+"m", esConfig); err != nil {
		panic(err)
	}
}

/*
需要自己指定打印的格式
Critical("%s%d%v","111",222,map[string]string{})
*/
func Critical(format string, v ...interface{}) { Log.Critical(format, v...) }

func Error(format string, v ...interface{}) { Log.Error(format, v...) }

func Warn(format string, v ...interface{}) { Log.Warn(format, v...) }

func Notice(format string, v ...interface{}) { Log.Notice(format, v...) }

func Trace(format string, v ...interface{}) { Log.Trace(format, v...) }

func Debug(format string, v ...interface{}) { Log.Debug(format, v...) }

func Info(format string, v ...interface{}) { Log.Informational(format, v...) }

/*
lazy打印方式,相当于全部用%v来打印参数
CriticalLazyCritical("111",222,map[string]string{})
相当于
Critical("%v%v%v","111",222,map[string]string{})
*/
func CriticalLazy(v ...interface{}) {
	format := strings.Repeat("%+v ", len(v))
	Log.Critical(format, v...)
}

func ErrorLazy(v ...interface{}) {
	format := strings.Repeat("%+v ", len(v))
	Log.Error(format, v...)
}

func WarnLazy(v ...interface{}) {
	format := strings.Repeat("%+v ", len(v))
	Log.Warn(format, v...)
}

func NoticeLazy(v ...interface{}) {
	format := strings.Repeat("%+v ", len(v))
	Log.Warn(format, v...)

}

func TraceLazy(v ...interface{}) {
	format := strings.Repeat("%+v ", len(v))
	Log.Warn(format, v...)
}

func DebugLazy(v ...interface{}) {
	format := strings.Repeat("%+v ", len(v))
	Log.Debug(format, v...)
}

func InfoLazy(v ...interface{}) {
	format := strings.Repeat("%+v ", len(v))
	Log.Informational(format, v...)
}

func Flush() { Log.Flush() }

func Close() { Log.Close() }

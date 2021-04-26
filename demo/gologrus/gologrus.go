package gologrus

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 日志格式化为JSON而不是默认的ASCII
	log.SetFormatter(&log.JSONFormatter{})

	// 输出 stdout 而不是默认的 stderr, 也可以是一个文件
	log.SetOutput(os.Stdout)

	// 只记录严重或以上警告
	log.SetLevel(log.WarnLevel)
}

func GlogrusMain() {

	/*log.WithFields(log.Fields{
		"animal": "wlarus",
	}).Info("A walrus appears")*/

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number incresed tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// 通过日志语句重用字段
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})
	contextLogger.Info("I'll be logged with common and other field !")
	contextLogger.Info("Me too!")
}

package utilities

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"time"
)

func LogInfo(format string, value ...interface{}) {
	pc, fPath, line, _ := runtime.Caller(1)
	funcN := runtime.FuncForPC(pc).Name()
	_, fName := filepath.Split(fPath)
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s INFO: %s:%s:%d - ", time.Now().Format("2006-02-01 15:04:05 Monday"), fName, funcN, line))
	log.Printf(format, value...)
}

func LogError(format string, value ...interface{}) {
	pc, fPath, line, _ := runtime.Caller(1)
	funcN := runtime.FuncForPC(pc).Name()
	_, fName := filepath.Split(fPath)
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s Error: %s:%s:%d - ", time.Now().Format("2006-02-01 15:04:05 Monday"), fName, funcN, line))
	log.Printf(format, value...)
}

func LogWarn(format string, value ...interface{}) {
	pc, fPath, line, _ := runtime.Caller(1)
	funcN := runtime.FuncForPC(pc).Name()
	_, fName := filepath.Split(fPath)
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s WARN: %s:%s:%d - ", time.Now().Format("2006-02-01 15:04:05 Monday"), fName, funcN, line))
	log.Printf(format, value...)
}

func LogFatal(format string, value ...interface{}) {
	pc, fPath, line, _ := runtime.Caller(1)
	funcN := runtime.FuncForPC(pc).Name()
	_, fName := filepath.Split(fPath)
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s FATAL: %s.%s:%d - ", time.Now().Format("2006-02-01 15:04:05 Monday"), fName, funcN, line))
	log.Fatal(fmt.Sprintf(format, value...))
}

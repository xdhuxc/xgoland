package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	logFileName := os.Getenv("LOG_FILE_NAME")
	if logFileName == "" {
		logFileName = "xdhuxc"
	}
	// 创建日志存储文件
	// 日期的格式化必须使用 2006-01-02 15:04:05.999999999 -0700 MST，否则会出错。
	logFile, err := os.Create(logFileName + "_" + time.Now().Format("2006-01-02") + ".log")
	if err != nil {
		fmt.Println(err)
	}

	// 创建一个 Logger
	logger := log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)

	// 输出一条日志
	logger.Output(2, "It is a Output log!")
	logger.Println("It is a Println log!")
}

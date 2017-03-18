package util

import (
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
)

//是否是windows
func IsWin() bool {
	var envOs = os.Getenv("OS")
	return envOs != "" && strings.Contains(envOs, "Windows")
}

//获取执行程序所在的路径
func GetRunDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//打印错误
func Log(err error) {
	if err != nil {
		log.Println(err)
		debug.PrintStack()
	}
}

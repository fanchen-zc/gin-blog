package logging

import (
	"fmt"
	"gin-test/pkg/file"
	"gin-test/pkg/setting"
	"os"
	"time"
)

//var (
//	LogSavePath = "runtime/logs/"
//	LogSaveName = "log"
//	LogFileExt  = "log"
//	TimeFormat  = "20060102"
//)

func getLogFilePath() string {
	//return fmt.Sprintf("%s", setting.AppSetting.LogSavePath)
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

//获取log文件完整路径
//func getLogFileFullPath() string {
//	prefixPath := getLogFilePath()
//	suffixPath := fmt.Sprintf("%s%s.%s", setting.AppSetting.LogSaveName, time.Now().Format(setting.AppSetting.TimeFormat), setting.AppSetting.LogFileExt)
//	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
//}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

// 打开文件
//
//	func openLogFile(filePath string) *os.File {
//		_, err := os.Stat(filePath)
//		switch {
//		case os.IsNotExist(err):
//			mkDir()
//		case os.IsPermission(err):
//			log.Fatalf("Permission :%v", err)
//		}
//		handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//		if err != nil {
//			log.Fatalf("Fail to OpenFile :%v", err)
//		}
//
//		return handle
//	}
func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}

//创建目录
//func mkDir() {
//	dir, _ := os.Getwd()
//	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
//	if err != nil {
//		panic(err)
//	}
//}

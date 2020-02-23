package common

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/astaxie/beego/logs"
)

var (
	customerLogger *logs.BeeLogger
)

const (
	DEFAULT_LOG_DIR string = "./logs"
)

func InitLogger() (err error) {
	customerLogger, err = InitCustomerLogger()
	if err != nil {
		return
	}
	return
}

//用于获取封装的logger，该logger默认打印调用上级的文件名
func InitCustomerLogger() (*logs.BeeLogger, error) {
	logger := logs.NewLogger()
	//默认打印输出行号和文件名
	logger.EnableFuncCallDepth(true)
	//默认打印上级调用
	logger.SetLogFuncCallDepth(3)

	if LogDir == "" {
		LogDir = DEFAULT_LOG_DIR
	}
	//create log_dir
	if err := createLogDir(LogDir); err != nil {
		fmt.Println(fmt.Sprintf("create dir %s error %s", LogDir, err))
		return logger, err
	}
	err := logger.SetLogger(logs.AdapterConsole, `{"level":8}`)
	if err != nil {
		fmt.Println(err.Error())
		return logger, err
	}
	logPath := LogDir + "/" + LogFileName
	config := make(map[string]interface{})
	config["filename"] = logPath
	config["level"] = logs.LevelNotice

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println(fmt.Sprintf("marshal failed, err: %s", err))
		return logger, err
	}
	if err := logger.SetLogger(logs.AdapterFile, string(configStr)); err != nil {
		fmt.Println(err.Error())
		return logger, err
	}
	if err != nil {
		fmt.Println(err.Error())
		return logger, err
	}
	return logger, nil
}

func createLogDir(log_dir string) error {
	if _, err := os.Stat(log_dir); err != nil {
		return os.MkdirAll(log_dir, 0777)
	}
	return nil
}

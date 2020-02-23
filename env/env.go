//used to initialize env configurations
package env

import (
	"github.com/astaxie/beego/logs"
	"github.com/oam-backend/common"
)

//control the initial sequence
func InitEnv() error {
	common.InitConfig()
	err := common.InitLogger()
	if err != nil {
		logs.Error("Logger init failed! err:%v", err.Error())
		return err
	}
	return nil
}

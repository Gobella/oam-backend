package common

import "github.com/astaxie/beego"

var (
	LogDir      string
	LogFileName string

	KubeAPIVersion string
)

func InitConfig() {
	LogFileName = beego.AppConfig.String("log_file_name")
	KubeAPIVersion = beego.AppConfig.String("kube_api_version")
}

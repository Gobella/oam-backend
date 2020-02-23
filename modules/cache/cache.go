package cache

import (
	"fmt"
	"github.com/oam-backend/common"
	"github.com/oam-backend/models"
	"sync"
	"time"
)

// used to cache the data instead of using a DB
var (
	ApplicationCache sync.Map
	YamlCache        sync.Map
	ChangeSheetCache sync.Map
)

func StoreApp(info models.ApplicationMetaInfo) (resp common.Response) {
	//app's name should be uniq in this demo
	ApplicationCache.Store(info.Name, info)
	return
}

func GetApps() (resp common.Response) {
	apps := make([]models.ApplicationMetaInfo, 0)
	ApplicationCache.Range(func(key, value interface{}) bool {
		valueInfo := value.(models.ApplicationMetaInfo)
		apps = append(apps, valueInfo)
		return true
	})
	resp.Data = apps
	return
}

func StoreChangeSheet(info models.ChangeSheet) (resp common.Response) {
	//app's name should be uniq in this demo
	key := fmt.Sprintf("%v-%v", info.App.Name, time.Now().Unix())
	ChangeSheetCache.Store(key, info)
	return
}

func GetChangeSheet() (resp common.Response) {
	sheets := make([]models.ChangeSheet, 0)
	ApplicationCache.Range(func(key, value interface{}) bool {
		valueInfo := value.(models.ChangeSheet)
		sheets = append(sheets, valueInfo)
		return true
	})
	resp.Data = sheets
	return
}

func StoreYaml(appname, info string) (resp common.Response) {
	//app's name should be uniq in this demo
	key := fmt.Sprintf("%v", appname)
	ChangeSheetCache.Store(key, info)
	return
}

func getYaml() (resp common.Response) {
	yamls := make([]string, 0)
	ApplicationCache.Range(func(key, value interface{}) bool {
		valueInfo := value.(string)
		yamls = append(yamls, valueInfo)
		return true
	})
	resp.Data = yamls
	return
}

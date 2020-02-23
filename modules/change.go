/*
*  Copyright 2020 Baidu Inc. All Rights Reserved.
*  @Author: caili02
*  @Date:   2020-02-23
*  @Description:
*/
package modules

import (
	"crypto/sha1"
	"encoding/hex"
	"yaml.v2"

	"github.com/astaxie/beego/logs"
	"github.com/oam-backend/common"
	"github.com/oam-backend/kube_api"
	"github.com/oam-backend/models"
	"github.com/oam-backend/modules/cache"
)

//! handle deploy tasks
//todo complete params
func HandleTask(sheet models.ChangeSheet) (resp common.Response) {
	resp = cache.StoreChangeSheet(sheet)
	if !resp.IsOk() {
		return
	}
	appDetail, ok := cache.ApplicationCache.Load(sheet.AppName)
	if !ok {
		return resp.ErrorDump(common.PARAM_ERROR, "can't find module")
	}
	appModel := appDetail.(models.ApplicationMetaInfo)
	//generate the ComponentSchematic first if this is fresh deplou
	if sheet.ChangeType == models.FreshDeployChangeType {
		component := kube_api.KubeAPIMeta{
			ApiVersion: common.KubeAPIVersion,
			Kind:       kube_api.ComponentSchematicKind,
			Spec: kube_api.ComponentSpec{
				WorkloadType: appModel.APPPublishedMode,
				Containers: []kube_api.ContainerSpec{
					kube_api.ContainerSpec{
						Name:  sheet.AppName,
						Image: sheet.Image,
					},
				},
			},
		}
		applicationConfigurationYaml, e := yaml.Marshal(component)
		if e != nil {
			return resp.ErrorDump(common.PARAM_ERROR, e.Error())
		} else {
			logs.Debug(string(applicationConfigurationYaml))
		}
		//todo store
	}

	//generate applicationConfiguration.yaml
	sheet.FakeID = Sha1String(sheet.AppName)
	appMeta := make(map[string]interface{})
	appMeta["name"] = sheet.AppName
	appMeta["uid"] = sheet.FakeID

	component := kube_api.ApplicationConfigurationComponent{
		Name:         sheet.AppName,
		InstanceName: sheet.AppName,
		Traits:       sheet.Traits,
	}
	appSpec := make(map[string]interface{})
	appSpec["Component"] = []kube_api.ApplicationConfigurationComponent{component}
	applicationConfiguration := kube_api.KubeAPIMeta{
		Kind:       kube_api.ApplicationConfigurationKind,
		ApiVersion: common.KubeAPIVersion,
		Metadata:   appMeta,
		Spec:       appSpec,
	}
	applicationConfigurationYaml, e := yaml.Marshal(applicationConfiguration)
	if e != nil {
		return resp.ErrorDump(common.PARAM_ERROR, e.Error())
	} else {
		logs.Debug(string(applicationConfigurationYaml))
		cache.StoreYaml(sheet.AppName, string(applicationConfigurationYaml))
		resp.Data = string(applicationConfigurationYaml)
	}

	return resp
}

func Sha1String(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}

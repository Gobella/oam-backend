/*
*  Copyright 2020 Baidu Inc. All Rights Reserved.
*  @Author: caili02
*  @Date:   2020-02-22
*  @Description:
*/
package kube_api

import (
	"testing"
	"yaml.v2"
)

func TestApplicationConfiguration_Valid(t *testing.T) {
	acObj := KubeAPIMeta{
		Kind:       ApplicationConfigurationKind,
		ApiVersion: "core.oam.dev/v1alpha1",
		Metadata:   make(map[string]interface{}),
		Spec:       make(map[string]interface{}),
	}
	acObj.Metadata["name"] = "test"
	paramsForTraits := make(map[string]interface{})
	paramsForTraits["hostName"] = "example.com"
	trait := ApplicationConfigurationTrait{
		Name:       "traits_test",
		Properties: paramsForTraits,
	}
	component := ApplicationConfigurationComponent{
		Name:         "hell0",
		InstanceName: "onajidesiyou",
		Traits:       []ApplicationConfigurationTrait{trait},
	}
	acObj.Spec["Component"] = []ApplicationConfigurationComponent{component}
	b, e := yaml.Marshal(acObj)
	if e != nil {
		t.Error(e.Error())
	} else {
		t.Log(string(b))
	}
}

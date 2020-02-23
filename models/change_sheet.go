package models

import "github.com/oam-backend/kube_api"

const (
	FreshDeployChangeType  = "freshDeploy"
	UpdateDeployChangeType = "updateDeploy"
	scaleDeployChangeType  = "scaleDeploy"
	deleteDeployChangeType = "deleteDeploy"
)

//store the deploy infomation
type ChangeSheet struct {
	FakeID     string                                   `json:"id"`
	AppName    string                                   `json:"app_name"`
	Image      string                                   `json:"image"`
	ChangeType string                                   `json:"change_type"`
	App        ApplicationMetaInfo                      `json:"app"`
	Traits     []kube_api.ApplicationConfigurationTrait `json:"traits"`
}

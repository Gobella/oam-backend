package models

import "time"

// the app's metadata demo
type ApplicationMetaInfo struct {
	Name             string    `json:"name"`
	ResourceVersion  string    `json:"resourceVersion"`
	NameSpace        string    `json:"namespace"`
	Port             string    `json:"port"`
	Generation       string    `json:"generation, omitempty"`
	CreatedTime      time.Time `json:"creationTimestamp, omitempty"`
	UpdatedTime      time.Time `json:"UpdatedTimestamp, omitempty"`
	APPPublishedMode string    `json:"appPublishedMode"`
}

func (am *ApplicationMetaInfo) Valid() bool {
	return true
}
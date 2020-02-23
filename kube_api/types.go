package kube_api

import "fmt"

const (
	ApplicationConfigurationKind = "ApplicationConfiguration"
	TraitKind                    = "Trait"
	ComponentSchematicKind       = "ComponentSchematic"
)

// ApplicationConfiguration.yaml
type KubeAPIMeta struct {
	ApiVersion string      `yaml:"apiVersion"`
	Kind       string      `yaml:"kind"`
	Metadata   interface{} `yaml:"metadata,omitempty"`
	Spec       interface{} `yaml:"spec,omitempty"`
}

func (ac KubeAPIMeta) Valid() (bool, error) {
	switch ac.Kind {
	case ApplicationConfigurationKind:
		return true, nil
	case TraitKind:
		return true, nil
	case ComponentSchematicKind:
		return true, nil
	default:
		return false, fmt.Errorf("the value(%v) is invalid in kind field", ac.Kind)
	}
}

type ApplicationConfigurationComponent struct {
	Name              string                              `yaml:"name"`
	InstanceName      string                              `yaml:"instanceName"`
	ParameterValues   []ApplicationConfiguratioParamValue `yaml:"parameterValues,omitempty"`
	Traits            []ApplicationConfigurationTrait     `yaml:"traits,omitempty"`
	ApplicationScopes []string                            `yaml:"applicationScopes,omitempty"`
}

type ApplicationConfigurationTrait struct {
	Name       string      `yaml:"name"`
	Properties interface{} `yaml:"properties,omitempty"`
}

type ApplicationConfiguratioParamValue struct {
	Name  string      `yaml:"name"`
	Value interface{} `yaml:"value"`
}

type TraitSpec struct {
	AppliesTo  []string    `yaml:"appliesTo" json:"appliesTo"`
	Properties interface{} `yaml:"properties" json:"properties"`
}

type ComponentSpec struct {
	WorkloadType string        `yaml:"workloadType" json:"workloadType"`
	Parameters   interface{} `yaml:"parameters" json:"parameters"`
	Containers   interface{} `yaml:"containers" json:"containers"`
}

type ContainerSpec struct {
	Name  string        `yaml:"name"`
	Image string        `yaml:"image"`
	Env   []interface{} `yaml:"env"`
}

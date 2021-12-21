// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	Name       *string        `json:",omitempty"`
	Resources  *PoolResources `json:",omitempty"`
	ShipaHost  *string        `json:",omitempty"`
	ShipaToken *string        `json:",omitempty"`
}

// PoolResources is autogenerated from the json schema
type PoolResources struct {
	General *PoolGeneral `json:",omitempty"`
	Node    *PoolNode    `json:",omitempty"`
}

// PoolGeneral is autogenerated from the json schema
type PoolGeneral struct {
	Setup           *PoolSetup           `json:",omitempty"`
	Plan            *PoolPlan            `json:",omitempty"`
	Security        *PoolSecurity        `json:",omitempty"`
	Access          *PoolServiceAccess   `json:",omitempty"`
	Services        *PoolServiceAccess   `json:",omitempty"`
	Router          *string              `json:",omitempty"`
	Volumes         []string             `json:",omitempty"`
	AppQuota        *PoolAppQuota        `json:",omitempty"`
	ContainerPolicy *PoolContainerPolicy `json:",omitempty"`
	NetworkPolicy   *PoolNetworkPolicy   `json:",omitempty"`
}

// PoolSetup is autogenerated from the json schema
type PoolSetup struct {
	Default             *bool   `json:",omitempty"`
	Public              *bool   `json:",omitempty"`
	Provisioner         *string `json:",omitempty"`
	KubernetesNamespace *string `json:",omitempty"`
}

// PoolPlan is autogenerated from the json schema
type PoolPlan struct {
	Name *string `json:",omitempty"`
}

// PoolSecurity is autogenerated from the json schema
type PoolSecurity struct {
	DisableScan        *bool    `json:",omitempty"`
	ScanPlatformLayers *bool    `json:",omitempty"`
	IgnoreComponents   []string `json:",omitempty"`
	IgnoreCVES         []string `json:",omitempty"`
}

// PoolServiceAccess is autogenerated from the json schema
type PoolServiceAccess struct {
	Append    []string `json:",omitempty"`
	Blacklist []string `json:",omitempty"`
}

// PoolAppQuota is autogenerated from the json schema
type PoolAppQuota struct {
	Limit *string `json:",omitempty"`
}

// PoolContainerPolicy is autogenerated from the json schema
type PoolContainerPolicy struct {
	AllowedHosts []string `json:",omitempty"`
}

// PoolNetworkPolicy is autogenerated from the json schema
type PoolNetworkPolicy struct {
	Ingress            *NetworkPolicyConfig `json:",omitempty"`
	Egress             *NetworkPolicyConfig `json:",omitempty"`
	DisableAppPolicies *bool                `json:",omitempty"`
}

// NetworkPolicyConfig is autogenerated from the json schema
type NetworkPolicyConfig struct {
	PolicyMode        *string             `json:",omitempty"`
	CustomRules       []NetworkPolicyRule `json:",omitempty"`
	ShipaRules        []NetworkPolicyRule `json:",omitempty"`
	ShipaRulesEnabled []string            `json:",omitempty"`
}

// NetworkPolicyRule is autogenerated from the json schema
type NetworkPolicyRule struct {
	ID           *string       `json:",omitempty"`
	Enabled      *bool         `json:",omitempty"`
	Description  *string       `json:",omitempty"`
	Ports        []NetworkPort `json:",omitempty"`
	Peers        []NetworkPeer `json:",omitempty"`
	AllowedApps  []string      `json:",omitempty"`
	AllowedPools []string      `json:",omitempty"`
}

// NetworkPort is autogenerated from the json schema
type NetworkPort struct {
	Protocol *string `json:",omitempty"`
	Port     *int    `json:",omitempty"`
}

// NetworkPeer is autogenerated from the json schema
type NetworkPeer struct {
	PodSelector       *NetworkPeerSelector `json:",omitempty"`
	NamespaceSelector *NetworkPeerSelector `json:",omitempty"`
	IPBlock           []string             `json:",omitempty"`
}

// NetworkPeerSelector is autogenerated from the json schema
type NetworkPeerSelector struct {
	MatchLabels      []Pair               `json:",omitempty"`
	MatchExpressions []SelectorExpression `json:",omitempty"`
}

// Pair is autogenerated from the json schema
type Pair struct {
	Key   *string `json:",omitempty"`
	Value *string `json:",omitempty"`
}

// SelectorExpression is autogenerated from the json schema
type SelectorExpression struct {
	Key      *string  `json:",omitempty"`
	Operator *string  `json:",omitempty"`
	Values   []string `json:",omitempty"`
}

// PoolNode is autogenerated from the json schema
type PoolNode struct {
	Drivers   []string       `json:",omitempty"`
	AutoScale *PoolAutoScale `json:",omitempty"`
}

// PoolAutoScale is autogenerated from the json schema
type PoolAutoScale struct {
	MaxContainer *int     `json:",omitempty"`
	MaxMemory    *int     `json:",omitempty"`
	ScaleDown    *float64 `json:",omitempty"`
	Rebalance    *bool    `json:",omitempty"`
}
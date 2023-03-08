package trufaas

type (
	FunctionInformation struct {
		Name               string             `json:"functionName"`
		Namespace          string             `json:"functionNamespace"`
		PackageInformation PackageInformation `json:"packageInformation"`
		Spec               FunctionSpec       `json:"functionSpec"`
	}

	PackageInformation struct {
		Name      string      `json:"packageName"`
		Namespace string      `json:"packageNamespace"`
		Spec      PackageSpec `json:"packageSpec"`
	}

	FunctionSpec struct {
		Environment Environment `json:"environment"`
		Package     struct {
			Packageref struct {
				Namespace       string `json:"namespace"`
				Name            string `json:"name"`
				Resourceversion string `json:"resourceversion"`
			} `json:"packageRef"`
			FunctionName string `json:"functionName,omitempty"`
		} `json:"package"`
		Resources struct {
		} `json:"resources"` //:TODO TruFaaS check later pkg\apis\core\v1\types.go has more attributes
		InvokeStrategy struct { //:TODO TruFaaS check this too
			ExecutionStrategy struct {
				ExecutorType          string `json:"ExecutorType"`
				MinScale              int    `json:"MinScale"`
				MaxScale              int    `json:"MaxScale"`
				TargetCPUPercent      int    `json:"TargetCPUPercent"`
				SpecializationTimeout int    `json:"SpecializationTimeout"`
			} `json:"executionStrategy"`
			StrategyType string `json:"strategyType"`
		} `json:"invokeStrategy"`
		FunctionTimeout int `json:"functionTimeout"`
		Idletimeout     int `json:"idleTimeout"`
		Concurrency     int `json:"concurrency"`
		RequestsPerPod  int `json:"requestsPerPod"`
	}

	PackageSpec struct {
		Environment Environment `json:"environment"`
		Source      Archive     `json:"source"`
		Deployment  Archive     `json:"deployment"`
		Buildcmd    string      `json:"buildcmd,omitempty"`
	}

	Environment struct {
		Namespace string `json:"namespace"`
		Name      string `json:"name"`
	}
	Archive struct {
		Type    string `json:"type"`
		Literal string `json:"literal,omitempty"`
		Url     string `json:"url,omitempty"`
	}
)

//:TODO TruFaaS check the nested levels inside the function spec, and limit the nested level

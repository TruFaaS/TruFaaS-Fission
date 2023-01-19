package trufaas

type (
	FunctionInformation struct {
		Name      string       `json:"function_name"`
		Namespace string       `json:"function_namespace"`
		Spec      FunctionSpec `json:"function_spec"`
	}

	PackageInformation struct {
		Name      string      `json:"package_name"`
		Namespace string      `json:"package_namespace"`
		Spec      PackageSpec `json:"package_spec"`
	}
)

type FunctionSpec struct {
	Environment struct {
		Namespace string `json:"namespace"`
		Name      string `json:"name"`
	} `json:"environment"`
	Package struct {
		Packageref struct {
			Namespace       string `json:"namespace"`
			Name            string `json:"name"`
			Resourceversion string `json:"resourceversion"`
		} `json:"packageref"`
		FunctionName string `json:"functionName,omitempty"`
	} `json:"package"`
	Resources struct {
	} `json:"resources"` // check later pkg\apis\core\v1\types.go has more attributes
	InvokeStrategy struct { //check this too
		ExecutionStrategy struct {
			ExecutorType          string `json:"ExecutorType"`
			MinScale              int    `json:"MinScale"`
			MaxScale              int    `json:"MaxScale"`
			TargetCPUPercent      int    `json:"TargetCPUPercent"`
			SpecializationTimeout int    `json:"SpecializationTimeout"`
		} `json:"ExecutionStrategy"`
		StrategyType string `json:"StrategyType"`
	} `json:"InvokeStrategy"`
	FunctionTimeout int `json:"functionTimeout"`
	Idletimeout     int `json:"idletimeout"`
	Concurrency     int `json:"concurrency"`
	RequestsPerPod  int `json:"requestsPerPod"`
}

type PackageSpec struct {
	Environment struct {
		Namespace string `json:"namespace"`
		Name      string `json:"name"`
	} `json:"environment"`
	Source struct {
		Type    string `json:"type"`
		Literal string `json:"literal,omitempty"`
		Url     string `json:"url,omitempty"`
	} `json:"source"`
	Deployment struct {
		Checksum struct {
		} `json:"checksum"`
	} `json:"deployment"`
	Buildcmd string `json:"buildcmd,omitempty"`
}

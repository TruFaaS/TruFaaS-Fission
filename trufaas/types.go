package trufaas

type (
	FunctionMetaData struct {
		FunctionInformation FunctionInformation `json:"function_information"`
		PackageInformation  PackageInformation  `json:"package_information"`
	}

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

	FunctionSpec struct {
		Environment     Environment    `json:"environment"`
		Package         Package        `json:"package"`
		InvokeStrategy  InvokeStrategy `json:"invoke_strategy"`
		FunctionTimeout int            `json:"function_timeout"`
		IdleTimeout     int            `json:"idle_timeout"`
		Concurrency     int            `json:"concurrency"`
		RequestsPerPod  int            `json:"requests_per_pod"`
	}

	Package struct {
		PackageRef struct {
			Namespace       string `json:"namespace"`
			Name            string `json:"name"`
			ResourceVersion string `json:"resource_version"`
		} `json:"packageRef"`
		FunctionName string `json:"functionName,omitempty"`
	}

	InvokeStrategy struct { //:TODO TruFaaS check this too
		ExecutionStrategy struct {
			ExecutorType          string `json:"executor-type"`
			MinScale              int    `json:"min_scale"`
			MaxScale              int    `json:"max_scale"`
			TargetCPUPercent      int    `json:"target_cpu_percent"`
			SpecializationTimeout int    `json:"specialization_timeout"`
		} `json:"execution_strategy"`
		StrategyType string `json:"strategy_type"`
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
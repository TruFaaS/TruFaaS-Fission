package trufaas

import (
	"bytes"
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
)

// createPkgInformation - util method to create/populate a trufaas.PackageInformation struct using a fv1.Package struct
func createPkgInformation(pkg fv1.Package) PackageInformation {

	var pkgSpec = PackageSpec{
		Environment: Environment{
			Namespace: pkg.Spec.Environment.Namespace,
			Name:      pkg.Spec.Environment.Name,
		},
		Source: Archive{
			Type:    string(pkg.Spec.Source.Type),
			Literal: pkg.Spec.Source.Literal,
			URL:     pkg.Spec.Source.URL,
		},
		Deployment: Archive{
			Type:    string(pkg.Spec.Deployment.Type),
			Literal: pkg.Spec.Deployment.Literal,
			URL:     pkg.Spec.Deployment.URL,
		},
		Buildcmd: pkg.Spec.BuildCommand,
	}

	pkgInformation := PackageInformation{
		Name:      pkg.Name,
		Namespace: pkg.Namespace,
		Spec:      pkgSpec,
	}

	return pkgInformation
}

// createFunctionInformation - util method to create/populate a trufaas.FunctionInformation struct using a fv1.Function struct
func createFunctionInformation(fn fv1.Function) FunctionInformation {
	var fnSpec = FunctionSpec{
		Environment: Environment{
			Namespace: fn.Spec.Environment.Namespace,
			Name:      fn.Spec.Environment.Name,
		},
		PackageRef: PackageRef{
			Namespace:       fn.Spec.Package.PackageRef.Namespace,
			Name:            fn.Spec.Package.PackageRef.Name,
			ResourceVersion: fn.Spec.Package.PackageRef.ResourceVersion,
		},
		InvokeStrategy: InvokeStrategy{
			ExecutionStrategy: ExecutionStrategy{
				ExecutorType:          string(fn.Spec.InvokeStrategy.ExecutionStrategy.ExecutorType),
				MinScale:              fn.Spec.InvokeStrategy.ExecutionStrategy.MinScale,
				MaxScale:              fn.Spec.InvokeStrategy.ExecutionStrategy.MaxScale,
				TargetCPUPercent:      fn.Spec.InvokeStrategy.ExecutionStrategy.TargetCPUPercent,
				SpecializationTimeout: fn.Spec.InvokeStrategy.ExecutionStrategy.SpecializationTimeout,
			},
			StrategyType: string(fn.Spec.InvokeStrategy.StrategyType),
		},
		FunctionTimeout: fn.Spec.FunctionTimeout,
		IdleTimeout:     *fn.Spec.IdleTimeout,
		Concurrency:     fn.Spec.Concurrency,
		RequestsPerPod:  fn.Spec.RequestsPerPod,
	}

	functionInformation := FunctionInformation{
		Name:      fn.Name,
		Namespace: fn.Namespace,
		Spec:      fnSpec,
	}

	return functionInformation
}

// SendToExternalComp - util method to send fnMetaData struct to trufaas external component
func SendToExternalComp(fnMetaData FunctionMetaData, URL string, method string) ([]byte, error) {
	jsonBody, err := jsoniter.Marshal(fnMetaData) // convert struct to json
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	AddTrustProtocolHeadersToReq(req) // Add necessary trust protocol headers

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // Read the response body
	if err != nil {
		return nil, err
	}
	GetTrustProtocolHeadersFromResp(resp)                                          // Get necessary trust protocol headers
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated { // If not success
		return body, fmt.Errorf("[Error Response] %s", string(body))
	}
	return body, nil
}

package trufaas

import (
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
	"reflect"
)

var fnMetaData *FunctionMetaData

func GetInstanceAtCreate() *FunctionMetaData {
	if fnMetaData == nil {
		fnMetaData = &FunctionMetaData{}
	}
	return fnMetaData
}
func (fnMetaData *FunctionMetaData) SavePkgInfoAtCreate(pkg fv1.Package) {
	fnMetaData.PackageInformation = createPkgInformation(pkg)
}

func (fnMetaData *FunctionMetaData) SaveFnInfoAtCreate(fn fv1.Function) {
	fnMetaData.FunctionInformation = createFunctionInformation(fn)
}

func SendInfoToAPIAtCreate() {
	fmt.Println("==========================Info send to api===============")
	//fmt.Println(fnMetaData.Name)
	fmt.Println(fnMetaData.PackageInformation.Name)
	//TODO:TruFaaS send data to API and store
}

func createPkgInformation(pkg fv1.Package) (pkgInformation PackageInformation) {

	var srcArchive Archive
	var deploymentArchive Archive
	if reflect.ValueOf(pkg.Spec.Source).FieldByName("Type").IsValid() {
		srcArchive = Archive{
			Type:    string(pkg.Spec.Source.Type),
			Literal: pkg.Spec.Source.Literal,
			URL:     pkg.Spec.Source.URL,
		}
	} else {
		deploymentArchive = Archive{
			Type:    string(pkg.Spec.Deployment.Type),
			Literal: pkg.Spec.Deployment.Literal,
			URL:     pkg.Spec.Deployment.URL,
		}
	}
	var pkgSpec = PackageSpec{
		Environment: Environment{
			Namespace: pkg.Spec.Environment.Namespace,
			Name:      pkg.Spec.Environment.Name,
		},
		Source:     srcArchive,
		Deployment: deploymentArchive,
		Buildcmd:   pkg.Spec.BuildCommand,
	}

	pkgInformation = PackageInformation{
		Name:      pkg.Name,
		Namespace: pkg.Namespace,
		Spec:      pkgSpec,
	}

	return pkgInformation
}

func createFunctionInformation(fn fv1.Function) (functionInformation FunctionInformation) {
	var fnSpec = FunctionSpec{
		Environment: Environment{},
		Package:     Package{},
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

	functionInformation = FunctionInformation{
		Name:      fn.Name,
		Namespace: fn.Namespace,
		Spec:      fnSpec,
	}

	return functionInformation
}

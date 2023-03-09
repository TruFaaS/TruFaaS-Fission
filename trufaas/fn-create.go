package trufaas

import (
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
)

var fnMetaData *FunctionMetaData

func GetInstanceAtCreate() *FunctionMetaData {
	if fnMetaData == nil {
		fnMetaData = &FunctionMetaData{}
	}
	return fnMetaData
}
func (fnMetaData *FunctionMetaData) SavePkgInfoAtCreate(pkg fv1.Package) {
	//fnInfoAtCreate.PackageInformation.Name = pkg.Name
	//TODO:TruFaaS populate rest of the fields
}

func (pkg fv1.Package) createPkgInformation(pkgInformation *PackageInformation) {
	pkgInformation = &PackageInformation{
		Name:      pkg.Name,
		Namespace: pkg.Namespace,
		Spec:      pkg.PackageSpec,
	}
}

func (fnMetaData *FunctionMetaData) SaveFnInfoAtCreate(fn fv1.Function) {
	//fnInfoAtCreate.Name = fn.Name
	//TODO:TruFaaS populate rest of the fields
}

func SendInfoToAPIAtCreate() {
	fmt.Println("==========================Info send to api===============")
	//fmt.Println(fnMetaData.Name)
	fmt.Println(fnMetaData.PackageInformation.Name)
	//TODO:TruFaaS send data to API and store
}

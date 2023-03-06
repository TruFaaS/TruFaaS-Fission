package trufaas

import (
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
)

var fnInfoAtCreate *FunctionInformation

func GetInstanceAtCreate() *FunctionInformation {
	if fnInfoAtCreate == nil {
		fnInfoAtCreate = &FunctionInformation{}
	}
	return fnInfoAtCreate
}
func (fnInfoAtCreate *FunctionInformation) SavePkgInfoAtCreate(pkg fv1.Package) {
	fnInfoAtCreate.PackageInformation.Name = pkg.Name
}

func (fnInfoAtCreate *FunctionInformation) SaveFnInfoAtCreate(fn fv1.Function) {
	fnInfoAtCreate.Name = fn.Name
}

func SendInfoToAPIAtCreate() {
	fmt.Println("==========================Info send to api===============")
	fmt.Println(fnInfoAtCreate.Name)
	fmt.Println(fnInfoAtCreate.PackageInformation.Name)
}

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
	fnMetaData.PackageInformation = createPkgInformation(pkg)
}

func (fnMetaData *FunctionMetaData) SaveFnInfoAtCreate(fn fv1.Function) {
	fnMetaData.FunctionInformation = createFunctionInformation(fn)
}

func SendInfoToAPIAtCreate() {
	body, err := SendToAPI(*fnMetaData, CreateURL, "POST")
	if err != nil {
		panic(err)
	}
	fmt.Println("==========================Response from API===============")
	fmt.Println(string(body))

}

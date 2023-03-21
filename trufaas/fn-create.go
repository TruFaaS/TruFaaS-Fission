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

func SendInfoToAPIAtCreate() error {
	body, err := SendToAPI(*fnMetaData, CreateURL, "POST")
	if err != nil {
		return fmt.Errorf("[Error Response] %s", string(body))
	}
	fmt.Println("<<<<<<<<<<<<<<<Response from TruFaaS API>>>>>>>>>>>>>>>>>>")
	fmt.Println(string(body))
	return nil
}

package trufaas

import (
	"encoding/json"
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
	fmt.Println("==========================Info send to api===============")
	val, _ := json.Marshal(fnMetaData)
	fmt.Println(string(val))
	//TODO:TruFaaS send data to API and store
}

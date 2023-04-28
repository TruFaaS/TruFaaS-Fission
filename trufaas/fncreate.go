package trufaas

import (
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
	"github.com/pkg/errors"
)

var fnMetaData *FunctionMetaData

// GetInstanceAtCreate returns singleton FnMetaData instance to populate with data
func GetInstanceAtCreate() *FunctionMetaData {
	if fnMetaData == nil {
		fnMetaData = &FunctionMetaData{}
	}
	return fnMetaData
}

// SavePkgInfoAtCreate populate the function pkg information
func (fnMetaData *FunctionMetaData) SavePkgInfoAtCreate(pkg fv1.Package) {
	fnMetaData.PackageInformation = createPkgInformation(pkg)
}

// SaveFnInfoAtCreate populate function information
func (fnMetaData *FunctionMetaData) SaveFnInfoAtCreate(fn fv1.Function) {
	fnMetaData.FunctionInformation = createFunctionInformation(fn)
}

// SendInfoToAPIAtCreate sends the populated FnMetaData to TruFaaS external component
func SendInfoToAPIAtCreate() error {
	_, err := SendToExternalComp(*fnMetaData, CreateURL, "POST")
	if err != nil {
		return errors.New(TrustValuesCreationFailMsg)
	}
	fmt.Println(TrustValuesCreationSuccessMsg)
	return nil
}

package trufaas

import (
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
)

func VerifyTrust(fn fv1.Function, pkg fv1.Package) error {

	fnMetaDataAtInvoke := createFnMetaDataAtInvocation(fn, pkg)
	_, err := SendToAPI(fnMetaDataAtInvoke, VerifyURL, "POST")
	if err != nil {
		return err
	}
	fmt.Printf("Trust of function with name %s verified by TruFaaS", fnMetaDataAtInvoke.FunctionInformation.Name)
	return nil
}

func createFnMetaDataAtInvocation(fn fv1.Function, pkg fv1.Package) FunctionMetaData {
	fnMetaDataInvocation := FunctionMetaData{
		FunctionInformation: createFunctionInformation(fn),
		PackageInformation:  createPkgInformation(pkg),
	}

	return fnMetaDataInvocation
}

package trufaas

import (
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
)

// VerifyTrust - used to verify trust of a function via the external component
func VerifyTrust(fn fv1.Function, pkg fv1.Package) error {

	fnMetaDataAtInvoke := createFnMetaDataAtInvocation(fn, pkg)
	_, err := SendToAPI(fnMetaDataAtInvoke, VerifyURL, "POST")
	if err != nil {
		return err
	}
	fmt.Printf("Trust of function with name %s verified by TruFaaS", fnMetaDataAtInvoke.FunctionInformation.Name)
	return nil
}

// createFnMetaDataAtInvocation - populates a FnMetaData struct with fn and pkg info at invocation
func createFnMetaDataAtInvocation(fn fv1.Function, pkg fv1.Package) FunctionMetaData {
	fnMetaDataInvocation := FunctionMetaData{
		FunctionInformation: createFunctionInformation(fn),
		PackageInformation:  createPkgInformation(pkg),
	}

	return fnMetaDataInvocation
}

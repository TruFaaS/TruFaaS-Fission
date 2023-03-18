package trufaas

import (
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
)

func VerifyTrust(fn fv1.Function, pkg fv1.Package) {

	fnMetaDataAtInvoke := createFnMetaDataAtInvocation(fn, pkg)
	_, err := SendToAPI(fnMetaDataAtInvoke, VerifyURL, "POST") //ToDo: TruFaaS change API URL
	if err != nil {
		panic(err)
	}

}

func createFnMetaDataAtInvocation(fn fv1.Function, pkg fv1.Package) FunctionMetaData {
	fnMetaDataInvocation := FunctionMetaData{
		FunctionInformation: createFunctionInformation(fn),
		PackageInformation:  createPkgInformation(pkg),
	}

	return fnMetaDataInvocation
}

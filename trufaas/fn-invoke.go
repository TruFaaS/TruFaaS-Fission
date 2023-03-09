package trufaas

import (
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
)

func VerifyTrust(fn fv1.Function, pkg fv1.Package) {
	fmt.Println(fn.Name)
	fmt.Println(pkg.Name)

	var fnMedatadata = createFnMetaDataAtInvocation(fn, pkg)
	//TODO:use variable before building

}

func createFnMetaDataAtInvocation(fn fv1.Function, pkg fv1.Package) (fnMetadata FunctionMetaData) {
	fnMetadata = FunctionMetaData{
		FunctionInformation: createFunctionInformation(fn),
		PackageInformation:  createPkgInformation(pkg),
	}

	return *fnMetaData
}

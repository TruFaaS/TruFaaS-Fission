package trufaas

import (
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
)

func VerifyTrust(fn fv1.Function, pkg fv1.Package) {

	fnMetaDataAtInvoke := createFnMetaDataAtInvocation(fn, pkg)
	//body, err := SendToAPI(fnMetaDataAtInvoke, CreateURL, "POST")
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println("==========================Response from API At Invoke===============")
	fmt.Println(fnMetaDataAtInvoke)

}

func createFnMetaDataAtInvocation(fn fv1.Function, pkg fv1.Package) (fnMetadata FunctionMetaData) {
	fnMetadata = FunctionMetaData{
		FunctionInformation: createFunctionInformation(fn),
		PackageInformation:  createPkgInformation(pkg),
	}

	return *fnMetaData
}

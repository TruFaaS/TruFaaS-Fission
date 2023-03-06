package trufaas

import (
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
)

func VerifyTrust(fn fv1.Function, pkg fv1.Package) {
	fmt.Println(fn.Name)
	fmt.Println(pkg.Name)
	//TODO:TruFaaS populate our own struct, send to API and get trust verified, if failed handle error
}

package trufaas

import (
	"bytes"
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
	"github.com/json-iterator/go"
	"go.uber.org/zap"
	"net/http"
)

func PrintFunctionStruct(msg string, f fv1.Function) {
	jsonData, err := jsoniter.Marshal(f) // convert struct to json
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonString := string(jsonData)
	fmt.Println("========================TruFaaS-Fn-Print======================")
	fmt.Println(msg)
	fmt.Println(jsonString)

}

func PrintPkgStruct(msg string, pkg fv1.Package) {
	jsonData, err := jsoniter.Marshal(pkg) // convert struct to json
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonString := string(jsonData)
	fmt.Println("========================TruFaaS-Pkg-Print======================")
	fmt.Println(msg)
	fmt.Println(jsonString)
	fmt.Println()
}

func PrintLiteral(msg string, l []byte) {
	fmt.Println("========================TruFaaS-Literal-Print======================")
	fmt.Println(msg)
	fmt.Println(string(l[:]))
}

func LogFnStruct(logger *zap.Logger, msg string, f fv1.Function) {
	jsonData, err := jsoniter.Marshal(f) // convert struct to json
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonString := string(jsonData)
	logger.Info(fmt.Sprintf("================TruFaaS====================== %s %s", msg, jsonString))
}

func SendStringTOAPI(msg string, pkg fv1.Package) {
	pkgJson, err := jsoniter.Marshal(pkg) // convert struct to json
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonString := string(pkgJson)
	jsonData := []byte(fmt.Sprintf(`{"content":"%s"}`, msg+" "+jsonString))
	req, err := http.NewRequest("POST", "https://test-data-api.onrender.com/info", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data Sent to API")
}

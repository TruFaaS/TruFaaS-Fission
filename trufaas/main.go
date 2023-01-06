package trufaas

import (
	"encoding/json"
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
)

func FnCreateInformationExtraction(f fv1.Function) {
	jsonData, err := json.Marshal(f) // convert struct to json
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonString := string(jsonData)
	fmt.Println("Following are the function info extracted at function creation")
	fmt.Println(jsonString)
}

func FnInvocationInformationExtraction(f fv1.Function) {
	jsonData, err := json.Marshal(f) // convert struct to json
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonString := string(jsonData)
	fmt.Println("Following are the function info extracted at function invocation")
	fmt.Println(jsonString)
}

// ===========================================
//func SendString(dataString string) {
//	data := fmt.Sprintf("Data from CLI %s ", dataString)
//	sendStringTOAPI(data)
//
//}

//func sendStringTOAPI(data string) {
//	jsonData := []byte(fmt.Sprintf(`{"content":"%s"}`, data))
//	req, err := http.NewRequest("POST", "https://test-data-api.onrender.com/info", bytes.NewBuffer(jsonData))
//	if err != nil {
//		panic(err)
//	}
//	req.Header.Set("Content-Type", "application/json")
//
//	client := &http.Client{}
//	_, err = client.Do(req)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("Data Sent to API")
//}

//func SavePackageInfo(storageSvcUrl string, archiveDownloadUrl string, filename string, pkg *fv1.Package) {
//
//	data := fmt.Sprintf("Storage Service URL %s \n Archive download URL %s \n File Name %s \n Package Name %s \n ", storageSvcUrl, archiveDownloadUrl, filename, pkg.Name)
//	sendStringTOAPI(data)
//
//}

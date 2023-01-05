package trufaas

import (
	"bytes"
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
	"io"
	"io/ioutil"
	"net/http"
)

func SavePackageInfo(storageSvcUrl string, archiveDownloadUrl string, filename string, pkg *fv1.Package) {

	data := fmt.Sprintf("Storage Service URL %s \n Archive download URL %s \n File Name %s \n Package Name %s \n ", storageSvcUrl, archiveDownloadUrl, filename, pkg.Name)
	sendStringTOAPI(data)

}

func SaveInfoFromCLI(dataString string) {
	data := fmt.Sprintf("Data from CLI %s ", dataString)
	sendStringTOAPI(data)

}

func sendStringTOAPI(data string) {
	jsonData := []byte(fmt.Sprintf(`{"content":"%s"}`, data))
	req, err := http.NewRequest("POST", "https://test-data-api.onrender.com/info", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

package truFaas

import (
	"bytes"
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
	v1 "github.com/fission/fission/pkg/apis/core/v1"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type (
	Connector struct {
		storageSvcUrl string
	}
)

func GetFunctionInfo(function v1.Function) {
	file, err := os.Create("funcSpec.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := fmt.Sprintf("Function Info %s", function.ObjectMeta.Name)
	_, err = io.WriteString(file, s)
	if err != nil {
		panic(err)
	}

}

func SavePackageInfo(storageSvcUrl string, archiveDownloadUrl string, filename string, pkg *fv1.Package) {
	//file, err := os.Create("funcSpec.txt")
	//if err != nil {
	//	panic(err)
	//}

	//defer file.Close()

	s := fmt.Sprintf("Storage Service URL %s \n Archive download URL %s \n File Name %s \n Package Name %s \n ", storageSvcUrl, archiveDownloadUrl, filename, pkg.Name)
	//_, err = io.WriteString(file, s)
	//if err != nil {
	//	panic(err)
	//}
	jsonData := []byte(fmt.Sprintf(`{"info":"%s"}`, s))
	// jsonData := []byte(`{"info":"Hi"}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/", bytes.NewBuffer(jsonData))
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

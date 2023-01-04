package truFaas

import (
	"fmt"
	fv1 "github.com/fission/fission/pkg/apis/core/v1"
	v1 "github.com/fission/fission/pkg/apis/core/v1"
	"io"
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
	file, err := os.Create("funcSpec.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := fmt.Sprintf("Storage Service URL %s \n Archive download URL %s \n File Name %s \n Package Name %s \n ", storageSvcUrl, archiveDownloadUrl, filename, pkg.Name)
	_, err = io.WriteString(file, s)
	if err != nil {
		panic(err)
	}
}

package mesh

import (
	"embed"
	"fmt"
	"github.com/chen-keinan/mesh-kridik/pkg/utils"
	"io/ioutil"
	"strings"
)

const IstioFolder = "istio"

var (
	//go:embed istio
	res embed.FS
)

func LoadIstioSpecs() ([]utils.FilesInfo, error) {
	dir, _ := res.ReadDir(IstioFolder)
	specs := make([]utils.FilesInfo, 0)
	for _, r := range dir {
		if strings.Contains(r.Name(), "_test") {
			continue
		}
		file, err := res.Open(fmt.Sprintf("%s/%s", IstioFolder, r.Name()))
		if err != nil {
			return specs, err
		}
		data, err := ioutil.ReadAll(file)
		spec := utils.FilesInfo{Name: r.Name(), Data: string(data)}
		if err != nil {
			return specs, err
		}
		if err != nil {
			return specs, err
		}
		specs = append(specs, spec)
	}
	return specs, nil
}

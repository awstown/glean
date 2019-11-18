package cmd

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Insights struct {
	HostOS              string  `json:"HostOS"`
	HostOSVersion       string  `json:"HostOSVersion"`
	Cores               int     `json:"Cores"`
	LoadAvg15           float32 `json:"LoadAvg15"`
	DiskUsage           int     `json:"DiskUsage"`
	DockerVersion       string  `json:"DockerVersion"`
	DockerStorageDriver string  `json:"DockerStorageDriver"`
}

func parseJS(reader io.Reader) {

}

func processFile(srcFile io.Reader) {
	// Unzip file
	archive, err := gzip.NewReader(srcFile)
	if err != nil {
		log.Fatal("Failed right out of the gate")
	}

	tr := tar.NewReader(archive)

	for true {
		h, err := tr.Next()
		if err == io.EOF {
			fmt.Println("woot worked")
			break
		}
		if filepath.Ext(h.Name) == ".json" {
			fmt.Println("True")
		}
	}

	fmt.Println("woot worked")
}

func getContents(file io.Reader) {
	// Unzip file
	archive, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal("Failed right out of the gate")
	}
	tr := tar.NewReader(archive)

	for true {
		h, err := tr.Next()
		if err == io.EOF {
			fmt.Println("woot worked")
			break
		}
		fmt.Println(h.Name)
		bs, err := ioutil.ReadAll(tr)
		fmt.Println(string(bs))

	}
}

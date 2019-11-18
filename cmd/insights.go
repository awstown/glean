package cmd

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type Insight struct {
	HostOS              string `json:"OS"`
	HostOSVersion       string `json:"OperatingSystem"`
	Cores               int    `json:"NCPU"`
	LoadAvg15           string
	DiskUsage           int
	DockerVersion       string `json:"Version"`
	DockerStorageDriver string `json:"Driver"`
}

func processFile(srcFile io.Reader) {
	// Define output
	var insight Insight
	// Unzip file
	archive, err := gzip.NewReader(srcFile)
	if err != nil {
		log.Fatal("Failed right out of the gate")
	}

	tr := tar.NewReader(archive)

	for true {
		h, err := tr.Next()
		if err == io.EOF {
			//fmt.Println("woot worked")
			break
		}
		if h.Typeflag == tar.TypeDir {
			continue
		}

		switch fname := filepath.Base(h.Name); fname {
		case "docker_version.json", "docker_info.json":
			bs, _ := ioutil.ReadAll(tr)
			err = json.Unmarshal(bs, &insight)
		case "loadavg":
			//fmt.Println("Running....", h.Name)
			bs, _ := ioutil.ReadAll(tr)
			col := strings.Split(string(bs), " ")
			insight.LoadAvg15 = col[2]
		case "df":
			// TODO: Implement parsing out the df for usage
			continue
		}
	}
	//fmt.Printf("%+v\n", insight)
	PrettyPrint(insight)
}

func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

package main

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"path"

	"testing"
)

var deps = make(map[string]Dependency)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	tmpDir, _ := ioutil.TempDir(os.TempDir(), "srlt")
	os.Setenv("BASEPATH", tmpDir)
	wd, _ := os.Getwd()
	os.Setenv("FILE", path.Join(wd, "srlt_test.json"))
	conf.Env()
	initConf()

	// read the file
	file, _ := conf.String("file")
	if buffer, err := ioutil.ReadFile(file); err == nil {
		json.Unmarshal(buffer, &deps)
	} else {
		t.Error(err)
	}

	Ω(len(deps)).Should(BeNumerically(">", 0))
	RunSpecs(t, "Clone Suite")
}

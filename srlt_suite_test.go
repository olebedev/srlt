package main

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var file = func() string {
	wd, _ := os.Getwd()
	return path.Join(wd, "srlt_test.yaml")
}()

var base = func() string {
	tmpDir, _ := ioutil.TempDir(os.TempDir(), "srlt")
	return tmpDir
}()

var srlt = &Srlt{
	Deps: make(map[string]*Dependency),
	Base: base,
}

func TestSrlt(t *testing.T) {
	RegisterFailHandler(Fail)

	if buffer, err := ioutil.ReadFile(file); err == nil {
		yaml.Unmarshal(buffer, srlt)
	} else {
		t.Error(err)
	}

	for _, d := range srlt.Deps {
		d.base = base
	}

	Ω(len(srlt.Deps)).Should(BeNumerically(">", 0))
	RunSpecs(t, "Srlt Suite")
}

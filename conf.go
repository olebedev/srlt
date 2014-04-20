package main

import (
	"github.com/olebedev/config"
	stdlog "log"
	"os"
	"path/filepath"
)

var log = stdlog.New(os.Stdout, "\033[1;33m[srlt] >>\033[m ", 0)
var conf, _ = config.ParseYaml(`
basepath: pwd
force: false
file: srlt.json
`)

func initConf() {
	// basepath
	bp, _ := conf.String("basepath")
	if bp == "pwd" {
		bp, _ = os.Getwd()
	}
	bp, _ = filepath.Abs(bp)
	conf.Set("basepath", bp)

	//file
	f, _ := conf.String("file")
	f, _ = filepath.Abs(f)
	conf.Set("file", f)
}

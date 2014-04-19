package main

import (
	"flag"
	"github.com/olebedev/config"
	"os"
	"path"
	"path/filepath"
)

var conf, _ = config.ParseYaml(`
gopath: $GOPATH
basepath: $GOPATH/src
force: false
file: sslc.json
`)

func init() {
	conf.Env().Flag()
	if flag.NArg() == 0 {
		usage()
		os.Exit(0)
	}
	// basepath
	bp, _ := conf.String("basepath")
	if bp == "$GOPATH/src" {
		gopath, _ := conf.String("gopath")
		bp = path.Join(gopath, "src")
	}
	bp, _ = filepath.Abs(bp)
	conf.Set("basepath", bp)

	//file
	f, _ := conf.String("file")
	f, _ = filepath.Abs(f)
	conf.Set("file", f)
}

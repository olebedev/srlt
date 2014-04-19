package main

import (
	"fmt"
)

func usage() {
	fmt.Println(`
Usage of sslc:
  sslc [options] snapshot  : save your current state
  sslc [options] restore   : restore state from file

Options:
  --basepath="$GOPATH/src" : path to scan and restore it
  --file="sslc.json"       : filename for read and write it
  --force="false"          : rewrite each dependency in file if exists
`)
}

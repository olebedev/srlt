package main

import (
	"fmt"
)

func usage() {
	fmt.Println(`
Usage of srlt:
  srlt [options] snapshot  : save your current state
  srlt [options] restore   : restore state from file

Options:
  --basepath="$GOPATH/src" : path to scan and restore it
  --file="srlt.json"       : filename for read and write it
`)
}

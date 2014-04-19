package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	conf.Env().Flag()
	if flag.NArg() == 0 {
		usage()
		os.Exit(0)
	}
	initConf()
	var err error
	switch flag.Arg(flag.NArg() - 1) {
	case "snapshot", "s":
		_, err = snapshot()
	case "restore", "r":
		err = restore()
	default:
		usage()
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

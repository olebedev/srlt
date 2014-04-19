package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var err error
	switch flag.Arg(flag.NArg() - 1) {
	case "snapshot", "s":
		err = save()
	case "restore", "r":
		err = install()
	default:
		usage()
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"

	"github.com/codegangsta/cli"
)

// save all dependency to file
func snapshotAction(c *cli.Context) {
	file := c.GlobalString("file")
	file, err := filepath.Abs(file)
	must(err)
	bp := c.GlobalString("path")
	verbose := c.GlobalBool("verbose")
	snapshot(file, bp, true, verbose)
}

func snapshot(file, base string, force, verbose bool) (map[string]*Dependency, error) {

	srlt := &Srlt{
		Deps:    make(map[string]*Dependency),
		Base:    base,
		verbose: verbose,
	}

	// read the file
	if buffer, err := ioutil.ReadFile(file); err == nil {
		yaml.Unmarshal(buffer, &srlt)
	} else if srlt.verbose {
		log("creating file %s ...\n", file)
	}

	bp, err := filepath.Abs(srlt.Base)
	if err != nil {
		return srlt.Deps, err
	}
	bp += string(os.PathSeparator)

	if srlt.verbose {
		log("lookup at %s...\n", bp)
	}

	err = filepath.Walk(bp, func(p string, f os.FileInfo, err error) error {
		dep := &Dependency{
			path:    p,
			base:    bp,
			verbose: srlt.verbose,
		}

		if err := dep.parse(); err == nil {
			if _, exist := srlt.Deps[dep.Name]; !exist && force {
				srlt.Deps[dep.Name] = dep
			}
		}

		return nil
	})

	if err != nil {
		return srlt.Deps, err
	}

	buffer, _ := yaml.Marshal(srlt)

	// save to the file
	err = ioutil.WriteFile(file, buffer, 0644)
	if err != nil {
		return srlt.Deps, err
	}

	return srlt.Deps, nil
}

func restoreAction(c *cli.Context) {
	file := c.GlobalString("file")
	file, err := filepath.Abs(file)
	must(err)

	verbose := c.GlobalBool("verbose")

	restore(file, verbose)
}

func restore(file string, verbose bool) error {
	srlt := &Srlt{verbose: verbose}
	// read the file
	if buffer, err := ioutil.ReadFile(file); err == nil {
		err = yaml.Unmarshal(buffer, srlt)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("file %s not found\n", file)
	}

	var wg sync.WaitGroup
	for _, dep := range srlt.Deps {
		dep.verbose = srlt.verbose
		dep.base = srlt.Base
		wg.Add(1)
		go func(dep *Dependency) {
			defer wg.Done()
			if err := dep.Install(); err != nil {
				panic(fmt.Sprintf("error: %v\ndep: #-v", err, dep)
			}
		}(dep)
	}
	wg.Wait()
	return nil
}

func execAction(c *cli.Context) {
	file := c.GlobalString("file")
	file, err := filepath.Abs(file)
	must(err)

	verbose := c.GlobalBool("verbose")
	tmpl := strings.Join(c.Args(), " ")
	execCommand(file, tmpl, verbose)
}

func execCommand(file, template string, verbose bool) {
	srlt := &Srlt{verbose: verbose}

	// read the file
	if buffer, err := ioutil.ReadFile(file); err == nil {
		yaml.Unmarshal(buffer, srlt)
		// } else {
		// 	return fmt.Errorf("file %s not found\n", file)
	}

	bp, err := filepath.Abs(srlt.Base)
	must(err)
	bp += string(os.PathSeparator)

	var wg sync.WaitGroup
	for _, dep := range srlt.Deps {
		dep.base = bp
		dep.verbose = srlt.verbose
		wg.Add(1)
		go func(dep *Dependency) {
			defer wg.Done()
			if err := dep.Exec(template); err != nil {
				panic(err)
			}
		}(dep)
	}
	wg.Wait()
}

// utils
func log(format string, i ...interface{}) {
	fmt.Printf(format, i...)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

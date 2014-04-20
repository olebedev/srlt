package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// Save all dependency to file
func snapshot(args ...string) (map[string]Dependency, error) {
	// search repositories
	force, _ := conf.Bool("force")
	basepath, _ := conf.String("basepath")
	basepath += string(os.PathSeparator)
	log.Printf("Lookup at '%s'...\n", basepath)

	deps := make(map[string]Dependency)

	// read the file
	file, _ := conf.String("file")
	if buffer, err := ioutil.ReadFile(file); err == nil {
		json.Unmarshal(buffer, &deps)
	} else {
		log.Printf("Creating file '%s'...\n", file)
	}

	err := filepath.Walk(basepath, func(p string, f os.FileInfo, err error) error {
		if dep, err := NewDependency(p); err == nil {
			if force {
				deps[dep.Name] = dep
			} else {
				if _, exist := deps[dep.Name]; !exist {
					deps[dep.Name] = dep
				}
			}
		}
		return nil
	})

	if err != nil {
		return deps, err
	}

	buffer, _ := json.MarshalIndent(deps, "", "  ")

	// save to the file
	err = ioutil.WriteFile(file, buffer, 0644)
	return deps, err
}

func restore(args ...string) error {
	// search repositories
	deps := make(map[string]Dependency)
	file, _ := conf.String("file")

	// read the file
	if buffer, err := ioutil.ReadFile(file); err == nil {
		json.Unmarshal(buffer, &deps)
	} else {
		return fmt.Errorf("file '%s' not found\n", file)
	}

	var wg sync.WaitGroup
	var err error
	for _, dep := range deps {
		wg.Add(1)
		go func(dep Dependency) {
			defer wg.Done()
			if _err := dep.Install(); err != nil {
				err = _err
			}
		}(dep)
	}
	wg.Wait()
	return err
}

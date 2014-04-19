package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

type Dependency struct {
	path   string `json:"-"`              // raw string `json:-,omitempty` with gopath
	Type   string `json:"type,omitempty"` // git, hg, bzr, svn
	Name   string `json:"name,omitempty"` // package name
	Remote string `json:"remote,omitempty"`
	Commit string `json:"commit,omitempty"` // commit hash
}

func (d *Dependency) parse() error {
	matched, err := regexp.MatchString("\\.git$|\\.hg$|\\.bzr$|\\.svn$", d.path)
	if err != nil {
		return err
	}

	if !matched {
		return fmt.Errorf("path '%s' is not repository.", d.path)
	}

	// set values
	d.Type = d.path[strings.LastIndex(d.path, ".")+1:]
	basepath, _ := conf.String("basepath")
	d.Name = strings.TrimPrefix(d.path[:strings.LastIndex(d.path, ".")-1], basepath)[1:]

	// ignore it if local
	// if strings.Index(d.Name, string(os.pathSeparator)) == 0 {
	// 	return fmt.Errorf("'%s' is not external repository.", d.Name)
	// }

	_, err = d.GetRemote()
	if err != nil {
		fmt.Println("remote err", err)
		return err
	}

	_, err = d.GetCommit()
	if err != nil {
		fmt.Println("commit err", err)
		return err
	}
	return nil
}

func (d *Dependency) GetRemote() (string, error) {
	if len(d.Remote) > 0 {
		return d.Remote, nil
	}

	switch d.Type {
	case "git":
		return d.gitGetRemote()
	case "hg":
		return d.hgGetRemote()
	case "bzr":
		return d.bzrGetRemote()
	case "svn":
		return d.svnGetRemote()
	default:
		return d.Remote, fmt.Errorf("VCS of type '%s' not found.", d.Type)
	}
}

func (d *Dependency) GetCommit() (string, error) {
	if len(d.Commit) > 0 {
		return d.Commit, nil
	}
	switch d.Type {
	case "git":
		return d.gitGetCommit()
	case "hg":
		return d.hgGetCommit()
	case "bzr":
		return d.bzrGetCommit()
	case "svn":
		return d.svnGetCommit()
	default:
		return d.Remote, fmt.Errorf("VCS of type '%s' not found.", d.Type)
	}
}

func NewDependency(s string) (Dependency, error) {
	dep := Dependency{path: s}
	if err := dep.parse(); err != nil {
		return Dependency{}, fmt.Errorf("path '%s' is not valid.")
	}
	return dep, nil
}

func (d *Dependency) Validate() error {
	if len(d.Name) <= 0 {
		return fmt.Errorf("Name '%s' is not valid", d.Name)
	}

	if len(d.Type) < 2 {
		return fmt.Errorf("Type '%s' is not valid", d.Type)
	}

	if len(d.Remote) <= 0 {
		return fmt.Errorf("Remote '%s' is not valid", d.Remote)
	}

	if len(d.Commit) <= 0 {
		return fmt.Errorf("Commit '%s' is not valid", d.Commit)
	}
	return nil
}

func (d *Dependency) Exists() bool {
	basepath, _ := conf.String("basepath")
	filename := path.Join(basepath, d.Name, "."+d.Type)
	_, err := os.Stat(filename)
	return err == nil
}

func (d *Dependency) Pull() error {
	fmt.Printf("pull %s\n", d.Name)
	switch d.Type {
	case "git":
		return d.gitPull()
	case "hg":
		return d.hgPull()
	case "bzr":
		return d.bzrPull()
	case "svn":
		return d.svnPull()
	default:
		return fmt.Errorf("VCS of type '%s' not found.", d.Type)
	}
}

func (d *Dependency) Clone() error {
	fmt.Printf("clone %s\n", d.Name)

	// mkdir all
	basepath, _ := conf.String("basepath")
	err := os.MkdirAll(filepath.Dir(path.Join(basepath, d.Name)), 0777)
	if err != nil {
		fmt.Println("clone error:", err)
		return err
	}

	switch d.Type {
	case "git":
		return d.gitClone()
	case "hg":
		return d.hgClone()
	case "bzr":
		return d.bzrClone()
	case "svn":
		return d.svnClone()
	default:
		return fmt.Errorf("VCS of type '%s' not found.", d.Type)
	}
}

func (d *Dependency) Checkout() error {
	fmt.Printf("checkout %s\n", d.Name)
	switch d.Type {
	case "git":
		return d.gitCheckout()
	case "hg":
		return d.hgCheckout()
	case "bzr":
		return d.bzrCheckout()
	case "svn":
		return d.svnCheckout()
	default:
		return fmt.Errorf("VCS of type '%s' not found.", d.Type)
	}
}

func (d *Dependency) Install() error {
	if err := d.Validate(); err != nil {
		return err
	}

	var err error
	if d.Exists() {
		err = d.Pull()
	} else {
		err = d.Clone()
	}
	if err != nil {
		return err
	}

	return d.Checkout()
}

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

type Srlt struct {
	verbose bool

	Base string                 `yaml:"base"`
	Deps map[string]*Dependency `yaml:"dependencies"`
}

type Dependency struct {
	path     string
	base     string
	verbose  bool
	execTmpl string

	Type   string `yaml:"type,omitempty"`
	Name   string `yaml:"name,omitempty"`
	Remote string `yaml:"remote,omitempty"`
	Commit string `yaml:"commit,omitempty"`
}

func (d *Dependency) parse() error {
	matched, err := regexp.MatchString("\\.git$|\\.hg$|\\.bzr$|\\.svn$", d.path)
	if err != nil {
		return err
	}

	if !matched {
		return fmt.Errorf("path %s is not a repository", d.path)
	}

	d.Type = d.path[strings.LastIndex(d.path, ".")+1:]
	d.Name = strings.TrimPrefix(d.path[:strings.LastIndex(d.path, ".")-1], d.base)
	d.Name = strings.TrimPrefix(d.Name, string(os.PathSeparator))

	_, err = d.GetRemote()
	if err != nil {
		return err
	}

	_, err = d.GetCommit()
	if err != nil {
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
		return d.Remote, fmt.Errorf("VCS of type %s not found.", d.Type)
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
		return d.Remote, fmt.Errorf("VCS of type %s not found.", d.Type)
	}
}

func NewDependency(s, base string) (Dependency, error) {
	dep := Dependency{path: s, base: base}
	if err := dep.parse(); err != nil {
		return Dependency{}, err
	}
	return dep, nil
}

func (d *Dependency) Validate() error {
	if len(d.Name) <= 0 {
		return fmt.Errorf("Name %s is not valid", d.Name)
	}

	if len(d.Type) < 2 {
		return fmt.Errorf("Type %s is not valid", d.Type)
	}

	if len(d.Remote) <= 0 {
		return fmt.Errorf("Remote %s is not valid", d.Remote)
	}

	if len(d.Commit) <= 0 {
		return fmt.Errorf("Commit %s is not valid", d.Commit)
	}
	return nil
}

func (d *Dependency) Exists() bool {
	filename := path.Join(d.base, d.Name, "."+d.Type)
	_, err := os.Stat(filename)
	return err == nil
}

func (d *Dependency) Pull() error {
	if d.verbose {
		log("pull: %s\n", d.Name)
	}
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
		return fmt.Errorf("VCS of type %s not found.", d.Type)
	}
}

func (d *Dependency) Clone() error {
	if d.verbose {
		log("clone: %s\n", d.Name)
	}

	err := os.MkdirAll(filepath.Dir(path.Join(d.base, d.Name)), 0777)
	if err != nil {
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
		return fmt.Errorf("VCS of type %s not found.", d.Type)
	}
}

func (d *Dependency) Checkout() error {
	if d.verbose {
		log("checkout: %s@%s\n", d.Name, d.Commit)
	}
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
		return fmt.Errorf("VCS of type %s not found.", d.Type)
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

func (d *Dependency) Exec(command string) error {

	t := template.Must(template.New("command").Parse(command))
	var doc bytes.Buffer
	t.Execute(&doc, d)

	command = doc.String()

	if d.verbose {
		log("exec: %s\n", command)
	}

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Dir = d.base
	return cmd.Run()
}

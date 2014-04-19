package main

import (
	"os/exec"
	"path"
	"strings"
)

func (d *Dependency) bzrGetRemote() (string, error) {

	basepath, _ := conf.String("basepath")
	cmd := exec.Command("bzr", "info")
	cmd.Dir = path.Join(basepath, d.Name)
	out, err := cmd.Output()

	if err != nil {
		return d.Remote, err
	}

	strArr := strings.Split(string(out), "\n")
	str := strArr[len(strArr)-2] // last element is empty string
	d.Remote = strings.TrimPrefix(str, "  parent branch: ")

	return d.Remote, nil
}

func (d *Dependency) bzrGetCommit() (string, error) {
	basepath, _ := conf.String("basepath")
	cmd := exec.Command("bzr", "revno")
	cmd.Dir = path.Join(basepath, d.Name)
	out, err := cmd.Output()

	if err != nil {
		return d.Remote, err
	}

	d.Commit = strings.TrimSpace(string(out))

	return d.Commit, nil
}

func (d *Dependency) bzrClone() error {
	basepath, _ := conf.String("basepath")
	cmd := exec.Command("bzr", "branch", d.Remote, d.Name)
	cmd.Dir = basepath
	_, err := cmd.Output()
	return err
}

func (d *Dependency) bzrCheckout() error {
	basepath, _ := conf.String("basepath")
	cmd := exec.Command("bzr", "revert", "-r"+d.Commit)
	cmd.Dir = path.Join(basepath, d.Name)
	_, err := cmd.Output()
	return err
}

func (d *Dependency) bzrPull() error {
	basepath, _ := conf.String("basepath")
	cmd := exec.Command("bzr", "pull", d.Remote, d.Name)
	cmd.Dir = path.Join(basepath, d.Name)
	_, err := cmd.Output()
	return err
}

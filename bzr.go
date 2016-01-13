package main

import (
	"os/exec"
	"path"
	"strings"
)

func (d *Dependency) bzrGetRemote() (string, error) {

	cmd := exec.Command("bzr", "info")
	cmd.Dir = path.Join(d.base, d.Name)
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
	cmd := exec.Command("bzr", "revno")
	cmd.Dir = path.Join(d.base, d.Name)
	out, err := cmd.Output()

	if err != nil {
		return d.Remote, err
	}

	d.Commit = strings.TrimSpace(string(out))

	return d.Commit, nil
}

func (d *Dependency) bzrClone() error {
	cmd := exec.Command("bzr", "branch", d.Remote, d.Name)
	cmd.Dir = d.base
	_, err := cmd.Output()
	return err
}

func (d *Dependency) bzrCheckout() error {
	cmd := exec.Command("bzr", "revert", "-r"+d.Commit)
	cmd.Dir = path.Join(d.base, d.Name)
	_, err := cmd.Output()
	return err
}

func (d *Dependency) bzrPull() error {
	cmd := exec.Command("bzr", "pull")
	cmd.Dir = path.Join(d.base, d.Name)
	_, err := cmd.Output()
	return err
}

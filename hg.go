package main

import (
	"os/exec"
	"path"
	"strings"
)

func (d *Dependency) hgGetRemote() (string, error) {

	basepath, _ := conf.String("basepath")
	cmd := exec.Command("hg", "paths", "default")
	cmd.Dir = path.Join(basepath, d.Name)
	out, err := cmd.Output()

	if err != nil {
		return d.Remote, err
	}

	d.Remote = strings.TrimSpace(string(out))
	return d.Remote, nil
}

func (d *Dependency) hgGetCommit() (string, error) {
	basepath, _ := conf.String("basepath")
	cmd := exec.Command("hg", "id", "-i")
	cmd.Dir = path.Join(basepath, d.Name)
	out, err := cmd.Output()

	if err != nil {
		return d.Remote, err
	}

	d.Commit = strings.TrimSpace(string(out))

	return d.Commit, nil
}

func (d *Dependency) hgClone() error {
	basepath, _ := conf.String("basepath")
	cmd := exec.Command("hg", "clone", d.Remote, d.Name)
	cmd.Dir = basepath
	_, err := cmd.Output()
	return err
}

func (d *Dependency) hgPull() error {
	basepath, _ := conf.String("basepath")
	cmd := exec.Command("hg", "pull")
	cmd.Dir = path.Join(basepath, d.Name)
	_, err := cmd.Output()
	return err
}

func (d *Dependency) hgCheckout() error {
	basepath, _ := conf.String("basepath")
	cmd := exec.Command("hg", "revert", "-r", d.Commit, "--all")
	cmd.Dir = path.Join(basepath, d.Name)
	_, err := cmd.Output()
	return err
}

package main

import (
	"os/exec"
	"path"
	"strings"
)

func (d *Dependency) gitGetRemote() (string, error) {
	cmd := exec.Command("git", "remote", "-v")
	cmd.Dir = path.Join(d.base, d.Name)
	out, err := cmd.Output()

	if err != nil {
		return d.Remote, err
	}

	strArr := strings.Split(string(out), "\n")
	str := strArr[len(strArr)-2]
	d.Remote = strings.TrimSpace(str[7 : len(str)-7])

	return d.Remote, nil
}

func (d *Dependency) gitGetCommit() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = path.Join(d.base, d.Name)
	out, err := cmd.Output()

	if err != nil {
		return d.Remote, err
	}

	d.Commit = strings.TrimSpace(string(out))

	return d.Commit, nil
}

func (d *Dependency) gitClone() error {
	cmd := exec.Command("git", "clone", d.Remote, d.Name)
	cmd.Dir = d.base
	_, err := cmd.Output()
	return err
}

func (d *Dependency) gitPull() error {
	cmd := exec.Command("git", "pull", "origin", "master")
	cmd.Dir = path.Join(d.base, d.Name)
	_, err := cmd.Output()
	return err
}

func (d *Dependency) gitCheckout() error {
	cmd := exec.Command("git", "checkout", d.Commit)
	cmd.Dir = path.Join(d.base, d.Name)
	_, err := cmd.Output()
	return err
}

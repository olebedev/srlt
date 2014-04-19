package main

import (
	"fmt"
)

func (d *Dependency) svnGetRemote() (string, error) {
	return d.Remote, fmt.Errorf("GetRemote for '%s' is not implemented.", d.Type)
}

func (d *Dependency) svnGetCommit() (string, error) {
	return d.Commit, fmt.Errorf("GetCommit for '%s' is not implemented.", d.Type)
}

func (d *Dependency) svnClone() error {
	return fmt.Errorf("Clone for '%s' is not implemented.", d.Type)
}

func (d *Dependency) svnPull() error {
	return fmt.Errorf("Clone for '%s' is not implemented.", d.Type)
}

func (d *Dependency) svnCheckout() error {
	return fmt.Errorf("Clone for '%s' is not implemented.", d.Type)
}

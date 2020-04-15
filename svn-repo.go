// ///////////////////////////////////////////////////////////////////
// Filename: svn-repo.go
// Description:
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Tue Apr 14 21:49:46 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package vcs

import (
	"os/exec"
)

type SvnRepo struct {
	Remote string
	Local  string
}

func NewSvnRepo(remote, local string) (*SvnRepo, error) {
	return &SvnRepo{remote, local}, nil
}

func (r *SvnRepo) Clone() error {
	cmd := exec.Command("svn", "checkout", r.Remote, r.Local)
	err := CallPipedOutput(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (r *SvnRepo) Checkout(commit string) error {
	cmd := exec.Command("svn", "update", "-r", commit)
	cmd.Dir = r.Local
	err := CallPipedOutput(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (r *SvnRepo) Pull() error {
	cmd := exec.Command("svn", "update")
	cmd.Dir = r.Local
	err := CallPipedOutput(cmd)
	if err != nil {
		return err
	}
	return nil
}

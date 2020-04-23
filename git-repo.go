// ///////////////////////////////////////////////////////////////////
// Filename: git-repo.go
// Description:
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Thu Apr 23 23:33:32 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package vcs

import (
	"os/exec"
)

type GitRepo struct {
	Remote string
	Local  string
}

func NewGitRepo(remote, local string) (*GitRepo, error) {
	return &GitRepo{remote, local}, nil
}

func (r *GitRepo) Clone() error {
	cmd := exec.Command("git", "clone", r.Remote, r.Local)
	err := CallPipedOutput(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (r *GitRepo) Checkout(commit string) error {
	cmd := exec.Command("git", "checkout", commit)
	cmd.Dir = r.Local
	err := CallPipedOutput(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (r *GitRepo) Pull() error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = r.Local
	err := CallPipedOutput(cmd)
	if err != nil {
		return err
	}
	return nil
}

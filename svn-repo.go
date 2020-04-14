// ///////////////////////////////////////////////////////////////////
// Filename: svn-repo.go
// Description:
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Tue Apr 14 21:49:46 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package vcs

type SvnRepo struct {
	Remote string
	Local  string
}

func NewSvnRepo(remote, local string) (*SvnRepo, error) {
	return &SvnRepo{remote, local}, nil
}

func (r *SvnRepo) Type() RepoType {
	return Svn
}

func (r *SvnRepo) Clone() {
	CallPipedOutput("svn", "checkout", r.Remote, r.Local)
}

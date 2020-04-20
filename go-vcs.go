// ///////////////////////////////////////////////////////////////////
// Filename: main.go
// Description:
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Tue Apr 14 21:17:02 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package vcs

type RepoType byte

// VCS types
const (
	Svn = 0
	Git = 1
)

var (
	Verbatim = true
)

type Repository interface {
	Clone() error
	Checkout(commit string) error
	Pull() error
}

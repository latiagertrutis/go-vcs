// ///////////////////////////////////////////////////////////////////
// Filename: cmd.go
// Description:
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Tue Apr 14 21:20:52 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package vcs

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

func CallDuplicatedOutput(LogFile *os.File, Dir string, Param ...string) error {
	var StdOutBuff, StderrBuff bytes.Buffer
	var ErrStdOut, ErrStderr error
	var sy sync.WaitGroup

	cmd := exec.Command(Param[0], Param[1:]...)
	cmd.Dir = Dir

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	stdout := io.MultiWriter(os.Stdout, &StdOutBuff)
	stderr := io.MultiWriter(os.Stderr, &StderrBuff)

	err := cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: on cmd.Start(): %s\n", err)
		return err
	}

	sy.Add(1)
	go func() {
		_, ErrStdOut = io.Copy(stdout, stdoutIn)
		sy.Done()
	}()

	_, ErrStderr = io.Copy(stderr, stderrIn)
	sy.Wait()

	cmd.Wait()

	if ErrStdOut != nil || ErrStderr != nil {
		fmt.Fprint(os.Stderr, "Error: failed to capture stdout and stderr from commnad\n")
		if ErrStdOut != nil {
			return ErrStdOut
		}
		return ErrStderr
	}

	_, err = LogFile.Write([]byte("\nSTANDARD OUTPUT\n##################################################\n"))
	if err != nil {
		return err
	}
	_, err = LogFile.Write(StdOutBuff.Bytes())
	if err != nil {
		return err
	}
	_, err = LogFile.Write([]byte("\nSTANDARD ERROR\n##################################################\n"))
	if err != nil {
		return err
	}
	_, err = LogFile.Write(StderrBuff.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func CallPipedOutput(Dir string, Param ...string) error {
	var ErrStdOut, ErrStderr error
	var sy sync.WaitGroup

	cmd := exec.Command(Param[0], Param[1:]...)
	cmd.Dir = Dir

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	err := cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: on cmd.Start(): %s\n", err)
		return err
	}

	sy.Add(1)
	go func() {
		_, ErrStdOut = io.Copy(os.Stdout, stdoutIn)
		sy.Done()
	}()

	_, ErrStderr = io.Copy(os.Stderr, stderrIn)
	sy.Wait()

	cmd.Wait()

	if ErrStdOut != nil || ErrStderr != nil {
		fmt.Fprint(os.Stderr, "Error: failed to capture stdout and stderr from commnad\n")
		if ErrStdOut != nil {
			return ErrStdOut
		}
		return ErrStderr
	}
	return nil
}

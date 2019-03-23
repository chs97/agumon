package main

import (
	"errors"
	"bytes"
	"os/exec"
	"time"

	C "github.com/chs97/agumon/constant"
)

func compilerCmd(language string) *exec.Cmd {
	src := C.Path.SrcPath()
	var _exec *exec.Cmd
	switch language {
	case "CPP":
		dist := C.Path.Program()
		_exec = exec.Command("g++", "-x", "c++", "-O2", "-std=c++11", "-D__USE_MINGW_ANSI_STDIO=0", "-o", dist, src)
		break
	case "JAVA":
		_exec = exec.Command("javac", "-encoding", "UTF-8", src)
		break
	}
	return _exec
}

func compiler(language string) (int, error) {
	timeLimit := 10000

	cmd := compilerCmd(language)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	resErr := cmd.Start()
	resState := 0
	if resErr != nil {
		return C.CE, resErr
	}
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case err := <- done:
		if err != nil {
			resErr = errors.New(errb.String())
			resState = C.CE
			break
		}
	case <- time.After(time.Duration(timeLimit) * time.Millisecond):
		cmd.Process.Kill()
		resState = C.CTLE
	}
	return resState, resErr
}
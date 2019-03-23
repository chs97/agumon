package main

import (
	"fmt"
	"os/exec"
	"syscall"
	"time"
	"io"
	"os"

	C "github.com/chs97/agumon/constant"
	U "github.com/chs97/agumon/utils"
)
 
type result struct {
	time   	float64
	memory 	int64
	input 	string
	output 	string 
	state		int
	err 		error
}

const outputDir = "/tmp/out"



func execCmd() *exec.Cmd {
	language := U.GetEnv("LANGUAGE", "CPP")
	execFile := C.Path.Program()
	workDir := C.Path.Workspace()
	fmt.Println(execFile)
	var _exec *exec.Cmd
	switch language {
	case "CPP":
		_exec = exec.Command(execFile)
		break
	case "JAVA":
		_exec = exec.Command("java", "-Dfile.encoding=UTF-8", "-Xmx256M", "-Xss64M", "-classpath", workDir, "Main")
		break
	}
	fmt.Println(_exec.Args)
	return _exec
}

func work(inputPath, outputPath string, timeLimit int) (*result, error){
	// fmt.Println(outputPath, inputPath)
	// execCmd := C.Path.ExecCmd()
	res := &result{ state: 0, input: inputPath, output: outputPath, err: nil }
	in, err := os.Open(inputPath)
	if err != nil {
		return res, err
	}
	defer in.Close()

	out, err := os.Create(outputPath)
	if err != nil {
		return res, err
	}
	defer out.Close()

	cmd := execCmd()
	cmd.Stdout = out
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return res, err
	}
	io.Copy(stdin, in)
	stdin.Close()
	start := time.Now()
	err = cmd.Start()
	if err != nil {
		return res, err
	}
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case err := <- done:
		if err != nil {
			res.err = err
			res.state = C.RE
			break
		}
		res.memory = cmd.ProcessState.SysUsage().(*syscall.Rusage).Maxrss
	case <- time.After(time.Duration(timeLimit) * time.Millisecond):
		cmd.Process.Kill()
		res.state = C.TLE
	}
	elapsed := time.Since(start)
	res.time = elapsed.Seconds() * 1000
	return res, res.err
}
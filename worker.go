package main

import (
	"syscall"
	"time"
	"io"
	"os"

	C "github.com/chs97/agumon/constant"
)
 
type result struct {
	time   	float64
	memory 	int64
	input 	string
	output 	string 
	state	int
}

const outputDir = "/tmp/out"

func work(inputPath, outputPath string, timeLimit int) (*result, error){
	// fmt.Println(outputPath, inputPath)
	// execCmd := C.Path.ExecCmd()
	res := &result{ state: 0, input: inputPath, output: outputPath }
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

	cmd := C.Path.ExecCmd()
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
	return res, nil
}
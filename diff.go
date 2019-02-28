package main

import (
	"os/exec"
)

func diff(src, dist string) (bool, error) {
	cmd := exec.Command("diff", "-b", "-B", src, dist)
	err := cmd.Run()
	if err != nil {
		return false, err
	}
	return true, nil
}
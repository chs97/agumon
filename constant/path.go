package constant

import (
	U "github.com/chs97/agumon/utils"
	P "path"
	"os/exec"
)

// Path is used to get config file path
var Path *path

type path struct {
	workspace string
}

func init() {
	workspace := U.GetEnv("WORKSPACE", "/workdir")
	Path = &path{workspace: workspace}
}

func (p *path) Config() string {
	config := U.GetEnv("RUNNER_CONFIG", "config.yml")
	return P.Join(p.workspace, config)
}

func (p *path) Workspace() string {
	return p.workspace
}

func (p *path) Program() string {
	name := U.GetEnv("PROGRAM", "dist")
	return P.Join(p.workspace, name)
}

func (p *path) ExecCmd() *exec.Cmd {
	language := U.GetEnv("LANGUAGE", "CPP")
	path := U.GetEnv("PROGRAM", "dist")
	var _exec *exec.Cmd
	switch language {
	case "CPP":
		_exec = exec.Command(path)
	case "JAVA":
		_exec = exec.Command("java", "-Dfile.encoding=UTF-8", "-Xmx256M", "-Xss64M", "'Main'", path)
	}
	return _exec
}

func (p *path) AbsPath(filename string) string {
	return P.Join(p.workspace, filename)
}


func (p *path) GetFilePath(filename string) (string ,string) {
	input := p.AbsPath(filename)
	output := p.AbsPath(filename + ".out")
	return input, output
}
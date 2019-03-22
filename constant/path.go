package constant

import (
	U "github.com/chs97/agumon/utils"
	P "path"
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

func (p *path) ExecCmd() string {
	language := U.GetEnv("LANGUAGE", "CPP")
	path := U.GetEnv("PROGRAM", "dist")
	exec := ""
	switch language {
	case "CPP":
		exec = path
	case "JAVA":
		exec = "java -Dfile.encoding=UTF-8 -Xmx512M -Xss64M 'Main' " + path
	}
	return exec
}

func (p *path) AbsPath(filename string) string {
	return P.Join(p.workspace, filename)
}


func (p *path) GetFilePath(filename string) (string ,string) {
	input := p.AbsPath(filename)
	output := p.AbsPath(filename + ".out")
	return input, output
}
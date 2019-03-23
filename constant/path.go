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
	language := U.GetEnv("LANGUAGE", "CPP")
	file := ""
	switch language {
	case "CPP":
		file = P.Join(p.workspace, "dist")
		break
	case "JAVA":
		file = P.Join(p.workspace, "Main.java")
		break
	}
	return file
}

func (p *path) SrcPath() string {
	src := U.GetEnv("SOURCE", "src")
	return P.Join(p.workspace, src)
}

func (p *path) AbsPath(filename string) string {
	return P.Join(p.workspace, filename)
}


func (p *path) GetFilePath(filename string) (string ,string) {
	input := p.AbsPath(filename)
	output := p.AbsPath(filename + ".out")
	return input, output
}
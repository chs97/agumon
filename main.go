package main

import (
	"os"
	Config "github.com/chs97/agumon/config"
	Const "github.com/chs97/agumon/constant"
	Util "github.com/chs97/agumon/utils"

	yaml "gopkg.in/yaml.v2"
)

type result2Yaml struct {
	In		string 		`yaml:"input"`
	Memory 	int64 		`yaml:"memory"`
	Time 	float64 	`yaml:"time"`
	Result 	int 		`yaml:"result"`
}

func main() {
	c, _ := Config.Parser()
	all := []result2Yaml{}
	for _, data := range c.Data {
		input, output := Const.Path.GetFilePath(data.In)
		answer := Const.Path.AbsPath(data.Out)
		res, err := work(input, output, c.TimeLimit)
		if err != nil {
			res.state = Const.SE
		} else if res.state == Const.WAIT {
			diff, _ := diff(output, answer)
			if !diff {
				res.state = Const.WA
			} else {
				res.state = Const.AC
			}
		}
		all = append(all, result2Yaml{In: data.In, Memory: res.memory, Time: res.time, Result: res.state})
	}
	ans := Const.Path.AbsPath("answer.yml")
	yml, err := yaml.Marshal(all)
	if err != nil {
		os.Exit(2)
	}
	err = Util.WriteFile(ans, yml)
	if err != nil {
		os.Exit(3)
	}
	// fmt.Println(string(yml))
	// workspace := getEnv("WORKSPACE", "/home/workspace")
	// fmt.Println(workspace)
}
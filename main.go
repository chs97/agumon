package main

import (
	"fmt"
	"os"
	Config "github.com/chs97/agumon/config"
	Const "github.com/chs97/agumon/constant"
	Util "github.com/chs97/agumon/utils"

	yaml "gopkg.in/yaml.v2"
)

type result2Yaml struct {
	In			string 		`yaml:"input"`
	Memory 	int64 		`yaml:"memory"`
	Time 		float64 	`yaml:"time"`
	Result 	int 			`yaml:"result"`
	Error 	string 		`yaml:"error"`
}

type answer2Yaml struct {
	Results []result2Yaml `yaml:"results"`
	Error   string 				`yaml:"error"`
	State		int 					`yaml:"state"`
}

func main() {
	language := Util.GetEnv("LANGUAGE", "CPP")
	answer := &answer2Yaml{}

	fmt.Println("languaga", language)

	defer func () {
		ans := Const.Path.AbsPath("answer.yml")
		yml, err := yaml.Marshal(answer)
		if err != nil {
			os.Exit(2)
		}
		err = Util.WriteFile(ans, yml)
		if err != nil {
			os.Exit(3)
		}
	}()

	state, CE := compiler(language)
	if CE != nil {
		answer.Error = CE.Error()
		answer.State = state
		return
	} 
	c, _ := Config.Parser()
	fmt.Println(c)
	all := []result2Yaml{}
	Data := c.Data
	if (Const.Path.Mode() == "test") {
		Data = c.Test
	}
	for _, data := range Data {
		input, output := Const.Path.GetFilePath(data.In)
		answer := Const.Path.AbsPath(data.Out)
		res, err := work(input, output, c.TimeLimit)
		fmt.Println("work error: ", err)
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
		message := ""
		if err != nil {
			message = err.Error()
		}
		if (res.memory > int64(c.MemoryLimit)) {
			res.state = Const.MLE
		}
		all = append(all, result2Yaml{In: data.In, Memory: res.memory, Time: res.time, Result: res.state, Error: message})
	}
	answer.Results = all
	
	// fmt.Println(string(yml))
	// workspace := getEnv("WORKSPACE", "/home/workspace")
	// fmt.Println(workspace)
}
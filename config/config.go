package config

import (
	"strings"
	"fmt"
	"io/ioutil"
	"os"

	C "github.com/chs97/agumon/constant"
	yaml "gopkg.in/yaml.v2"
)

type data struct {
	In 	string 		`yaml:"in"`
	Out string 		`yaml:"out"`
}

// Config is used for runner
type Config struct {
	TimeLimit	int 		`yaml:"time-limit"`
	MemoryLimit int			`yaml:"memory-limit"`
	Data 		[]data 		`yaml:"data"`
}

func readConfig (path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("Configuration file %s is empty", path)
	}

	config := &Config{}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// Parser config and get inputs file
func Parser() (*Config, error) {
	path := C.Path.Config()
	config, err := readConfig(path)
	if (err != nil) {
		return nil, err
	}

	// inputs := config.Inputs
	// if len(config.Prefix) != 0 {
	// 	inputs = getInputs(config.Prefix)
	// }
	// fmt.Println(config)

	return config, nil
}

func getInputs(prefix string) []string {
	res := make([]string, 0)
	workdir := C.Path.Workspace()
	files, err := ioutil.ReadDir(workdir)
	
	if err != nil {
		return res
	}
	for _, file := range files {
		if file.IsDir() != true && file.Size() != 0 {
			name := file.Name()
			if strings.HasPrefix(name, prefix) {
				res = append(res, name)
			}
		}
	}
	return res
}
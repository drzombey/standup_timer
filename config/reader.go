package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Timer struct{
		SitTimeInMinutes int `yaml:"sitTimeInMinutes"`
		StandTimeInMinutes int `yaml:"standTimeInMinutes"`
	} `yaml:"timer"`
}

func processError(err error){
	fmt.Println(err.Error())
	os.Exit(2)
}

func Read(filename string) (config Config) {
	f, err := os.Open(filename)

	if err != nil {
		processError(err)
	}

	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		processError(err)
	}

	return cfg
}

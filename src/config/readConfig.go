package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	_ "io/ioutil"
)

type config struct {
	Host        string         `yaml:"host"`
	User        string         `yaml:"user"`
	Pwd         string         `yaml:"pwd"`
	Dbname      string         `yaml:"dbname"`
	Devs        []string       `yaml:"devs"`
	FilterTypes map[string]int `yaml:"filter_types"`
}

func ReadConfig() config {
	var c config
	// 首先初始化struct对象，用于保存配置文件信息
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(c.User)
	return c
}

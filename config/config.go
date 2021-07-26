package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	ErrConfigFileNotFound = errors.New("配置文件不存在")
)

type Config struct {
	Foo string
	Bar string
}

type ConfigLoader interface {
	Load(configFile string) (Config, error)
}

func GetConfigLoader(ext string) (ConfigLoader, error) {
	switch ext {
	case "json":
		return JsonConfigLoader{}, nil
	case "yaml":
		return YmlConfigLoader{}, nil
	default:
		return nil, errors.New("config loader not supported")
	}
}

func fileExists(file string) bool {
	info, err := os.Stat(file)
	return err == nil && !info.IsDir()
}

type JsonConfigLoader struct {
}

func (j JsonConfigLoader) Load(configFile string) (Config, error) {
	var c Config
	if !fileExists(configFile) {
		return c, ErrConfigFileNotFound
	}
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(data, &c)
	return c, err
}

type YmlConfigLoader struct {
}

func (j YmlConfigLoader) Load(configFile string) (Config, error) {
	var c Config
	if !fileExists(configFile) {
		return c, ErrConfigFileNotFound
	}
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal([]byte(data), &c)
	return c, err
}

func main() {
	loader, err := GetConfigLoader("json")
	if err != nil {
		fmt.Println(err)
		return
	}
	config, err := loader.Load("config.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", config)

	loader, err = GetConfigLoader("yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	config, err = loader.Load("config.yml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", config)
}

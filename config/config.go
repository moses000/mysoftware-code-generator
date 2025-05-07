package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
}

type AppConfig struct {
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	LogLevel    string `yaml:"log_level"`
}

type ServerConfig struct {
	Port    int `yaml:"port"`
	Timeout int `yaml:"timeout"`
}

type ApiConfig struct {
	BaseUrl string `yaml:"base_url"`
}

type CrudTemplateConfig struct {
	ModelTemplate    string `yaml:"model_template"`
	HandlerTemplate  string `yaml:"handler_template"`
	RoutesTemplate   string `yaml:"routes_template"`
	ServicesTemplate string `yaml:"services_template"`
	DtoTemplate      string `yaml:"dto_template"`
	OutputDir        string `yaml:"output_dir"`
}

type Config struct {
	Database     DatabaseConfig     `yaml:"database"`
	App          AppConfig          `yaml:"app"`
	Server       ServerConfig       `yaml:"server"`
	Api          ApiConfig          `yaml:"api"`
	CrudTemplate CrudTemplateConfig `yaml:"crud_template"`
}

var config Config

func LoadConfig(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not read config file: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("could not unmarshal config file: %v", err)
	}

	return nil
}

func GetConfig() *Config {
	return &config
}

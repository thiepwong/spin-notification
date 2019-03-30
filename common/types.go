package common

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Server struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}

type Db struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	DbName   string `yaml:"DbName"`
}
type Database struct {
	Pg    Db `yaml:"Pg"`
	Redis Db `yaml:"Redis"`
	Mqtt  Db `yaml:"Mqtt"`
}

type Config struct {
	Server Server   `yaml:"Server"`
	Db     Database `yaml:"Db"`
}

type VehicleHolder struct {
	Mobile       string `json:"mobile"`
	VehiclePlate string `json:"vehicle_plate"`
}

func (c *Config) GetConf() *Config {

	yamlFile, err := ioutil.ReadFile("config/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

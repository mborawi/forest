package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Database database
	Server   server
}
type database struct {
	Username string
	Password string
	Port     string
	DbName   string
}
type server struct {
	Port string
}

func ReadConfig(conf *Config) {
	pwd, _ := os.Getwd()
	p := path.Dir(pwd)

	filePath := fmt.Sprintf("%s/config.toml", p)
	if _, err := toml.DecodeFile(filePath, &conf); err != nil {
		log.Fatalln("could not parse config", err)
	}
}

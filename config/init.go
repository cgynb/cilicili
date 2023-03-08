package config

import (
	"github.com/BurntSushi/toml"
)

func Init() {
	if _, err := toml.DecodeFile("config.toml", &Conf); err != nil {
		panic(err)
	}
}

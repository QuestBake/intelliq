package config

import (
	toml "github.com/pelletier/go-toml"
)

//Conf is global reference to config data.
var Conf *toml.Tree

//LoadConfig reads config file
func LoadConfig(configFile string) (config *toml.Tree, err error) {
	return toml.LoadFile(configFile)
}

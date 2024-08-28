/*
Copyright © 2023 Stepan Zubkov stepanzubkov@florgon.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package services

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const (
    configDir = ".config"
    configFile = "ghclone.toml"
)

type Config struct {
    GithubAccessToken string
    DefaultUsername   string
}

// Parses .toml config file to Config struct and returns it. If any error occured - returns default config
func ParseConfig() *Config {
    cfg := &Config{}
    fullConfigDir :=  filepath.Join(os.ExpandEnv("$HOME"), configDir)
    toml.DecodeFile(filepath.Join(fullConfigDir, configFile), cfg)
    return cfg
}

// Writes Config struct to .toml config file
func (cfg *Config) WriteConfig() {
    tomlData, err := toml.Marshal(cfg)
    if err != nil {
        log.Fatalf("Error marshaling to TOML: %v", err)
    }

    fullConfigDir :=  filepath.Join(os.ExpandEnv("$HOME"), configDir)
    err = os.MkdirAll(fullConfigDir, 0755)
    if err != nil {
        log.Fatalf("Can't create config dir: %v", err)
    }
    err = os.WriteFile(filepath.Join(fullConfigDir, configFile), tomlData, 0644)
    if err != nil {
        log.Fatalf("Error writing to config file: %v", err)
    }
}



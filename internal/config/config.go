package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

const configPath = "./config.yml"

var once sync.Once

// instance единый экземпляр конфигурации сервиса.
var instance Config

// Config структура основной конфигурации сервиса.
type Config struct {
	Server    Address `yaml:"server"`
	Simulator Address `yaml:"simulator"`

	TempDir string `yaml:"tmp_dir"`
}

// Address конфигурация адреса (хост, порт).
type Address struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Get возвращает копию структуры Config.
func Get() Config {
	return instance
}

func init() {
	once.Do(func() {
		if err := cleanenv.ReadConfig(configPath, &instance); err != nil {
			panic(err)
		}
	})
}

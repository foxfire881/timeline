package config

import (
	"encoding/json"
	"os"
)

// var Setting Config

// type Config struct {
// 	kv any
// }

// func Load() {
// 	data, _ := os.ReadFile("config.json")
// 	_ = json.Unmarshal(data, &Setting.kv)
// }

// func (this Config) Get(keys ...string) any {
// 	v := this.kv
// 	for _, k := range keys {
// 		v = v.(map[string]any)[k]
// 	}
// 	return v
// }

type AppConfig struct {
	Server struct {
		Host string
		Port string
		Max  struct {
			Header int
			Upload int64
		}
	}
	Database struct {
		Timeline string
	}
}

var Config AppConfig

func Load(filepath string) {
	file, _ := os.Open(filepath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.Decode(&Config)
}

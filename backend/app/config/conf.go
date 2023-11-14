// Package          config
// @Title           conf.go
// @Description
// @Author          zhengzongwei<zhengzongwei@foxmail.com> 2023/11/14 17:02

package config

import (
	"backend/app/env"
	"bytes"
	_ "embed"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"io"
	"os"
	"path/filepath"
)

var conf = new(Config)

type Config struct {
	Language struct {
		Local string `toml:"local"`
	}
}

var (
	//go:embed conf/dev.toml
	devConfig []byte

	//go:embed conf/prod.toml
	prodConfig []byte
)

// IsExists 文件是否存在
func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	return f, err == nil || os.IsExist(err)
}

func init() {
	var r io.Reader
	switch env.Active().Value() {
	case "dev":
		r = bytes.NewBuffer(devConfig)
	case "prod":
		r = bytes.NewBuffer(prodConfig)
	default:
		r = bytes.NewBuffer(devConfig)
	}

	viper.SetConfigType("toml")
	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(conf); err != nil {
		panic(err)
	}

	viper.SetConfigName(env.Active().Value())
	viper.AddConfigPath("./config")

	configFile := "conf/" + env.Active().Value() + ".toml"
	_, ok := IsExists(configFile)
	if !ok {
		if err := os.MkdirAll(filepath.Dir(configFile), 0766); err != nil {
			panic(err)
		}
		f, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)

		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}

	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(conf); err != nil {
			panic(err)
		}

	})
}

func Get() Config {
	return *conf
}

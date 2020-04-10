package config

import (
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
)

func init() {
	// get current file path
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)

	var configPath string

	configPath = dir + "/../../../config/"
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	/*
		viper.SetConfigName("provider")
		viper.WatchConfig()
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
	*/
}

// Get get config value by key
func Get(key string) interface{} {
	return viper.Get(key)
}

// Set set config
func Set(key string, value interface{}) {
	viper.Set(key, value)
}

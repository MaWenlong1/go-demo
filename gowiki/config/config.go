package config

import (
	"github.com/jinzhu/configor"
)

//ConfigInfo 配置信息
type ConfigInfo struct {
	Port     string ""
	DataPath string ""
	TmplPath string ""
}

// NewConfig ..
func NewConfig(configFile string) *ConfigInfo {
	if configFile != "" {
		config := &ConfigInfo{}
		configor.Load(config, configFile)
		return config
	}
	return &ConfigInfo{Port: "8080", DataPath: "data/", TmplPath: "tmpl/"}
}

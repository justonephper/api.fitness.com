package config

import "time"

//框架设置
type GinConfig struct {
	AppName string `mapstructure:"app-name" json:"app-name" yaml:"app-name"`
	RunMode string `mapstructure:"run-mode" json:"run-mode" yaml:"run-mode"`

	HttpPort     int           `mapstructure:"http-port" json:"http-port" yaml:"http-port"`
	ReadTimeout  time.Duration `mapstructure:"read-timeout" json:"read-timeout" yaml:"read-timeout"`
	WriteTimeout time.Duration `mapstructure:"write-timeout" json:"write-timeout" yaml:"write-timeout"`
}

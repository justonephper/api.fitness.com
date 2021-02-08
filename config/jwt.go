package config

type JWT struct {
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
	Subject     string `mapstructure:"subject" json:"subject" yaml:"subject"`
	JwtSecret   string `mapstructure:"jwt-secret" json:"jwt-secret" yaml:"jwt-secret"`
	ExpiresTime int64  `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"`
	BufferTime  int64  `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`
}

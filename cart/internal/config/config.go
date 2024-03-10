package config

type Config struct {
	Port string `mapstructure:"PORT"`
	Host string `mapstructure:"HOST"`

	ProductServiceUrl   string `mapstructure:"PRODUCT_SERVICE_URL"`
	ProductServiceToken string `mapstructure:"PRODUCT_SERVICE_TOKEN"`
}

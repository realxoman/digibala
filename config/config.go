package config

type Properties struct {
	Port           string `env:"MY_APP_PORT" env-default:"8080"`
	Host           string `env:"HOST" env-default:"localhost"`
	DBHost         string `env:"DB_HOST" env-default:"localhost"`
	DBPort         string `env:"DB_PORT" env-default:"3306"`
	DBName         string `env:"DB_NAME" env-default:"digibala-db"`
	JwtTokenSecret string `env:"JWT_TOKEN_SECRET" env-default:"abrakadabra"`
}

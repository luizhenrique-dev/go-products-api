package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

func NewConfig() *conf {
	return cfg
}

type conf struct {
	dbDriver      string `mapstructure:"DB_DRIVER"`
	dbHost        string `mapstructure:"DB_HOST"`
	dbPort        string `mapstructure:"DB_PORT"`
	dbUser        string `mapstructure:"DB_USER"`
	dbPassword    string `mapstructure:"DB_PASSWORD"`
	dbName        string `mapstructure:"DB_NAME"`
	webServerPort string `mapstructure:"WEB_SERVER_PORT"`
	jwtScret      string `mapstructure:"JWT_SECRET"`
	jwtExpiresIn  int    `mapstructure:"JWT_EXPIRES_IN"`
	tokenAuth     *jwtauth.JWTAuth
}

func (c *conf) GetDBConnectionString() string {
	return "host=" + c.dbHost +
		" user=" + c.dbUser +
		" password=" + c.dbPassword +
		" dbname=" + c.dbName +
		" port=" + c.dbPort +
		" sslmode=disable TimeZone=America/Sao_Paulo"
}

func init() {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile("cmd/server/.env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	cfg = &conf{
		dbDriver:      viper.GetString("DB_DRIVER"),
		dbHost:        viper.GetString("DB_HOST"),
		dbPort:        viper.GetString("DB_PORT"),
		dbUser:        viper.GetString("DB_USER"),
		dbPassword:    viper.GetString("DB_PASSWORD"),
		dbName:        viper.GetString("DB_NAME"),
		webServerPort: viper.GetString("WEB_SERVER_PORT"),
		jwtScret:      viper.GetString("JWT_SECRET"),
		jwtExpiresIn:  viper.GetInt("JWT_EXPIRES_IN"),
	}

	cfg.tokenAuth = jwtauth.New("HS256", []byte(cfg.jwtScret), nil)
}

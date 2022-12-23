package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	BinPath       string `mapstructure:"BIN_PATH"`
	UploadPath    string `mapstructure:"UPLOAD_PATH"`
	TemplatesPath string `mapstructure:"TEMPLATES_PATH"`

	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort string `mapstructure:"REDIS_PORT"`

	DB            DB            `mapstructure:",squash"`
	Email         Email         `mapstructure:",squash"`
	Social        Social        `mapstructure:",squash"`
	ElasticSearch ElasticSearch `mapstructure:",squash"`
	Token         Token         `mapstructure:",squash"`
	Server        Server        `mapstructure:",squash"`
}

type Server struct {
	Port         string `mapstructure:"SERVER_PORT"`
	Host         string `mapstructure:"SERVER_HOST"`
	ReadTimeout  int    `mapstructure:"SERVER_READ_TIMEOUT"`
	WriteTimeout int    `mapstructure:"SERVER_WRITE_TIMEOUT"`
}

type Token struct {
	TokenSecret        string `mapstructure:"TOKEN_SECRET"`
	TokenAccessMinutes int    `mapstructure:"TOKEN_ACCESS_MINUTES"`
	TokenRefreshHours  int    `mapstructure:"TOKEN_REFRESH_HOURS"`
}

type DB struct {
	DB             string `mapstructure:"DB_DB"`
	User           string `mapstructure:"DB_USER"`
	Password       string `mapstructure:"DB_PASSWORD"`
	Host           string `mapstructure:"DB_HOST"`
	Port           string `mapstructure:"DB_PORT"`
	Name           string `mapstructure:"DB_NAME"`
	LogPath        string `mapstructure:"DB_LOG_PATH"`
	RemoteUser     string `mapstructure:"DB_REMOTE_USER"`
	RemotePassword string `mapstructure:"DB_REMOTE_PASSWORD"`
}

type Email struct {
	User          string `mapstructure:"EMAIL_USER"`
	Password      string `mapstructure:"EMAIL_PASSWORD"`
	From          string `mapstructure:"EMAIL_FROM"`
	Server        string `mapstructure:"EMAIL_SERVER"`
	Port          string `mapstructure:"EMAIL_PORT"`
	WriteTestFile bool   `mapstructure:"EMAIL_WRITE_TEST_FILE"`
}

type Social struct {
	InstagramToken string `mapstructure:"INSTAGRAM_TOKEN"`
	InstagramID    string `mapstructure:"INSTAGRAM_ID"`

	YouTubeApiKey    string `mapstructure:"YOUTUBE_API_KEY"`
	YouTubeChannelID string `mapstructure:"YOUTUBE_CHANNEL_ID"`
}

type ElasticSearch struct {
	ElasticSearchURL string `mapstructure:"ELASTIC_SEARCH_URL"`
	ElasticSearchOn  bool   `mapstructure:"ELASTIC_SEARCH_ON"`
}

func LoadConfig() (config *Config, err error) {
	viper.AddConfigPath(getEnvLocation())
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return config, err
}

func LoadTestConfig() (config *Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env.test")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return config, err
}

func getEnvLocation() string {
	envLocation := os.Getenv("ENV_LOCATION")
	if envLocation != "" {
		return envLocation
	}
	envLocation, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return envLocation
}

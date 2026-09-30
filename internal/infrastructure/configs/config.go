package configs

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	defaultHTTPPort               = "8000"
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
	defaultAccessTokenTTL         = 15 * time.Minute
	defaultRefreshTokenTTL        = 24 * time.Hour * 30
	defaultLimiterRPS             = 10
	defaultLimiterBurst           = 2
	defaultLimiterTTL             = 10 * time.Minute
	defaultVerificationCodeLength = 8

	EnvLocal = "local"
	Prod     = "prod"
)

type (
	Config struct {
		Environment string
		HTTP        HTTPConfig
		Auth        AuthConfig
		Catalog     CatalogConfig `mapstructure:"catalog"`
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}

	AuthConfig struct {
		JWT                    JWTConfig
		PasswordSalt           string
		VerificationCodeLength int `mapstructure:"verificationCodeLength"`
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
		SigningKey      string
	}

	CatalogListConfig struct {
		CatalogPath     string `mapstructure:"catalogPath"`
		CatalogItemCode string `mapstructure:"catalogItemCode"`
	}

	CatalogConfig struct {
		CatalogCode                 string            `mapstructure:"catalogCode"`
		ApiDrillCollar              CatalogListConfig `mapstructure:"apiDrillCollar"`
		ApiDrillPipe                CatalogListConfig `mapstructure:"apiDrillPipe"`
		AdjustableGaugeStablilizers CatalogListConfig `mapstructure:"adjustableGaugeStablilizers"`
		Additional                  CatalogListConfig `mapstructure:"additional"`
	}
)

// Init populates Config struct with values from config file
// located at filepath and environment variables.
func Init(configsDir string) (*Config, error) {
	if strings.TrimSpace(os.Getenv("APP_ENV")) == "" {
		return nil, fmt.Errorf("required environment variable APP_ENV is empty")
	}

	viper.AutomaticEnv()
	populateDefaults()

	if err := parseConfigFile(configsDir, os.Getenv("APP_ENV")); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)
	if port := os.Getenv("HTTP_PORT"); port != "" {
		n, err := strconv.Atoi(port)
		if err != nil || n < 1 || n > 65535 {
			return nil, fmt.Errorf("HTTP_PORT must be an integer from 1 to 65535")
		}
		cfg.HTTP.Port = port
	}
	if host := os.Getenv("HTTP_HOST"); host != "" {
		cfg.HTTP.Host = host
	}
	if cfg.HTTP.Host == "" {
		cfg.HTTP.Host = "127.0.0.1"
	}
	for _, key := range []string{"USER_ACCESS_TOKEN_SECRET", "USER_REFRESH_TOKEN_SECRET", "PASSWORD_SALT"} {
		if strings.TrimSpace(os.Getenv(key)) == "" {
			return nil, fmt.Errorf("required environment variable %s is empty", key)
		}
	}
	for _, key := range []string{"ACCESS_TOKEN_LIFETIME_MINUTES", "REFRESH_TOKEN_LIFETIME_MINUTES"} {
		value, err := strconv.Atoi(os.Getenv(key))
		if err != nil || value <= 0 {
			return nil, fmt.Errorf("%s must be a positive integer", key)
		}
	}

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("catalog", &cfg.Catalog); err != nil {
		return err
	}

	return viper.UnmarshalKey("http", &cfg.HTTP)
}

func setFromEnv(cfg *Config) {
	// TODO use envconfig https://github.com/kelseyhightower/envconfig
	cfg.Environment = os.Getenv("APP_ENV")
	cfg.Auth.JWT.SigningKey = os.Getenv("JWT_SIGNING_KEY")
}

func parseConfigFile(folder, env string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if env == EnvLocal {
		return nil
	}

	viper.SetConfigName(env)

	return viper.MergeInConfig()
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.max_header_megabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHTTPRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHTTPRWTimeout)
}

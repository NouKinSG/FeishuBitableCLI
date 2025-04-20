package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"

	"Base/internal/utils"
)

type Config struct {
	AppID           string `yaml:"app_id"`
	AppSecret       string `yaml:"app_secret"`
	AppToken        string `yaml:"app_token"`
	TableID         string `yaml:"table_id"`
	UserAccessToken string `yaml:"user_access_token"`
	Debug           bool   `yaml:"debug"`
}

var C *Config

func Load(env string) (*Config, error) {
	if env == "" {
		env = "local"
	}
	path := filepath.Join("configs", fmt.Sprintf("config.%s.yaml", env))
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置失败: %w", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}
	utils.SetDebugMode(cfg.Debug)
	C = &cfg
	return &cfg, nil
}

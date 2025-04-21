package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// LoadMap 从 configs 目录读取 env 环境下的 YAML 文件，
// 并反序列化到 map[string]interface{} 返回
func LoadMap(env string) (map[string]interface{}, error) {
	path := filepath.Join("configs", "config."+env+".yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := yaml.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// SaveMap 将 map 序列化成 YAML 写回 configs/config.<env>.yaml
func SaveMap(env string, m map[string]interface{}) error {
	path := filepath.Join("configs", "config."+env+".yaml")
	data, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// Reset 覆盖当前配置文件为示例文件。示例路径：configs/config.<env>.example.yaml
func Reset(env string) error {
	src := filepath.Join("configs", "config."+env+".example.yaml")
	dst := filepath.Join("configs", "config."+env+".yaml")
	body, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, body, 0644)
}

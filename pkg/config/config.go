package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

type Config struct {
	GCS GCS `json:"gcs"`
	//Client *storage.Client
}

type GCS struct {
	Bucket string `json:"bucket" yaml:"Bucket"`
	Creds  string `json:"creds" yaml:"Creds"`
}

func ReadOrCreate(configFile string) (*Config, error) {
	cfg := &Config{}
	var err error
	if len(configFile) == 0 {
		baseDir, err := os.UserHomeDir()
		if err != nil {
			return cfg, err
		}
		configDir := path.Join(baseDir, ".config", "photo-uploader")
		if os.MkdirAll(configDir, 0700) != nil {
			return cfg, err
		}
		configFile = path.Join(configDir, "config")
	}
	_, err = os.Stat(configFile)
	if os.IsNotExist(err) {
		CreateConfig(cfg)
		b, err := json.Marshal(cfg)
		if err != nil {
			panic(err)
		}
		os.WriteFile(configFile, b, 0600)
	}
	b, err := os.ReadFile(configFile)
	if err != nil {
		return cfg, err
	}
	if json.Unmarshal(b, cfg) != nil {
		return cfg, err
	}

	return cfg, err
}

func CreateConfig(cfg *Config) {
	cfg.GCS.Bucket = ReadValue("bucket")
	cfg.GCS.Creds = ReadFile("GCS credentials json file")
}

func ReadFile(name string) string {
	fmt.Printf("Enter path to %s: ", name)
	var reader = bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')

	b, err := os.ReadFile(strings.ReplaceAll(value, "\n", ""))
	if err != nil {
		panic(err)
	}
	return string(b)
}

func ReadValue(name string) string {
	fmt.Printf("Enter %s value: ", name)
	var reader = bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	return strings.ReplaceAll(value, "\n", "")
}

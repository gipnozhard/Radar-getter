package config

import (
	"encoding/json"
	"os"
)

type RadarConfig struct {
	ID      string `json:"id"`
	BaseURL string `json:"base_url"`
}

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

type MongoDBConfig struct {
	URI    string `json:"uri"`
	DBName string `json:"dbname"`
}

type Config struct {
	Radars   []RadarConfig  `json:"radars"`
	Postgres PostgresConfig `json:"postgres"`
	MongoDB  MongoDBConfig  `json:"mongodb"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

package config

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/ducho-metson/polysurance-test/utils"
)

type Config struct {
	DiscountsFilePath string
	OrdersFilePath    string
	ProductsFilePath  string

	DataPath          string `json:"dataPath"`
	DiscountsFilename string `json:"discountsFilename"`
	OrdersFilename    string `json:"ordersFilename"`
	ProductsFilename  string `json:"productsFilename"`
}

func LoadConfig(part string) (*Config, error) {
	var config Config

	data, err := utils.ReadFile("./config/config.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config data: %w", err)
	}

	config.DiscountsFilePath = filepath.Join(config.DataPath, part, config.DiscountsFilename)
	config.OrdersFilePath = filepath.Join(config.DataPath, part, config.OrdersFilename)
	config.ProductsFilePath = filepath.Join(config.DataPath, part, config.ProductsFilename)
	return &config, nil
}

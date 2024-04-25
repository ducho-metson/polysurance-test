package config

import (
	"encoding/json"
	"fmt"

	"github.com/ducho-metson/polysurance-test/utils"
)

type Config struct {
	DiscountsFilePath string `json:"discountsFilePath"`
	OrdersFilePath    string `json:"ordersFilePath"`
	ProductsFilePath  string `json:"productsFilePath"`
}

func LoadConfig() (*Config, error) {
	var config Config

	data, err := utils.ReadFile("./config/config.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config data: %w", err)
	}

	return &config, nil
}

package main

import (
	"fmt"

	"github.com/ducho-metson/polysurance-test/config"
	"github.com/ducho-metson/polysurance-test/model"
	"github.com/ducho-metson/polysurance-test/sales"
)

const (
	DiscountFilePath = "./data/discounts.json"
	OrdersFilePath   = "./data/discounts.json"
	ProductsFilePath = "./data/discounts.json"
)

func main() {
	fmt.Println("Starting Fullstack Devoloper Test...")

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("failed loading config file")
		return
	}

	data, err := model.ParseData(cfg.DiscountsFilePath, cfg.OrdersFilePath, cfg.ProductsFilePath)
	if err != nil {
		fmt.Println("failed parsing sales data file")
		return
	}

	sales.Calculate(data)
}

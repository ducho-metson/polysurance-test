package main

import (
	"fmt"

	"github.com/ducho-metson/polysurance-test/config"
	"github.com/ducho-metson/polysurance-test/model"
	"github.com/ducho-metson/polysurance-test/sales"
)

func main() {
	fmt.Println("Starting Fullstack Devoloper Test...")

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("failed loading config file")
		return
	}

	salesData, err := model.ParseSalesData(cfg.DiscountsFilePath, cfg.OrdersFilePath, cfg.ProductsFilePath)
	if err != nil {
		fmt.Println("failed parsing sales data file")
		return
	}

	sales.Calculate(salesData)
}

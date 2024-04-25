package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ducho-metson/polysurance-test/config"
	"github.com/ducho-metson/polysurance-test/model"
	"github.com/ducho-metson/polysurance-test/sales"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing argument: part1 or part2\nexample: go run main.go part1 ")
		return
	}

	part := os.Args[1]
	part = strings.Replace(part, " ", "", -1)
	if part != "part1" && part != "part2" {
		fmt.Println("Wrong argument: part1 or part2\nexample: go run main.go part1 ")
		return
	}

	fmt.Println("Starting Fullstack Devoloper Test...")

	cfg, err := config.LoadConfig(part)
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

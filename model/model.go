package model

import (
	"encoding/json"
	"fmt"

	"github.com/ducho-metson/polysurance-test/utils"
)

type Data struct {
	Products  []Product
	Orders    []Order
	Discounts []Discount
}

type Product struct {
	SKU   int     `json:"sku"`
	Price float64 `json:"price"`
}

type Item struct {
	SKU      int `json:"sku"`
	Quantity int `json:"quantity"`
}

type Order struct {
	OrderID  int    `json:"orderId"`
	Discount string `json:"discount"`
	Items    []Item `json:"items"`
}

type Discount struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

func ParseData(discountFilePath, ordersFilePath, productsFilePath string) (*Data, error) {
	discountsFileAsByte, err := utils.ReadFile(discountFilePath)
	if err != nil {
		fmt.Println("failed reading file", discountFilePath)
		return nil, err
	}

	discounts, err := parseDiscount(discountsFileAsByte)
	if err != nil {
		fmt.Println("failed parsing file", discountFilePath)
		return nil, err
	}

	ordersFileAsByte, err := utils.ReadFile(ordersFilePath)
	if err != nil {
		fmt.Println("failed reading file", ordersFilePath)
		return nil, err
	}

	orders, err := parseOrder(ordersFileAsByte)
	if err != nil {
		fmt.Println("failed parsing file", ordersFilePath)
		return nil, err
	}

	productsFileAsByte, err := utils.ReadFile(productsFilePath)
	if err != nil {
		fmt.Println("failed reading file", productsFilePath)
		return nil, err
	}

	products, err := parseProduct(productsFileAsByte)
	if err != nil {
		fmt.Println("failed parsing file", productsFilePath)
		return nil, err
	}

	return &Data{
		Products:  products,
		Orders:    orders,
		Discounts: discounts,
	}, nil
}

func parseProduct(jsonData []byte) ([]Product, error) {
	var products []Product
	err := json.Unmarshal([]byte(jsonData), &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func parseOrder(jsonData []byte) ([]Order, error) {
	var orders []Order
	err := json.Unmarshal([]byte(jsonData), &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func parseDiscount(jsonData []byte) ([]Discount, error) {
	var discounts []Discount
	err := json.Unmarshal([]byte(jsonData), &discounts)
	if err != nil {
		return nil, err
	}

	return discounts, nil
}

func GetPriceFromSku(products []Product, sku int) float64 {
	for _, product := range products {
		if product.SKU == sku {
			return product.Price
		}
	}

	return -1.0
}

func GetDiscountAsFloat(discount string) float64 {
	switch discount {
	case "SALE10":
		return 0.1
	case "SALE20":
		return 0.2
	case "SALE30":
		return 0.3
	default:
		return 0.0
	}
}

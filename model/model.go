package model

import (
	"encoding/json"

	"github.com/ducho-metson/polysurance-test/utils"
)

type Data struct {
	Products  map[int]float64
	Orders    []Order
	Discounts map[string]float64
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
		return nil, err
	}

	discounts, err := parseDiscount(discountsFileAsByte)
	if err != nil {
		return nil, err
	}

	ordersFileAsByte, err := utils.ReadFile(ordersFilePath)
	if err != nil {
		return nil, err
	}

	orders, err := parseOrder(ordersFileAsByte)
	if err != nil {
		return nil, err
	}

	productsFileAsByte, err := utils.ReadFile(productsFilePath)
	if err != nil {
		return nil, err
	}

	products, err := parseProduct(productsFileAsByte)
	if err != nil {
		return nil, err
	}

	return &Data{
		Products:  products,
		Orders:    orders,
		Discounts: discounts,
	}, nil
}

func parseProduct(jsonData []byte) (map[int]float64, error) {
	var products []Product
	err := json.Unmarshal([]byte(jsonData), &products)
	if err != nil {
		return nil, err
	}

	productsAsMap := make(map[int]float64)

	for _, product := range products {
		productsAsMap[product.SKU] = product.Price
	}

	return productsAsMap, nil
}

func parseOrder(jsonData []byte) ([]Order, error) {
	var orders []Order
	err := json.Unmarshal([]byte(jsonData), &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func parseDiscount(jsonData []byte) (map[string]float64, error) {
	var discounts []Discount
	err := json.Unmarshal([]byte(jsonData), &discounts)
	if err != nil {
		return nil, err
	}

	discountsAsMap := make(map[string]float64)

	for _, discount := range discounts {
		discountsAsMap[discount.Key] = discount.Value
	}

	return discountsAsMap, nil
}

func GetPriceFromSku(data *Data, sku int) float64 {
	if price, ok := data.Products[sku]; ok {
		return price
	}

	return -1.0
}

func GetDiscountAsFloat(data *Data, discountKey string) float64 {
	if discount, ok := data.Discounts[discountKey]; ok {
		return discount
	}

	return 0.0
}

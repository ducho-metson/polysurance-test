package model

import (
	"encoding/json"

	"github.com/ducho-metson/polysurance-test/utils"
)

// SalesData is a struct suposed to hold every information about Products, Discounts and Orders provided by
// given files. Both Products and Discounts fields are map in order to optimize searching operations.
type SalesData struct {
	Products  map[int]float64
	Orders    []Order
	Discounts map[string]Discount
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
	Key    string  `json:"key"`
	Value  float64 `json:"value"`
	Stacks string  `json:"stacks,omitempty"`

	isStackable bool
}

// ParseSalesData reads orders, products and discounts files from data folder and parse every information
// to a reference of type Data.
func ParseSalesData(discountFilePath, ordersFilePath, productsFilePath string) (*SalesData, error) {
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

	return &SalesData{
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

func parseDiscount(jsonData []byte) (map[string]Discount, error) {
	var discounts []Discount
	err := json.Unmarshal([]byte(jsonData), &discounts)
	if err != nil {
		return nil, err
	}

	discountsAsMap := make(map[string]Discount)

	for _, discount := range discounts {
		if discount.Stacks == "TRUE" {
			discount.isStackable = true
		}
		discountsAsMap[discount.Key] = discount
	}

	return discountsAsMap, nil
}

func (d *SalesData) GetPriceFromSku(sku int) float64 {
	if price, ok := d.Products[sku]; ok {
		return price
	}

	return -1.0
}

func (d *SalesData) GetDiscountAsFloat(discountKey string) float64 {
	if discount, ok := d.Discounts[discountKey]; ok {
		return discount.Value
	}

	return 0.0
}

func (d *SalesData) IsDiscountStackable(discountKey string) bool {
	if discount, ok := d.Discounts[discountKey]; ok {
		return discount.isStackable
	}

	return false
}

package sales

import (
	"fmt"

	"github.com/ducho-metson/polysurance-test/model"
	"github.com/ducho-metson/polysurance-test/utils"
)

type OrderInfo struct {
	price    float64
	discount float64
}

func Calculate(salesData *model.SalesData) {
	orderInfoArray := buildOrderInfo(salesData)

	totalSalesBeforeDiscount := calculateTotalSalesBeforeDiscountApplied(orderInfoArray)
	fmt.Println("Total sales before discount is applied: ", utils.RoundTo2DecimalPlaces(totalSalesBeforeDiscount))

	totalSalesAfterDiscount := calculateTotalSalesAfterDiscountApplied(orderInfoArray)
	fmt.Println("Total sales after discount is applied: ", utils.RoundTo2DecimalPlaces(totalSalesAfterDiscount))

	moneyLostViaDiscount := utils.RoundTo2DecimalPlaces(totalSalesBeforeDiscount - totalSalesAfterDiscount)
	fmt.Println("Total amount of money lost via customer using discount codes: ", moneyLostViaDiscount)

	averageDiscountPerCustomer := calculateAverageDiscountPerCustomer(orderInfoArray)
	fmt.Println("Average discount per customer as a percentage: ", utils.RoundTo2DecimalPlaces(averageDiscountPerCustomer), "%")
}

// buildOrderInfo builds a Order Info array from SalesData. Getting every order total price and discount, if exists.
func buildOrderInfo(salesData *model.SalesData) []OrderInfo {
	var orderInfoArray []OrderInfo

	for _, order := range salesData.Orders {
		orderPrice := 0.0
		for _, item := range order.Items {
			price := salesData.GetPriceFromSku(item.SKU)
			if price <= 0.0 {
				continue
			}

			orderPrice = orderPrice + price*float64(item.Quantity)
		}

		orderInfoArray = append(orderInfoArray, OrderInfo{
			price:    orderPrice,
			discount: salesData.GetDiscountAsFloat(order.Discount),
		})

	}

	return orderInfoArray
}

func calculateTotalSalesBeforeDiscountApplied(orderInfoArray []OrderInfo) float64 {
	total := 0.0
	for _, orderInfo := range orderInfoArray {
		total = total + orderInfo.price
	}

	return total
}

func calculateTotalSalesAfterDiscountApplied(orderInfoArray []OrderInfo) float64 {
	total := 0.0
	for _, orderInfo := range orderInfoArray {
		total = total + orderInfo.price*(1-orderInfo.discount)
	}

	return total
}

func calculateAverageDiscountPerCustomer(orderInfoArray []OrderInfo) float64 {
	totalDiscount := 0.0
	for _, orderInfo := range orderInfoArray {
		totalDiscount = totalDiscount + orderInfo.discount
	}

	numberCustomers := float64(len(orderInfoArray))
	return totalDiscount * 100.0 / numberCustomers
}

package lists

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices)
	prices[1] = 9.99

	prices = append(prices, 5.99, 12.99)
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99}
	prices = append(prices, discountPrices...)
}

/* type Product struct {
	title string
	id    string
	price float64
}

func main() {
	productNames := [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"
	fmt.Println(productNames)

	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))

}
*/

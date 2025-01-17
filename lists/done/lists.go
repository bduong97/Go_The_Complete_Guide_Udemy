package done

// import "fmt"

// func main() {
// 	prices := []float64{10.99, 8.99} //dynamic slice
// 	fmt.Println(prices[0:1])
// 	updatedPrices := append(prices, 5.99)
// 	fmt.Println(updatedPrices) //original prices array does not change

// }
// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := [4]float64{1, 2, 3, 4}
// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:]
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(featuredPrices)
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) 
// 	// you can always select more to the right, but can't select any from the left
// 	// cap can show potentially how much more you can expand your slice
// }
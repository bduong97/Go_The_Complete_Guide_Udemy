package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("an error occured")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var prices []float64
	for scanner.Scan() {
		convertedIntoFloat, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("could not parse floats in file")
			return
		}
		prices = append(prices, convertedIntoFloat)
	}
	job.InputPrices = prices
	
	err = scanner.Err() //once reaching the end of file
	if err != nil {
		fmt.Println("reading the file content failed.")
		fmt.Println(err)
		file.Close()
		return
	}
}
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		finalPrice, err := strconv.ParseFloat(fmt.Sprintf("%.2f", taxIncludedPrice), 64)
		if err != nil {
			fmt.Println("can't convert string to float")
			return
		}
		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = finalPrice
	}

	fmt.Println(job.TaxIncludedPrices)
}
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices:       []float64{10, 20, 30},
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]float64),
	}

}

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
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("reading the file content failed.")
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("conversion to float failed")
			fmt.Println(err)
			file.Close()
			return
		}
		prices[lineIndex] = floatPrice
		job.InputPrices = prices
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

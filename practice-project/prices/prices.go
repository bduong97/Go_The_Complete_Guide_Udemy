package prices

import (
	"fmt"
	"strconv"

	"example.com/practice-project/conversion"
	"example.com/practice-project/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {

	prices, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	job.InputPrices, err = conversion.StringsToFloat(prices)
	if err != nil {
		fmt.Println(err)
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

	filemanager.WriteJson(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices:       []float64{10, 20, 30},
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]float64),
	}

}

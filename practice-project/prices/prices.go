package prices

import (
	"fmt"
	"strconv"

	"example.com/practice-project/conversion"
	"example.com/practice-project/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadData() {

	prices, err := job.IOManager.ReadLines()
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

	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices:       []float64{10, 20, 30},
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]float64),
		IOManager:         iom,
	}

}

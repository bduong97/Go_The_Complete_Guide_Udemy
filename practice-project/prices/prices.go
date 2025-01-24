package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) Process() {
	for _, price := range job.InputPrices {
		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(job.TaxIncludedPrices)
}
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
		TaxIncludedPrices: make(map[string]float64),
	}

}
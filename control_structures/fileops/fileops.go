package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find file")
	}
	valueText := string(data)
	value, _ := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}
	return value, nil
}
func WriteFloatToFile(value float64, fileName string) {
	balanceText := fmt.Sprint(value) //converting balance to string to write to file
	os.WriteFile(fileName, []byte(balanceText), 0644)

}
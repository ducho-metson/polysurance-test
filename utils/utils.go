package utils

import (
	"fmt"
	"io"
	"math"
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed opening file %s: %w", filePath, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed reading file %s: %w", filePath, err)
	}

	return data, nil
}

func RoundTo2DecimalPlaces(value float64) float64 {
	return math.Round(value*100) / 100
}

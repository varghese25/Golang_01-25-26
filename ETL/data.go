package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// --- 1. EXTRACT ---
	inputFile, _ := os.Open("data.csv")
	defer inputFile.Close()

	reader := csv.NewReader(inputFile)
	records, _ := reader.ReadAll()

	// --- 2. TRANSFORM & 3. LOAD ---
	outputFile, _ := os.Create("transformed_data.csv")
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write Header for new file
	writer.Write([]string{"id", "item", "total_cost"})

	for i, row := range records {
		if i == 0 { continue } // Skip header row

		id := row[0]
		item := row[1]
		price, _ := strconv.ParseFloat(row[2], 64)
		tax, _ := strconv.ParseFloat(row[3], 64)

		// The Transformation
		total := fmt.Sprintf("%.2f", price+tax)

		// Loading into the new slice
		writer.Write([]string{id, item, total})
	}

	fmt.Println("ETL Complete: Check transformed_data.csv")
}
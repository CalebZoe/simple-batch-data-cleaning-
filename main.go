
import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"errors"
)

type Dataset struct {
	Headers []string
	Records [][]string
}

func main() {
	// Open the source CSV file.
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a CSV reader.
	reader := csv.NewReader(file)
	rawRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading CSV: %v\n", err)
		return
	}

	if len(rawRecords) == 0 {
		fmt.Println("Empty CSV file.")
		return
	}

	// Split headers and records.
	headers, records := rawRecords[0], rawRecords[1:]

	// Create the dataset.
	dataset := Dataset{
		Headers: headers,
		Records: cleanRecords(headers, records),
	}

	// Perform analysis.
	dataset.SummaryStats()
	dataset.UniqueValues()
	dataset.NormalizeNumericCols()

	// Write cleaned data to a new CSV file.
	if err = dataset.WriteToCSV("cleaned_data.csv"); err != nil {
		fmt.Printf("Error writing to CSV: %v\n", err)
	}
}

func cleanRecords(headers []string, records [][]string) [][]string {
	cleanedRecords := [][]string{}
	recordMap := make(map[string][]string) // To remove duplicates.

	// Find the index for the ID column.
	idIndex := -1
	for i, header := range headers {
		if strings.ToLower(header) == "id" {
			idIndex = i
			break
		}
	}

	for _, rec := range records {
		if len(rec) != len(headers) {
			continue // Skip records with missing columns.
		}

		// Trim spaces for all fields.
		for i := range rec {
			rec[i] = strings.TrimSpace(rec[i])
		}

		// Skip records with any empty fields.
		skip := false
		for _, field := range rec {
			if len(field) == 0 {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		// Remove duplicates by keeping the latest record for each ID.
		if idIndex != -1 {
			recordMap[rec[idIndex]] = rec
		} else {
			// If there's no ID column, simply append.
			cleanedRecords = append(cleanedRecords, rec)
		}
	}

	// Collect unique records if ID column was found.
	if idIndex != -1 {
		for _, rec := range recordMap {
			cleanedRecords = append(cleanedRecords, rec)
		}
	}

	return cleanedRecords
}

func (d *Dataset) SummaryStats() {
	for i, header := range d.Headers {
		numericValues := []float64{}
		for _, rec := range d.Records {
			val, err := strconv.ParseFloat(rec[i], 64)
			if err == nil {
				numericValues = append(numericValues, val)
			}
		}
		if len(numericValues) > 0 {
			mean, median := calculateMean(numericValues), calculateMedian(numericValues)
			fmt.Printf("Column: %s, Mean: %.2f, Median: %.2f\n", header, mean, median)
		}
	}
}

func calculateMean(nums []float64) float64 {
	sum := 0.0
	for _, v := range nums {
		sum += v
	}
	return sum / float64(len(nums))
}

func calculateMedian(nums []float64) float64 {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// Sort the numbers
	sort.Float64s(nums)

	if n%2 == 0 {
		return (nums[n/2-1] + nums[n/2]) / 2
	}

	return nums[n/2]
}

func (d *Dataset) UniqueValues() {
	for i, header := range d.Headers {
		uniqueValues := make(map[string]bool)
		for _, rec := range d.Records {
			uniqueValues[rec[i]] = true
		}
		fmt.Printf("Column: %s, Unique Values: %d\n", header, len(uniqueValues))
	}
}

func (d *Dataset) NormalizeNumericCols() {
	for i, header := range d.Headers {
		numericValues := []float64{}
		for _, rec := range d.Records {
			val, err := strconv.ParseFloat(rec[i], 64)
			if err == nil {
				numericValues = append(numericValues, val)
			}
		}
		if len(numericValues) > 0 {
			minVal, maxVal := min(numericValues), max(numericValues)
			for j := range d.Records {
				val, err := strconv.ParseFloat(d.Records[j][i], 64)
				if err == nil {
					normalized := (val - minVal) / (maxVal - minVal)
					d.Records[j][i] = fmt.Sprintf("%.2f", normalized)
				}
			}
			fmt.Printf("Column: %s normalized.\n", header)
		}
	}
}

func min(nums []float64) float64 {
	minVal := nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

func max(nums []float64) float64 {
	maxVal := nums[0]
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func (d *Dataset) WriteToCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	err = writer.Write(d.Headers)
	if err != nil {
		return err
	}

	// Write records
	for _, rec := range d.Records {
		err = writer.Write(rec)
		if err != nil {
			return err
		}
	}

	return nil
}

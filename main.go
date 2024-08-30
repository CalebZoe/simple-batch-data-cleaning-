import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Record struct {
	ID    string
	Name  string
	Age   string
	Email string
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
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading CSV: %v\n", err)
		return
	}

	// Process the records (skip the header).
	cleanedRecords := []Record{}
	recordMap := make(map[string]Record) // To remove duplicates.

	for i, rec := range records {
		if i == 0 {
			// Skip header.
			continue
		}

		// Create a Record object.
		r := Record{ID: rec[0], Name: rec[1], Age: rec[2], Email: rec[3]}

		// Check for missing values and normalize data.
		if len(r.ID) == 0 || len(r.Name) == 0 || len(r.Age) == 0 || len(r.Email) == 0 {
			continue // Skip records with missing values.
		}
		r.Name = strings.TrimSpace(r.Name)
		r.Email = strings.ToLower(strings.TrimSpace(r.Email))

		// Remove duplicates by keeping the latest record for each ID.
		recordMap[r.ID] = r
	}

	// Collect unique records.
	for _, rec := range recordMap {
		cleanedRecords = append(cleanedRecords, rec)
	}

	// Open the destination CSV file.
	outputFile, err := os.Create("cleaned_data.csv")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer outputFile.Close()

	// Create a CSV writer.
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Write the header.
	err = writer.Write([]string{"ID", "Name", "Age", "Email"}
	if err != nil {
		fmt.Printf("Error writing header to CSV: %v\n", err)
		return
	}
	
	// Write the cleaned records.
	for _, rec := range cleanedRecords {
		err = writer.Write([]string{rec.ID, rec.Name, rec.Age, rec.Email})
		if err != nil {
			fmt.Printf("Error writing record to CSV: %v\n", err)
			return
		}
	}

	fmt.Println("Data cleaning and CSV writing done successfully.")
}

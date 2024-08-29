package handlers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"go-webapp/utils"
	"net/http"
	"os"
	"strconv"
)

// SearchInJSON verarbeitet die Suchanfrage in JSON-Dateien
func SearchInJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	var data []int
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	queryStr := r.URL.Query().Get("query")
	query, err := strconv.Atoi(queryStr)
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	index := utils.BinarySearch(data, query)
	if index == -1 {
		fmt.Fprintf(w, "Value %d not found\n", query)
	} else {
		fmt.Fprintf(w, "Value %d found at index %d\n", query, index)
	}
}

// SearchInCSV verarbeitet die Suchanfrage in CSV-Dateien
func SearchInCSV(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	separator := r.URL.Query().Get("separator")
	if separator == "" {
		separator = ","
	}

	reader := csv.NewReader(file)
	reader.Comma = rune(separator[0])

	records, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "Error reading CSV", http.StatusBadRequest)
		return
	}

	var data []int
	for _, record := range records {
		for _, field := range record {
			num, err := strconv.Atoi(field)
			if err != nil {
				http.Error(w, "CSV contains non-integer value", http.StatusBadRequest)
				return
			}
			data = append(data, num)
		}
	}

	queryStr := r.URL.Query().Get("query")
	query, err := strconv.Atoi(queryStr)
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	index := utils.BinarySearch(data, query)
	if index == -1 {
		fmt.Fprintf(w, "Value %d not found\n", query)
	} else {
		fmt.Fprintf(w, "Value %d found at index %d\n", query, index)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// sumNumbers handles POST requests to the /sum endpoint
func sumNumbers(w http.ResponseWriter, r *http.Request) {
	var input []interface{} // Decode input as a slice of interfaces

	// Read and log the raw input
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	log.Printf("Request body: %s", string(bodyBytes))

	// Decode the JSON input
	if err := json.Unmarshal(bodyBytes, &input); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, fmt.Sprintf("Invalid JSON input: %v", err), http.StatusBadRequest)
		return
	}

	var sum int64 // Use int64 to handle larger sums
	for _, item := range input {
		var num int64
		switch v := item.(type) {
		case float64:
			// JSON numbers are unmarshalled as float64
			num = int64(v)
		case string:
			// Clean up the string by removing newline characters
			numStr := strings.ReplaceAll(v, "\n", "")
			numStr = strings.TrimSpace(numStr)

			// Convert string to integer
			parsedNum, err := strconv.ParseInt(numStr, 10, 64)
			if err != nil {
				log.Printf("Error converting string to int: %v", err)
				http.Error(w, fmt.Sprintf("Invalid number format: %v", err), http.StatusBadRequest)
				return
			}
			num = parsedNum
		default:
			log.Printf("Unsupported data type: %T", v)
			http.Error(w, "Unsupported data type in input", http.StatusBadRequest)
			return
		}
		sum += num
	}

	// Convert sum to JSON integer
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sum); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/sum", sumNumbers)

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

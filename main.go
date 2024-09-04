package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// sumNumbers menangani permintaan POST ke endpoint /sum
func sumNumbers(w http.ResponseWriter, r *http.Request) {
	var numbers []int32 // Gunakan slice string untuk mendecode body

	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		// Log kesalahan dan kirimkan respons dengan detail input yang gagal
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, fmt.Sprintf("Invalid JSON input: %v", err), http.StatusBadRequest)
		return
	}

	var sum int32
	for _, numStr := range numbers {
		// Tambah ke total
		sum += numStr
	}

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

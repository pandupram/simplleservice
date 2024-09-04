package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ResponseBody mendefinisikan struktur data yang dikirimkan ke client
type ResponseBody struct {
	Result int32 `json:"result"`
}

// sumNumbers menangani permintaan POST ke endpoint /sum
func sumNumbers(w http.ResponseWriter, r *http.Request) {
	var numbers []string // Gunakan slice string untuk mendecode body

	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var sum int32
	for _, numStr := range numbers {
		// Hapus karakter newline, spasi, dan karakter escape lainnya
		numStr = strings.ReplaceAll(numStr, "\\n", "")
		numStr = strings.TrimSpace(numStr)

		// Konversi string ke integer
		num, err := strconv.Atoi(numStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error converting '%s' to integer: %v", numStr, err), http.StatusBadRequest)
			return
		}

		// Tambah ke total
		sum += int32(num)
	}

	response := ResponseBody{
		Result: sum,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/sum", sumNumbers)

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ResponseBody mendefinisikan struktur data yang dikirimkan ke client
type ResponseBody struct {
	Result int32 `json:"result"`
}

// sumNumbers menangani permintaan POST ke endpoint /sum
func sumNumbers(w http.ResponseWriter, r *http.Request) {
	var numbers []int32
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var sum int32
	for _, num := range numbers {
		sum += num
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sum)
}

func main() {
	http.HandleFunc("/sum", sumNumbers)

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

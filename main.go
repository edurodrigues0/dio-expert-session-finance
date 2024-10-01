package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleGetTransactions)

	http.ListenAndServe(
		":8080",
		nil,
	)
}

type Transaction struct {
	Title     string  `json:"title"`
	Amount    float32 `json:"amount"`
	Type      int
	CreatedAt time.Time
}

type Transactions []Transaction

func handleGetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	layout := "2006-01-02T15:04:05"
	salaryReceived, err := time.Parse(layout, "2024-09-28T11:04:23")
	if err != nil {
		panic(err)
	}

	var transactions = Transactions{
		Transaction{
			Title:     `Salary`,
			Amount:    1200.00,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

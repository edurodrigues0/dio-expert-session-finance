package http

import (
	"net/http"

	"github.com/edurodrigues0/dio-expert-session-finance/adapter/http/actuator"
	"github.com/edurodrigues0/dio-expert-session-finance/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init() error {
	http.HandleFunc("/health", actuator.HandleHealth)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/transactions", transaction.HandleGetTransactions)
	http.HandleFunc("/transactions/create", transaction.HandleCreateTransaction)

	return http.ListenAndServe(":8080", nil)
}

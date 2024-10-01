package main

import (
	"log"

	"github.com/edurodrigues0/dio-expert-session-finance/adapter/http"
)

func main() {
	log.Fatal(http.Init())
}

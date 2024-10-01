package actuator

import (
	"net/http"

	"github.com/edurodrigues0/dio-expert-session-finance/adapter/utils"
)

type healthBody struct {
	Status string `json:"status"`
}

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.SendJSON(
			w,
			utils.ApiResponse{Error: "Method not allowed."},
			http.StatusMethodNotAllowed,
		)
		return
	}

	var status = healthBody{Status: "alive"}

	utils.SendJSON(
		w,
		utils.ApiResponse{Data: status},
		http.StatusOK,
	)
}

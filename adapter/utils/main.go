package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type ApiResponse struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func SendJSON(w http.ResponseWriter, resp ApiResponse, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Failed to marshal json data", "error", err)
		SendJSON(
			w,
			ApiResponse{Error: "Something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("Failed to write response to client", "error", err)
		return
	}
}

const layout = "2006-01-02T15:04:05"

func StringToTime(value string) time.Time {
	convertedTime, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}

	return convertedTime
}

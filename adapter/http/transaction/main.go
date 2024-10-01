package transaction

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/edurodrigues0/dio-expert-session-finance/adapter/utils"
	"github.com/edurodrigues0/dio-expert-session-finance/model/transaction"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PostBody struct {
	Title  string  `json:"title" validate:"required"`
	Amount float32 `json:"amount" validate:"required"`
	Type   int     `json:"type" validate:"required"`
}

var validate = validator.New()

func HandleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.SendJSON(
			w,
			utils.ApiResponse{Error: "Method not allowed"},
			http.StatusMethodNotAllowed,
		)
		return
	}

	var body PostBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.SendJSON(
			w,
			utils.ApiResponse{Error: "Invalid body"},
			http.StatusUnprocessableEntity,
		)

		return
	}

	if err := validate.Struct(body); err != nil {
		utils.SendJSON(
			w,
			utils.ApiResponse{Error: "Missing required fields"},
			http.StatusBadRequest,
		)
		return
	}

	transactionID, err := uuid.NewUUID()
	if err != nil {
		utils.SendJSON(
			w,
			utils.ApiResponse{Error: "Error on create uuid"},
			http.StatusBadRequest,
		)

		return
	}

	transaction := transaction.Transaction{
		ID:        transactionID,
		Title:     body.Title,
		Amount:    body.Amount,
		Type:      body.Type,
		CreatedAt: time.Now(),
	}

	utils.SendJSON(
		w,
		utils.ApiResponse{Data: transaction},
		http.StatusCreated,
	)
}

func HandleGetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.SendJSON(
			w,
			utils.ApiResponse{Error: "Method not allowed."},
			http.StatusMethodNotAllowed,
		)
		return
	}

	transactionID, _ := uuid.NewUUID()

	var transactions = transaction.Transactions{
		transaction.Transaction{
			ID:        transactionID,
			Title:     `Salary`,
			Amount:    1200.00,
			Type:      0,
			CreatedAt: utils.StringToTime("2024-09-01T11:45:26"),
		},
	}

	utils.SendJSON(
		w,
		utils.ApiResponse{Data: transactions},
		http.StatusOK,
	)
}

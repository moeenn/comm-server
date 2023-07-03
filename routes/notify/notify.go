package notify

import (
	"comm/pkg/error"
	"comm/pkg/services/auth"
	"comm/pkg/validator"
	"encoding/json"
	"net/http"
)

type NotifyRequest struct {
	UserIds []string `json:"user_ids" validate:"required"`
	// Payload string   `json:"payload"`
}

type NotifyResponse struct {
	UserId  string   `json:"userId"`
	UserIds []string `json:"user_ids"`
}

func NotifyHandler(jwtSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		decoder := json.NewDecoder(r.Body)
		encoder := json.NewEncoder(w)

		userId, err := auth.ValidateBearerToken(jwtSecret, r)
		if err != nil {
			errorResponse := error.ErrorResponse{
				Error: err.Error(),
			}

			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(errorResponse)
			return
		}

		body := NotifyRequest{}
		if err := decoder.Decode(&body); err != nil {
			errorResponse := error.ErrorResponse{
				Error: err.Error(),
			}

			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(errorResponse)
			return
		}

		if err := validator.Struct(body); len(err) != 0 {
			w.WriteHeader(http.StatusUnprocessableEntity)
			encoder.Encode(err)
			return
		}

		res := NotifyResponse{
			UserId:  userId,
			UserIds: body.UserIds,
		}

		encoder.Encode(res)
	}
}

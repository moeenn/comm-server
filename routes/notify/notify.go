package notify

import (
	"comm/pkg/error"
	"comm/pkg/validator"
	"encoding/json"
	"net/http"
)

type NotifyRequest struct {
	UserIds []string `json:"user_ids" validate:"required"`
	// Payload string   `json:"payload"`
}

type NotifyResponse struct {
	Message string   `json:"message"`
	UserId  string   `json:"user_id,omitempty"`
	UserIds []string `json:"user_ids"`
}

func NotifyHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	userId := r.Context().Value("userId").(string)
	if userId == "" {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(error.ErrorResponse{
			Error: "invalid or expired JWT",
		})
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
		Message: "you have reached notify endpoint",
		UserId:  userId,
		UserIds: body.UserIds,
	}

	encoder.Encode(res)
}

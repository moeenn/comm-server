package notify

import (
	apiError "comm/pkg/error"
	"comm/pkg/notification"
	"comm/pkg/server/utils"
	"comm/pkg/validator"
	"encoding/json"
	"net/http"
)

type NotifyRequest struct {
	UserIds []string `json:"user_ids" validate:"required"`
	Payload string   `json:"payload" validate:"required"`
}

type NotifyResponse struct {
	Message string `json:"message"`
}

func NotifyHandler(notificationChannel chan<- notification.Notification) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		isAuthenticated := r.Context().Value("isAuthenticated").(bool)
		if !isAuthenticated {
			utils.JSON(w, http.StatusUnauthorized, apiError.ErrorResponse{
				Error: "invalid or expired JWT",
			})
			return
		}

		body := NotifyRequest{}
		if err := decoder.Decode(&body); err != nil {
			utils.JSON(w, http.StatusBadRequest, apiError.ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		if err := validator.Struct(body); len(err) != 0 {
			utils.JSON(w, http.StatusUnprocessableEntity, err)
			return
		}

		if isPayloadValid := json.Valid([]byte(body.Payload)); !isPayloadValid {
			utils.JSON(w, http.StatusUnprocessableEntity, apiError.ErrorResponse{
				Error: "invalid JSON provided in payload",
			})
			return
		}

		go func(body NotifyRequest) {
			notificationChannel <- notification.Notification{
				UserIds: body.UserIds,
				Payload: []byte(body.Payload),
			}
		}(body)

		utils.JSON(w, http.StatusOK, NotifyResponse{
			Message: "notification request will be processed",
		})
	}
}

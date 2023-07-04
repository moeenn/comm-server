package notify

import (
	apiError "comm/pkg/error"
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
	Message string   `json:"message"`
	UserIds []string `json:"user_ids"`
	Payload string   `json:"payload"`
}

func NotifyHandler(w http.ResponseWriter, r *http.Request) {
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

	utils.JSON(w, http.StatusOK, NotifyResponse{
		Message: "you have reached notify endpoint",
		UserIds: body.UserIds,
		Payload: body.Payload,
	})
}

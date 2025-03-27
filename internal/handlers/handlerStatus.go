package handlers

import (
	"net/http"

	"github.com/GuiCezaF/task-collector-v2/internal/utils"
)

func HandlerStatus(w http.ResponseWriter, r *http.Request) {
	response := utils.Response{
		Message: "Server is running",
		Status:  http.StatusOK,
	}
	utils.JSONResponse(w, response, http.StatusOK)
}

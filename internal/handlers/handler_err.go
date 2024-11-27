package handlers

import (
	"net/http"

	"github.com/cyberkillua/go-api-template/internal/utils"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
}

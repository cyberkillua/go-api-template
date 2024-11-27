package handlers

import (
	"net/http"

	"github.com/cyberkillua/go-api-template/internal/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, struct{}{})
}

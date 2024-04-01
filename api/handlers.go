package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gordonBusyman/lpsh_api/storage"
)

// Settings handles the GET /settings endpoint.
func (api API) Settings(w http.ResponseWriter, r *http.Request) {
	code, err := strconv.Atoi(chi.URLParam(r, "code"))
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "invalid input")

		return
	}

	s := storage.NewUsers(api.DB)

	settings, err := s.RetrieveSettings(code)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())

		return
	}

	json.NewEncoder(w).Encode(settings)
	w.Header().Set("Content-Type", "application/json")
}

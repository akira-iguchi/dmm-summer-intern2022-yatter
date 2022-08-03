package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// Handle request for POST /v1/accounts
func (h handler) Show(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	ctx := r.Context()
	if user, err := h.app.Dao.Account().FindByUsername(ctx, username); err != nil {
		httperror.InternalServerError(w, err)
		return
	} else if user == nil {
		httperror.Error(w, http.StatusNotFound)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(user); err != nil {
			httperror.InternalServerError(w, err)
			return
		}
	}
}

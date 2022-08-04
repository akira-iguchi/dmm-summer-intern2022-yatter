package statuses

import (
	"encoding/json"
	"net/http"

	"strconv"

	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

func (h handler) Show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	statusId, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if status, err := h.app.Dao.Status().FindByStatusId(ctx, statusId); err != nil {
		httperror.InternalServerError(w, err)
		return
	} else if status == nil {
		httperror.Error(w, http.StatusNotFound)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(status); err != nil {
			httperror.InternalServerError(w, err)
			return
		}
	}
}

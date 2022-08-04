package timelines

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/handler/httperror"
)

func (h handler) Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if timelines, err := h.app.Dao.Timeline().AllStatuses(ctx); err != nil {
		httperror.InternalServerError(w, err)
		return
	} else if timelines == nil {
		httperror.Error(w, http.StatusNotFound)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(timelines); err != nil {
			httperror.InternalServerError(w, err)
			return
		}
	}
}

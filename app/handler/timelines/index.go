package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"

	"yatter-backend-go/app/handler/httperror"
)

type AddRequest struct {
	OnlyMedia bool
	MaxId     int
	SinceId   int
	Limit     int
}

func (h handler) Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	maxId, _ := strconv.Atoi(r.FormValue("max_id"))
	sinceId, _ := strconv.Atoi(r.FormValue("since_id"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))

	if timelines, err := h.app.Dao.Timeline().PublicTimelines(ctx, maxId, sinceId, limit); err != nil {
		httperror.InternalServerError(w, err)
		return
	} else if timelines == nil {
		httperror.Error(w, http.StatusBadRequest)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(timelines); err != nil {
			httperror.InternalServerError(w, err)
			return
		}
	}
}

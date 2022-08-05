package accounts

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"

	"io"
	"log"
	"os"
)

// Request body for `POST /v1/statuses`
type UpdateRequest struct {
	DisplayName *string
	Note        *string
	Avatar      *string
	Header      *string
}

func saveImage(w http.ResponseWriter, r *http.Request, v string) (string, error) {
	file, name, err := r.FormFile(v)
	if err != nil {
		return "err", err
	}
	value := name.Filename
	defer file.Close()
	filePath := filepath.Join("app", "assets", "images", "avatars", value)
	if v == "header" {
		filePath = filepath.Join("app", "assets", "images", "headers", value)
	}
	dst, err := os.Create(filePath)
	if err != nil {
		return "err", err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, file); err != nil {
		return "err", err
	}
	return value, nil
}

// Handle request for `POST /v1/statuses`
func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	formValues := new(UpdateRequest)
	queries := []string{"display_name", "note", "avatar", "header"}
	for _, v := range queries {
		value := r.FormValue(v)
		if v == "avatar" || v == "header" {
			imageValue, err := saveImage(w, r, v)
			if err != nil {
				if err.Error() == "http: no such file" {
					imageValue = ""
				} else {
					httperror.Error(w, http.StatusBadRequest)
					return
				}
			}
			value = imageValue
		}
		switch v {
		case "display_name":
			formValues.DisplayName = &value
		case "note":
			formValues.Note = &value
		case "avatar":
			formValues.Avatar = &value
		default:
			formValues.Header = &value
		}
	}

	ctx := r.Context()
	accountId := auth.AccountOf(r).ID
	if updated_status, err := h.app.Dao.Account().UpdateUser(ctx, accountId, formValues.DisplayName, formValues.Note, formValues.Avatar, formValues.Header); err != nil {
		httperror.InternalServerError(w, err)
		return
	} else if updated_status == nil {
		httperror.Error(w, http.StatusBadRequest)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(updated_status); err != nil {
			httperror.InternalServerError(w, err)
			return
		}
	}
}

package pfmapi

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/govinda-attal/dlt-pfm/pkg/core/status"
	"github.com/govinda-attal/dlt-pfm/pkg/pfm"
)

type AssetHandler struct {
	api pfm.AssetAPI
}

func NewAssetHandler(api pfm.AssetAPI) *AssetHandler {
	return &AssetHandler{api}
}

func (ah *AssetHandler) FetchAsset(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return status.ErrBadRequest
	}
	_, err := ah.api.Get(id)
	if err != nil {
		return err
	}
	return nil
}

func (ah *AssetHandler) AddAsset(w http.ResponseWriter, r *http.Request) error {
	return status.ErrNotImplemented
}

// WrapperHandler is wrapper function to wrap API handlers and retuns as http.HandlerFunc.
// API Handlers may return error, and this wrapper simplifies error handling for API Handlers.
func WrapperHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			if errSvc, ok := err.(status.ErrServiceStatus); ok {
				w.WriteHeader(errSvc.Code)
				json.NewEncoder(w).Encode(&errSvc)
				return
			}
			errSvc := status.ErrInternal.WithMessage(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&errSvc)
		}
	}
}

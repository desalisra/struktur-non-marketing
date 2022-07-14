package city

import (
	"log"
	"net/http"
	"struktur-non-marketing/pkg/response"

	"github.com/gorilla/mux"

	httpHelper "struktur-non-marketing/internal/delivery/http"
)

func (h *Handler) GetCity(w http.ResponseWriter, r *http.Request) {
	var err error
	var result interface{}

	resp := response.Response{}
	defer resp.RenderJSON(w, r)
	ctx := r.Context()

	params := r.URL.Query()
	searchBy := params["type"][0]
	searchValue := params["value"][0]

	switch searchBy {
	case "id":
		result, err = h.service.GetCityById(ctx, searchValue)
		if err != nil {
			resp = httpHelper.ParseErrorCode(err.Error())
			log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
			return
		}
	case "name":
		result, err = h.service.GetCityByName(ctx, searchValue)
		if err != nil {
			resp = httpHelper.ParseErrorCode(err.Error())
			log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
			return
		}
	}

	resp.Data = result
	resp.Metadata = "metadata"
	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) GetBranch(w http.ResponseWriter, r *http.Request) {
	var err error
	var result interface{}

	resp := response.Response{}
	defer resp.RenderJSON(w, r)
	ctx := r.Context()

	params1 := mux.Vars(r)
	cityID := params1["id"]

	params2 := r.URL.Query()
	searchBy := params2["type"][0]
	searchValue := params2["value"][0]

	switch searchBy {
	case "id":
		result, err = h.service.GetCityBranchByID(ctx, cityID, searchValue)
		if err != nil {
			resp = httpHelper.ParseErrorCode(err.Error())
			log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
			return
		}
	case "name":
		result, err = h.service.GetCityBranchByName(ctx, cityID, searchValue)
		if err != nil {
			resp = httpHelper.ParseErrorCode(err.Error())
			log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
			return
		}
	}

	resp.Data = result
	resp.Metadata = "metadata"
	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

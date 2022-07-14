package groupteri

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"struktur-non-marketing/pkg/response"

	httpHelper "struktur-non-marketing/internal/delivery/http"
	entity "struktur-non-marketing/internal/entity/groupteri"
)

func (h *Handler) GetStrukturTeri(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)
	ctx := r.Context()

	params := r.URL.Query()
	periode := params["periode"][0]
	ptID := params["pt_id"][0]
	dptID := params["dpt_id"][0]
	karNip := params["nip"][0]

	result, err := h.service.GetStruktur(ctx, periode, ptID, dptID, karNip)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = "metadata"
	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) AddStrukturTeri(w http.ResponseWriter, r *http.Request) {
	var request entity.Grpteri

	resp := response.Response{}
	defer resp.RenderJSON(w, r)
	ctx := r.Context()

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &request)

	result, err := h.service.AddStrukturTeri(ctx, request)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = "metadata"
	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) EditStrukturTeri(w http.ResponseWriter, r *http.Request) {
	var request entity.Grpteri

	resp := response.Response{}
	defer resp.RenderJSON(w, r)
	ctx := r.Context()

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &request)

	result, err := h.service.EditStrukturTeri(ctx, request)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = "metadata"
	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) DeleteStrukturTeri(w http.ResponseWriter, r *http.Request) {
	var request entity.Grpteri

	resp := response.Response{}
	defer resp.RenderJSON(w, r)
	ctx := r.Context()

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &request)

	result, err := h.service.DeleteStrukturTeri(ctx, request)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = "metadata"
	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

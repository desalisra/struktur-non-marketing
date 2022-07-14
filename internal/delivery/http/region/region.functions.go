package region

import (
	"log"
	"net/http"
	"struktur-non-marketing/pkg/response"

	httpHelper "struktur-non-marketing/internal/delivery/http"
)

func (h *Handler) GetStrukturRegion(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)
	ctx := r.Context()

	params := r.URL.Query()
	periode := params["periode"][0]
	ptID := params["pt_id"][0]
	dptID := params["dpt_id"][0]
	karNip := params["nip"][0]

	result, err := h.service.GetStrukturRegion(ctx, periode, ptID, dptID, karNip)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = "metadata"
	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

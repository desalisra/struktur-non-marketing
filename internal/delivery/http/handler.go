package http

import (
	"errors"
	"log"
	"net/http"

	"struktur-non-marketing/pkg/response"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// Handler will initialize mux router and register handler
func (s *Server) Handler() *mux.Router {
	r := mux.NewRouter()
	// Jika tidak ditemukan, jangan diubah.
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	// Health Check
	r.HandleFunc("", defaultHandler).Methods("GET")
	r.HandleFunc("/", defaultHandler).Methods("GET")

	// Tambahan Prefix di depan API endpoint
	router := r.PathPrefix("/struktur-non-mkt").Subrouter()

	// Routes
	master := router.PathPrefix("/master").Subrouter()
	// skeleton.Use(s.JWTMiddleware)

	master.HandleFunc("/departments", s.Department.GetDepartments).Methods("GET")
	master.HandleFunc("/department/{id}", s.Department.GetDepartmentById).Methods("GET")
	master.HandleFunc("/position/{id}", s.Department.GetPosition).Methods("GET")

	master.HandleFunc("/city", s.City.GetCity).Methods("GET")
	master.HandleFunc("/branch/{id}", s.City.GetBranch).Methods("GET")

	master.HandleFunc("/jabatan-iklan", s.Iklan.GetJabIklan).Methods("GET")
	master.HandleFunc("/karyawan", s.Karyawan.GetKaryawan).Methods("GET")

	// Routes path MR (GroupTeri)
	grpt := router.PathPrefix("/mr").Subrouter()
	grpt.HandleFunc("/struktur", s.GroupTeri.GetStrukturTeri).Methods("GET")
	grpt.HandleFunc("/struktur", s.GroupTeri.AddStrukturTeri).Methods("POST")
	grpt.HandleFunc("/struktur", s.GroupTeri.EditStrukturTeri).Methods("PUT")
	grpt.HandleFunc("/struktur", s.GroupTeri.DeleteStrukturTeri).Methods("DELETE")

	// Routes path SPV (SubArea)
	sub := router.PathPrefix("/spv").Subrouter()
	sub.HandleFunc("/struktur", s.SubArea.GetStrukturSubarea).Methods("GET")

	// Routes path ASM (Area)
	asm := router.PathPrefix("/asm").Subrouter()
	asm.HandleFunc("/struktur", s.Area.GetStrukturSubarea).Methods("GET")

	// Routes path SM (Region)
	sm := router.PathPrefix("/sm").Subrouter()
	sm.HandleFunc("/struktur", s.Region.GetStrukturRegion).Methods("GET")

	// Routes path NSM
	nsm := router.PathPrefix("/nsm").Subrouter()
	nsm.HandleFunc("/struktur", s.Nsm.GetStrukturNsm).Methods("GET")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return r
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Example Service API"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp   *response.Response
		err    error
		errRes response.Error
	)
	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	err = errors.New("404 Not Found")

	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   404,
			Msg:    "404 Not Found",
			Status: true,
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.StatusCode = 404
		resp.Error = errRes
		return
	}
}

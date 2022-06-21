package http

import (
	"net/http"

	"struktur-non-marketing/pkg/grace"

	"github.com/rs/cors"
)

// DepartmentHandler ...
type DepartmentHandler interface {
	GetDepartments(w http.ResponseWriter, r *http.Request)
	GetDepartmentById(w http.ResponseWriter, r *http.Request)
	GetPosition(w http.ResponseWriter, r *http.Request)
}

// CityHandler ...
type CityHandler interface {
	GetCitys(w http.ResponseWriter, r *http.Request)
	GetCity(w http.ResponseWriter, r *http.Request)
	GetBranchByCityId(w http.ResponseWriter, r *http.Request)
}

// IklanHandler ...
type IklanHandler interface {
	GetJabIklan(w http.ResponseWriter, r *http.Request)
}

// KaryawanHandler ...
type KaryawanHandler interface {
	GetKaryawan(w http.ResponseWriter, r *http.Request)
}

// GroupTeriHandler ...
type GroupTeriHandler interface {
	GetStrukturTeri(w http.ResponseWriter, r *http.Request)
}

// SubAreaHandler ...
type SubAreaHandler interface {
	GetStrukturSubarea(w http.ResponseWriter, r *http.Request)
}

// AreaHandler ...
type AreaHandler interface {
	GetStrukturSubarea(w http.ResponseWriter, r *http.Request)
}

// RegionHandler ...
type RegionHandler interface {
	GetStrukturRegion(w http.ResponseWriter, r *http.Request)
}

// NsmHandler ...
type NsmHandler interface {
	GetStrukturNsm(w http.ResponseWriter, r *http.Request)
}



// Server ...
type Server struct {
	// server   *http.Server
	Department	DepartmentHandler
	City		CityHandler
	Iklan		IklanHandler
	Karyawan	KaryawanHandler
	GroupTeri	GroupTeriHandler
	SubArea		SubAreaHandler
	Area		AreaHandler
	Region		RegionHandler
	Nsm			NsmHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}

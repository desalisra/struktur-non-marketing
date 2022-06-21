package boot

import (
	"log"
	"net/http"
	"struktur-non-marketing/docs"
	"struktur-non-marketing/pkg/httpclient"

	"struktur-non-marketing/internal/config"
	"struktur-non-marketing/internal/viper"

	"github.com/jmoiron/sqlx"

	server "struktur-non-marketing/internal/delivery/http"

	authData "struktur-non-marketing/internal/data/auth"
	authService "struktur-non-marketing/internal/service/auth"

	departmentData "struktur-non-marketing/internal/data/department"
	departmentHandler "struktur-non-marketing/internal/delivery/http/department"
	departmentService "struktur-non-marketing/internal/service/department"

	cityData "struktur-non-marketing/internal/data/city"
	cityHandler "struktur-non-marketing/internal/delivery/http/city"
	cityService "struktur-non-marketing/internal/service/city"

	iklanData "struktur-non-marketing/internal/data/jabiklan"
	iklanHandler "struktur-non-marketing/internal/delivery/http/jabiklan"
	iklanService "struktur-non-marketing/internal/service/jabiklan"

	kryData "struktur-non-marketing/internal/data/karyawan"
	kryHandler "struktur-non-marketing/internal/delivery/http/karyawan"
	kryService "struktur-non-marketing/internal/service/karyawan"

	// client sementara
	clientData "struktur-non-marketing/internal/data/data_client"

	grpteriData "struktur-non-marketing/internal/data/groupteri"
	grpteriHandler "struktur-non-marketing/internal/delivery/http/groupteri"
	grpteriService "struktur-non-marketing/internal/service/groupteri"

	subData "struktur-non-marketing/internal/data/subarea"
	subHandler "struktur-non-marketing/internal/delivery/http/subarea"
	subService "struktur-non-marketing/internal/service/subarea"

	areaData "struktur-non-marketing/internal/data/area"
	areaHandler "struktur-non-marketing/internal/delivery/http/area"
	areaService "struktur-non-marketing/internal/service/area"

	regData "struktur-non-marketing/internal/data/region"
	regHandler "struktur-non-marketing/internal/delivery/http/region"
	regService "struktur-non-marketing/internal/service/region"

	nsmData "struktur-non-marketing/internal/data/nsm"
	nsmHandler "struktur-non-marketing/internal/delivery/http/nsm"
	nsmService "struktur-non-marketing/internal/service/nsm"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}

	errViper := viper.Init()
	if errViper != nil {
		log.Fatalf("[CONFIG2] Failed to initialize config: %v", errViper)
	}

	v := viper.GetConn()
	db := v.Db

	cfg := config.Get()
	
	// Open MySQL DB Connection
	db, err = sqlx.Open("mysql", cfg.Database.Master)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	//
	docs.SwaggerInfo.Host = cfg.Swagger.Host
	docs.SwaggerInfo.Schemes = cfg.Swagger.Schemes

	httpc := httpclient.NewClient()
	ad := authData.New(httpc, cfg.API.Auth)
	as := authService.New(ad)

	//Sementara
	_ = as 

	dptd := departmentData.New(db)
	dpts := departmentService.New(dptd)
	dpth := departmentHandler.New(dpts)

	cityd := cityData.New(db)
	citys := cityService.New(cityd)
	cityh := cityHandler.New(citys)

	ikland := iklanData.New(db)
	iklans := iklanService.New(ikland)
	iklanh := iklanHandler.New(iklans)

	kryd := kryData.New(db)
	krys := kryService.New(kryd)
	kryh := kryHandler.New(krys)

	cd := clientData.New(httpc, cfg.API.Client)
	
	grpterid := grpteriData.New(db)
	grpteris := grpteriService.New(grpterid, cd)
	grpterih := grpteriHandler.New(grpteris)

	subd := subData.New(db)
	subs := subService.New(subd, cd)
	subh := subHandler.New(subs)

	aread := areaData.New(db)
	areas := areaService.New(aread, cd)
	areah := areaHandler.New(areas)

	regd := regData.New(db)
	regs := regService.New(regd, cd)
	regh := regHandler.New(regs)

	nsmd := nsmData.New(db)
	nsms := nsmService.New(nsmd, cd)
	nsmh := nsmHandler.New(nsms)


	s := server.Server{
		Department: dpth,
		City: cityh,
		Iklan: iklanh,
		Karyawan: kryh,
		GroupTeri: grpterih,
		SubArea: subh,
		Area: areah,
		Region: regh,
		Nsm: nsmh,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}

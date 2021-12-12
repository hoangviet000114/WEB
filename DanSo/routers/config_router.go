package routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"viet/test/api"
)

func RunServer() {
	r := mux.NewRouter()
	customerURL := r.PathPrefix("/api").Subrouter()
	customerURL.HandleFunc("/nguoidanS", api.LayDSNguoiDan).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	fmt.Println("Server start on domain: http://localhost:10000")
	log.Fatal(http.ListenAndServe(":10000", handler))
}

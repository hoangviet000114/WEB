package routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"viet/test/api"
	"viet/test/api/api_captaikhoan"
	"viet/test/api/api_khaibao"
	"viet/test/api/api_themxoasua_nguoidan"
)

func RunServer() {
	r := mux.NewRouter()
	routerChung := r.PathPrefix("/api").Subrouter()
	routerChung.HandleFunc("/login", api.Login).Methods(http.MethodPost)
	routerChung.HandleFunc("/them_dan", api_themxoasua_nguoidan.ThemNguoiDan).Methods(http.MethodPost)
	routerChung.HandleFunc("/getAllDan", api_themxoasua_nguoidan.LayDSNguoiDan).Methods(http.MethodPost)
	routerChung.HandleFunc("/ds", api.LayDSTinhHuyenXaXom).Methods(http.MethodPost)
	routerChung.HandleFunc("/khaibao", api_khaibao.KhaiBaoTinh).Methods(http.MethodPost)
	routerChung.HandleFunc("/captaikhoan", api_captaikhoan.CapTaiKhoan).Methods(http.MethodPost)
	routerChung.HandleFunc("/xoanguoidan", api_themxoasua_nguoidan.XoaNguoiDan).Methods(http.MethodPost)
	routerChung.HandleFunc("/suanguoidan", api_themxoasua_nguoidan.SuaNguoiDan).Methods(http.MethodPost)
	routerChung.HandleFunc("/baocaohoanthanh", api.BaoCaoHoanThanh).Methods(http.MethodPost)
	routerChung.HandleFunc("/capquyenkhaibao", api.CapQuyenKhaiBao).Methods(http.MethodPost)
	routerChung.HandleFunc("/logout", api.Logout).Methods(http.MethodGet)
	routerChung.HandleFunc("/thongke", api.ThongKe).Methods(http.MethodPost)

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

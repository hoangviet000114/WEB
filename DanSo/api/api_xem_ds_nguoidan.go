package api

import (
	"encoding/json"
	"net/http"
	"viet/test/DB_process"
)

func LayDSNguoiDan(w http.ResponseWriter, r *http.Request) {
	nguoidanS, err := DB_process.Get_NguoiDan()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"ds_nguoidan": nguoidanS})
}

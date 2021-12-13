package api_themxoasua_nguoidan

import (
	"encoding/json"
	"net/http"
	"viet/test/DB_process"
	"viet/test/models"
)

func XoaNguoiDan(w http.ResponseWriter, r *http.Request) {
	var reqData models.XoaNguoiDanRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": err.Error()})
		return
	}

	err := DB_process.XoaNguoiDanInDB(&reqData)
	if err != nil {
		if err.Error() != "success" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"message": err.Error()})
			return
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"message": "success"})
		}
	}
}

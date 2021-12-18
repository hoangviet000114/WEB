package api

import (
	"encoding/json"
	"net/http"
	"viet/test/DB_process/DB_thongke"
	"viet/test/models"
)

func ThongKe(w http.ResponseWriter, r *http.Request) {
	var reqData models.ThongKeRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": err.Error()})
		return
	}

	if len(reqData.ID) == 0 || len(reqData.ID) == 2 || len(reqData.ID) == 4 || len(reqData.ID) == 6 || len(reqData.ID) == 8 {
		result, err := DB_thongke.ThongKeAll(&reqData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"message": err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"nam":          result[0].Nam,
			"nu":           result[0].Nu,
			"thathoc":      result[0].ThatHoc,
			"tieuhoc":      result[0].TieuHoc,
			"thcs":         result[0].THCS,
			"thpt":         result[0].THPT,
			"cotongiao":    result[0].CoTonGiao,
			"khongtongiao": result[0].KhongTonGiao,
			"message":      "success"})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Sai ID"})
		return
	}
}

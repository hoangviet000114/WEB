package api

import (
	"encoding/json"
	"net/http"
	"viet/test/DB_process"
)

func Login(w http.ResponseWriter, r *http.Request) {

	//Kiểm tra username, password có rỗng không

	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	nguoidungS, err := DB_process.Get_NguoiDung()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": err.Error()})
		return
	}

	//Tìm username có tồn tại trong db chưa
	//Nếu có,
}

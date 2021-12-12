package models

type NguoiDan struct {
	CCCD            string `json:"cccd/cmnd"`
	HoTen           string `json:"hoten"`
	GioiTinh        string `json:"gioitinh"`
	NgaySinh        string `json:"ngaysinh"`
	QueQuan         string `json:"quequan"`
	DiaChiThuongTru string `json:"diachi_thuongtru"`
	DiaChiTamTru    string `json:"diachi_tamtru"`
	TonGiao         string `json:"tongiao"`
	TrinhDo         string `json:"trinhdo"`
	NgheNghiep      string `json:"nghenghiep"`
	ID_Xom          string `json:"id_xom"`
}
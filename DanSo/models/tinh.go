package models

type Tinh struct {
	ID_Tinh         string `json:"id_tinh"`
	TenTinh         string `json:"ten_tinh"`
	TrangThai       string `json:"trang_thai"`
	ThoiDiemBatDau  string `json:"thoi_diem_bat_dau"`
	ThoiDiemKetThuc string `json:"thoi_diem_ket_thuc"`
}
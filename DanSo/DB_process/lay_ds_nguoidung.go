package DB_process

import (
	"log"
	"viet/test/database"
	"viet/test/models"
)

func Get_NguoiDung() ([]models.NguoiDung, error) {
	db := database.Connect()
	defer db.Close()

	query := "SELECT * FROM nguoidung"

	var nguoidungS []models.NguoiDung
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error: ", err)
		return nguoidungS, err
	}
	for rows.Next() {
		var nguoidung models.NguoiDung
		err = rows.Scan(
			&nguoidung.TaiKhoan,
			&nguoidung.MatKhau,
			&nguoidung.HoTen,
			&nguoidung.SDT,
			&nguoidung.Email,
			&nguoidung.DiaChi,
		)
		if err != nil {
			log.Println("ERROR: ", err)
			return nguoidungS, err
		}
		nguoidungS = append(nguoidungS, nguoidung)
	}

	return nguoidungS, err
}

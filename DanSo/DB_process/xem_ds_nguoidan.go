package DB_process

import (
	"log"
	"viet/test/database"
	"viet/test/models"
)

func Get_NguoiDan() ([]models.NguoiDan, error) {
	db := database.Connect()
	defer db.Close()

	query := "SELECT * FROM nguoidan"

	var nguoidanS []models.NguoiDan
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error: ", err)
		return nguoidanS, err
	}
	for rows.Next() {
		var nguoidan models.NguoiDan
		err = rows.Scan(
			&nguoidan.CCCD,
			&nguoidan.HoTen,
			&nguoidan.GioiTinh,
			&nguoidan.NgaySinh,
			&nguoidan.QueQuan,
			&nguoidan.DiaChiThuongTru,
			&nguoidan.DiaChiTamTru,
			&nguoidan.TonGiao,
			&nguoidan.TrinhDo,
			&nguoidan.NgheNghiep,
			&nguoidan.ID_Xom,
		)
		if err != nil {
			log.Println("ERROR: ", err)
			return nguoidanS, err
		}
		nguoidanS = append(nguoidanS, nguoidan)
	}

	return nguoidanS, err
}

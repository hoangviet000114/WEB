package DB_thongke

import (
	"log"
	"strconv"
	s "strings"
	"viet/test/database"
	"viet/test/models"
)

func ThongKeAll(request *models.ThongKeRequest) ([]models.KetQuaThongKeRequest, error) {
	db := database.Connect()
	defer db.Close()

	query := "SELECT * FROM nguoidan"

	var result []models.KetQuaThongKeRequest
	var temp models.KetQuaThongKeRequest

	temp.Nam = "0"
	temp.Nu = "0"
	temp.ThatHoc = "0"
	temp.TieuHoc = "0"
	temp.THCS = "0"
	temp.THPT = "0"
	temp.CoTonGiao = "0"
	temp.KhongTonGiao = "0"
	result = append(result, temp)

	var Nam = 0
	var Nu = 0
	var ThatHoc = 0
	var TieuHoc = 0
	var THCS = 0
	var THPT = 0
	var CoTonGiao = 0
	var KhongTonGiao = 0

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error: ", err)
		return result, err
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
			return result, err
		}
		if s.HasPrefix(nguoidan.ID_Xom, request.ID) {
			if nguoidan.GioiTinh == "Nam" {
				Nam++
			} else {
				Nu++
			}

			if nguoidan.TrinhDo == "Thất học" {
				ThatHoc++
			} else if nguoidan.TrinhDo == "Tiểu học" {
				TieuHoc++
			} else if nguoidan.TrinhDo == "THCS" {
				THCS++
			} else if nguoidan.TrinhDo == "THPT" {
				THPT++
			}

			if nguoidan.TonGiao == "Không" {
				KhongTonGiao++
			} else {
				CoTonGiao++
			}
		}
	}

	result[0].Nam = strconv.Itoa(Nam)
	result[0].Nu = strconv.Itoa(Nu)
	result[0].ThatHoc = strconv.Itoa(ThatHoc)
	result[0].TieuHoc = strconv.Itoa(TieuHoc)
	result[0].THCS = strconv.Itoa(THCS)
	result[0].THPT = strconv.Itoa(THPT)
	result[0].CoTonGiao = strconv.Itoa(CoTonGiao)
	result[0].KhongTonGiao = strconv.Itoa(KhongTonGiao)

	return result, err
}

package DB_capquyen_khaibao

import (
	"fmt"
	"viet/test/database"
	"viet/test/models"
)

func CapQuyenTKInDB(request *models.CapQuyenKhaiBaoTKRequest) error {
	db := database.Connect()
	defer db.Close()

	_, err := db.Query("UPDATE nguoidung SET quyen = ? WHERE tai_khoan = ?",
		request.Quyen, request.ID)

	if err != nil {
		return err
	}
	return fmt.Errorf("success")

}

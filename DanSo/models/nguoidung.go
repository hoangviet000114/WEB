package models

import (
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
)

type NguoiDung struct {
	TaiKhoan string `json:"tai_khoan"`
	MatKhau  string `json:"mat_khau"`
	HoTen    string `json:"ho_ten"`
	SDT      string `json:"sdt"`
	Email    string `json:"email"`
	DiaChi   string `json:"dia_chi"`
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}


func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}

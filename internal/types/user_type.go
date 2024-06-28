package types

import (
	"time"
)

type FindUser struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	NoHp      string    `json:"no_hp"`
	HakAkses  string    `json:"hak_akses"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	NoHp     string `json:"no_hp"`
	HakAkses string `json:"hak_akses"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	NoHp     string `json:"no_hp"`
	HakAkses string `json:"hak_akses"`
	Password string `json:"password"`
}

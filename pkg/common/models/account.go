package models

type Account struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Holder string `json:"holder"`
	Email  string `json:"email"`
}

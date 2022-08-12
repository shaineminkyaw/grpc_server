package model

import "time"

type User struct {
	ID             uint64    `gorm:"column:id;primary_key,AUTO_INCREMENT" json:"id"`
	Username       string    `gorm:"column:username" jsn:"username"`
	Password       string    `gorm:"column:password" json:"password"`
	Email          string    `gorm:"column:email" json:"email"`
	PhoneAreaCode  string    `gorm:"column:phone_area_code" json:"phone_area_code"`
	PhoneNumber    string    `gorm:"column:phone_number" json:"phone_number"`
	Balance        float64   `gorm:"column:balance" json:"balance"`
	Currency       string    `gorm:"column:currency" json:"currency"`
	BankCardNumber string    `gorm:"column:bank_crad_number" json:"bank_card_number"`
	NationID       string    `gorm:"nation_id" json:"nation_id"`
	City           string    `gorm:"column:city" json:"city"`
	RegesiterIP    string    `gorm:"column:register_ip" json:"register_ip"`
	LastLoginIP    string    `gorm:"column:last_login_ip" json:"last_login_ip"`
	Type           int8      `gorm:"gender_type" json:"gender_type"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt      time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (u *User) TableName() string {
	return "bank_user"
}

type VerifyCode struct {
	Id         uint64    `gorm:"column:id" json:"id"`
	Email      string    `gorm:"column:email" json:"email"`
	Code       string    `gorm:"column:code" json:"code"`
	ExpireTime time.Time `gorm:"column:expire_time" json:"expire_time"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (uv *VerifyCode) TableName() string {
	return "user_verify_code"
}

// user token
type UserToken struct {
	Id         uint64    `gorm:"column:id" json:"id"`
	Uid        uint64    `gorm:"column:uid" json:"uid"`
	TokenID    string    `gorm:"column:token_id" json:"token_id"`
	Token      string    `gorm:"column:token" json:"token"`
	ExpireTime time.Time `gorm:"column:expire_time" json:"expire_time"`
	CreatedAt  time.Time `gorm:"column:creaed_at" json:"created_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (ut *UserToken) TableName() string {
	return "user_token"
}

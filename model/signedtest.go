package model

import "time"

type SignedTest struct {
	Id        int    `gorm:"type:int;primary_key"`
	UserId    string `gorm:"type:uuid;index:user_sign"`
	Questions string `gorm:"type:text"`
	Answers   string `gorm:"type:text"`
	Signature string `gorm:"type:char(32);index:user_sign;unique;not null"`
	Timestamp time.Time
}

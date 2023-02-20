package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Uid      string `json:"uid" gorm:"primaryKey,type:char(27)" `    // 用户ID
	Name     string `json:"name" form:"name" gorm:"size:64"`         // 名称
	Email    string `json:"email" form:"email" gorm:"index,size:64"` // 邮箱
	Password string `json:"password" form:"password"`                // 密码

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

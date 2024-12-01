// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID              int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	AccountID       int64     `gorm:"column:account_id;type:bigint;not null;uniqueIndex:idx_users_account_id,priority:1" json:"account_id"`
	Mobile          string    `gorm:"column:mobile;type:varchar(20);not null" json:"mobile"`
	Email           string    `gorm:"column:email;type:varchar(50);not null" json:"email"`
	Name            string    `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Avatar          string    `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`
	BackgroundImage string    `gorm:"column:background_image;type:varchar(255);not null" json:"background_image"`
	Signature       string    `gorm:"column:signature;type:varchar(255);not null" json:"signature"`
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
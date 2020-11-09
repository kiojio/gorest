package models

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	RoleId   uint32 `gorm:"primary_key;auto_increment" json:"role_id"`
	RoleName string `gorm:"size:100;not null" json:"role_name"`
}

func (u *Role) SaveUser(db *gorm.DB) (*Role, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Role{}, err
	}
	return u, nil
}

func (u *Role) FindAllUsers(db *gorm.DB) (*[]Role, error) {
	var err error
	roles := []Role{}
	err = db.Debug().Model(&Role{}).Limit(100).Find(&roles).Error
	if err != nil {
		return &[]Role{}, err
	}
	return &roles, err
}

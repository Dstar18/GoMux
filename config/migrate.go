package config

import "GoMux/entity"

func InitMigrate() {
	DB.AutoMigrate(&entity.User{})
}

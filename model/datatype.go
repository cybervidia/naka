package model

import (
	"gorm.io/gorm"
)

type SecretEntry struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Mail     string //nome o mail associato alla password o al segreto
	Password string //password o segreto da salvare criptato!!!
	Note     string
	IV       string //initilization vector
}

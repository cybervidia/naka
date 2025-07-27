/*
Copyright © 2025 maKs <eliteKnow@theyKnowWhere.it>
*/
package model

import (
	"gorm.io/gorm"
)

type SecretEntry struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Tag      string //optional se non c'è il --tag o -t il suo valore è "#default"
	Mail     string //nome o mail associato alla password o al segreto
	Password string //password o segreto da salvare criptato!!!
	Note     string
	IV       string //initilization vector
	Salt     string
}

/*
Copyright © 2025 maKs <eliteKnow@theyKnowWhere.it>
*/
package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/cybervidia/naka/model"
	"github.com/cybervidia/naka/vault"
	"github.com/pterm/pterm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func AddSecret(secret *model.SecretEntry) {

	dbPath, err := getDatabasePath()
	if err != nil {
		log.Fatalf("Failed to get database path: %v", err)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.SecretEntry{})

	//db.Create(&bkmrk)
	result := db.Create(&secret)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "UNIQUE constraint failed") {
			fmt.Println("Error: name already in use")
		} else {
			fmt.Println("Generic error:", result.Error)
		}
		return // o os.Exit(1)
	}

	fmt.Println("Secret <", secret.Name, "> successfully inserted")

}

func ListSecret(tag string) {

	dbPath, err := getDatabasePath()
	if err != nil {
		log.Fatalf("Failed to get database path: %v", err)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.SecretEntry{})

	var secrets []model.SecretEntry

	var result *gorm.DB

	if tag == "" {
		result = db.Find(&secrets)
	} else {
		result = db.Where("tag = ?", tag).Find(&secrets)
	}

	err = result.Error // returns error

	if err != nil {
		fmt.Println("argh something wrong", err)
	}

	// Creazione della tabella
	tableData := pterm.TableData{
		{"Name", "Mail", "Password", "Note"},
	}

	for _, secret := range secrets {
		row := []string{
			secret.Name,
			secret.Mail,
			secret.Password,
			secret.Note,
		}
		tableData = append(tableData, row)
	}

	alternateStyle := pterm.NewStyle(pterm.BgDarkGray)
	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).WithAlternateRowStyle(alternateStyle).Render()
}

func GetSecret(name string) {

	dbPath, err := getDatabasePath()
	if err != nil {
		log.Fatalf("Failed to get database path: %v", err)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.SecretEntry{})

	secret := model.SecretEntry{}
	db.First(&secret, "name = ?", name) // carica il record

	//qui devo decriptare la pwd
	vault.Unlock(&secret)

	err = clipboard.WriteAll(secret.Password)
	if err != nil {
		fmt.Println("Error in coping to clipboard:", err)
		return
	}
	fmt.Println("Password copied in clipboard:") //, secret.Password)
}

func DeleteSecret(name string) {

	dbPath, err := getDatabasePath()
	if err != nil {
		log.Fatalf("Failed to get database path: %v", err)
	}

	//Open DB
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.SecretEntry{})

	secret := model.SecretEntry{}
	db.First(&secret, "name = ?", name) // carica il record
	db.Unscoped().Delete(&model.SecretEntry{}, secret.ID)

	fmt.Println("secret deleted:", secret.Name)
}

func UpdateSecret(secret *model.SecretEntry) {
	dbPath, err := getDatabasePath()
	if err != nil {
		log.Fatalf("Failed to get database path: %v", err)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.SecretEntry{})

	oldSecret := model.SecretEntry{}
	db.First(&oldSecret, "name = ?", secret.Name) // carica il record
	oldSecret.Name = secret.Name
	oldSecret.Mail = secret.Mail
	oldSecret.Password = secret.Password
	oldSecret.Note = secret.Note
	oldSecret.IV = secret.IV
	oldSecret.Salt = secret.Salt

	db.Save(&oldSecret)
}

/*
Helper method that retrn the path where the executable live with a db fine added at the end
*/
func getDatabasePath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir := filepath.Dir(exePath)
	dbPath := filepath.Join(exeDir, ".secret.db")
	// fmt.Println(dbPath)
	return dbPath, nil
}

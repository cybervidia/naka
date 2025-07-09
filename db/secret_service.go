package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cybervidia/naka/model"
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
			fmt.Println("Errore: nome duplicato")
		} else {
			fmt.Println("Errore generico:", result.Error)
		}
		return // o os.Exit(1)
	}

	fmt.Println("Secret <", secret.Name, "> inserito con successo")

}

func getDatabasePath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir := filepath.Dir(exePath)
	dbPath := filepath.Join(exeDir, "secret.db")
	fmt.Println(dbPath)
	return dbPath, nil
}

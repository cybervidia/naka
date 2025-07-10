package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/cybervidia/naka/model"
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
			fmt.Println("Errore: nome duplicato")
		} else {
			fmt.Println("Errore generico:", result.Error)
		}
		return // o os.Exit(1)
	}

	fmt.Println("Secret <", secret.Name, "> inserito con successo")

}

func ListSecret() {
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

	result := db.Find(&secrets)

	// codice di esempio per list da bookmark app in go, poi CANCELLALO!!!!!

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
			//fmt.Sprintf("%d", secret.ID),
			secret.Name,
			secret.Mail,
			secret.Password,
			secret.Name,
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

	err = clipboard.WriteAll(secret.Password)
	if err != nil {
		fmt.Println("Errore nel copiare nella clipboard:", err)
		return
	}
	fmt.Println("Testo copiato nella clipboard:", secret.Password)
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
	dbPath := filepath.Join(exeDir, "secret.db")
	fmt.Println(dbPath)
	return dbPath, nil
}

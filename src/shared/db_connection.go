package shared

import (
	"GiftWize/src/entity/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Env = GetEnvs()

func Init() *gorm.DB {
	var db *gorm.DB
	dsn := "host=" + Env["DB_HOST"] + " user=" + Env["DB_USER"] + " password=" + Env["DB_PASSWORD"] + " dbname=" + Env["DB_NAME"] + " port=" + Env["DB_PORT"] + " sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	migration(db)
	log.Println("Database connected")
	return db
}

func migration(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Customer{},
		&models.Order{},
		&models.Company{},
		&models.CompanyOrder{},
		&models.API{},
		&models.AuditLog{},
		&models.Campaign{},
		&models.GiftCard{},
		&models.Inventory{},
		&models.Report{},
		&models.Setting{},
		&models.Transaction{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Database migrated")

}

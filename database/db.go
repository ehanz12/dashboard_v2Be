package database

import (
	"be_dashboard/config"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	cfg := config.AppConfig

	//lakukan koneksi ke database dengan config tersebut
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get underlying SQL DB: " + err.Error())
	}

	// SetMaxIdleConns: jumlah maksimal koneksi yang dibiarkan idle di pool
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns: jumlah maksimal koneksi yang terbuka ke database
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime: durasi maksimal sebuah koneksi dapat digunakan kembali
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	fmt.Println("Database connected successfully! 👌")
}

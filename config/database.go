package config

import (
	"fmt"
	"homework_mitramas/model"
	"log"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	var err error
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	seeder, err := strconv.ParseBool(os.Getenv("SEEDER"))
	if err != nil {
		panic("Seeder status must be filled")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost, dbPort, username, password, dbname)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect database")
	}

	sqlDB, _ := DB.DB()
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Database not responding with err: %s", err)
	}
	DB.AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.Member{},
	)
	if seeder {
		SeederRole()
		SeederAdmin()
	}
	log.Println("Success connected database")
	return DB
}

func SeederRole() {
	role := []model.Role{
		{
			RoleName: "Admin",
		},
		{
			RoleName: "Client",
		},
		{
			RoleName: "Member",
		},
	}
	DB.Create(&role)
}

func SeederAdmin() {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 4)
	user := model.User{
		UserName: "Admin",
		Email:    "admin@gmail.com",
		RoleId:   1,
		Password: string(hashPassword),
	}
	DB.Create(&user)
}

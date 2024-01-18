package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dba *gorm.DB

func GetCommonDBInstance() (db *gorm.DB) {

	// @TODO: Read from config

	user := "postgres"
	password := "1505"

	host := "localhost"
	port := "5432"
	dbname := "postgres"
	schema := "did"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=%s", host, user, password, dbname, port, schema)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	dba = db
	if err != nil {
		log.Panic().Msgf("Error connecting to the database at %s:%s/%s", host, port, dbname)
	}
	sqlDB, err := dba.DB()
	if err != nil {
		log.Panic().Msgf("Error getting GORM DB definition")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	log.Info().Msgf("Successfully established connection to %s:%s/%s", host, port, dbname)

	return dba

}

type UserTest struct {
	ID       string `gorm:"column:id"`
	Username string `gorm:"column:user_name"`
	Password string `gorm:"column:pass_word"`
}

func (UserTest) TableName() string {
	return "did.user_test"
}

func main() {
	storage := GetCommonDBInstance()
	var Usertest1 UserTest
	ID := uuid.NewString()
	Usertest1.ID = ID
	Username := "Sharan"
	Usertest1.Username = Username
	Password := "1505"
	Usertest1.Password = Password

	err := storage.Create(&Usertest1).Error
	if err != nil {
		fmt.Println("Error in Inserting", err)
	} else {
		fmt.Println("Record inserted successfully")
	}
}

package config

import (
	// "database/sql"
	"fmt"
	"os"

	// "webserver/config"
	"webserver/model"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// db *sql.DB
	db  *gorm.DB
	err error
)

// func Connect()(*sql.DB, error){
func Connect() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	psqlInfo := fmt.Sprintf(
		`host=%s 
		port=%d 
		user=%s
		password=%s 
		dbname=%s 
		sslmode=disable`,
		host, port, user, password, dbname,
	)

	// db, err := sql.Open("postgres", psqlInfo)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// err = db.Ping()
	// if err != nil{
	// 	panic(err)
	// }
	fmt.Println("successfully connected to database")
	// return db,nil
	db.Debug().AutoMigrate(&model.Employee{}, &model.Item{})
}

func GetDB() *gorm.DB {
	return db
}

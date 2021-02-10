package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"test/config"
)

var DB *sqlx.DB

func ConnectDB() {
	conf := config.New()
	var err error
	databaseUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DbHost, conf.DbPort, conf.DbUser, conf.DbPassword, conf.DbDatabase)

	DB, err = sqlx.Connect("postgres", databaseUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to PostgreSQL was established")
}

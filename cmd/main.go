package main

import (
	"database/sql"
	"log"

	"github.com/ayushn2/go_ecom.git/cmd/api"
	"github.com/ayushn2/go_ecom.git/config"
	db "github.com/ayushn2/go_ecom.git/db"
	"github.com/go-sql-driver/mysql"
)

func main(){

	db, err := db.NewMySQLStorage(mysql.Config{
		User : config.Envs.DBUser,
		Passwd : config.Envs.DBPassword,
		Addr : config.Envs.DBAddress,
		DBName : config.Envs.DBName,
		Net : "tcp",
		AllowNativePasswords : true,
		ParseTime : true,
	})

	if err != nil{
		log.Panic(err)
	}

	initStorage(db)

	server := 	api.NewAPIServer(":8080",db)
	if err := server.Run() ; err != nil{
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB){
	err := db.Ping()
	if err != nil{
		log.Fatal(err)
	}

	log.Println("Database successfully connected!!")
}
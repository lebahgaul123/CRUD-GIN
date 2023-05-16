package main

import (
	"crud-gin/controllers"
	"crud-gin/database"
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = 1234
// 	dbname   = "person_db"
// )

var (
	DB  *sql.DB
	err error
)

func main() {

	// ENV Config
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file enviroment")
	} else {
		fmt.Println("Success read file enviroment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	fmt.Println(psqlInfo)
	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Coneection FAILED")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	// router gin

	r := gin.Default()
	r.GET("/persons", controllers.GetAllPerson)
	r.POST("/persons", controllers.InsertPerson)
	r.PUT("/persons/:id", controllers.UpdatePerson)
	r.DELETE("/persons/:id", controllers.DeletePerson)

	r.Run("localhost:8080")

}

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

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DATABASE"))
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
	r.GET("/", controllers.Cekaja)
	r.GET("/persons", controllers.GetAllPerson)
	r.POST("/persons", controllers.InsertPerson)
	r.PUT("/persons/:id", controllers.UpdatePerson)
	r.DELETE("/persons/:id", controllers.DeletePerson)

	r.Run(os.Getenv("PG_PORT"))
	fmt.Println(r.Run())

}

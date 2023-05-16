package main

import (
	"crud-gin/controllers"
	"crud-gin/database"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {

	// ENV Config
	// err = godotenv.Load("config/.env")
	// if err != nil {
	// 	fmt.Println("failed load file enviroment")
	// } else {
	// 	fmt.Println("Success read file enviroment")
	// }

	psqlInfo := "PxzgJwJV60QSDHMF00bl psql -h containers-us-west-73.railway.app -U postgres -p 7511 -d railway"
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

	r.Run("")

}

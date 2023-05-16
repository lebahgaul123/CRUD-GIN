package controllers

import (
	"crud-gin/database"
	"crud-gin/repository"
	"crud-gin/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Cekaja(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": true,
		"data":    "COBAAPI",
	})
}

func GetAllPerson(c *gin.Context) {
	var (
		result gin.H
	)

	persons, err := repository.GetAllperson(*database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": "berhasil mendapatkan data",
			"data":   err,
		}
	} else {
		result = gin.H{
			"result": persons,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertPerson(c *gin.Context) {
	var person structs.Person

	err := c.ShouldBindJSON(&person)

	if err != nil {
		panic(err)
	}

	err = repository.InsertPerson(database.DbConnection, person)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert person",
		"data":   person,
	})
}

func UpdatePerson(c *gin.Context) {
	var person structs.Person

	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&person)

	if err != nil {
		panic(err)
	}

	person.ID = int64(id)

	err = repository.UpdatePerson(database.DbConnection, person)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update person",
		"data":   person,
	})
}

func DeletePerson(c *gin.Context) {
	var person structs.Person
	id, err := strconv.Atoi(c.Param("id"))

	person.ID = int64(id)

	err = repository.DeletePerson(database.DbConnection, person)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Person",
	})
}

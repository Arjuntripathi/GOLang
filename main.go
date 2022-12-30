package main

//inport for use router
import (
	"GOLang/StudentInfo"

	"github.com/gin-gonic/gin"
)

//main function which perform every action
//this for rest api

func main() {
	router := gin.Default()
	router.GET("/Student", StudentInfo.GetStudent)
	router.POST("/Student", StudentInfo.PostStudent)
	router.GET("/Student/:id", StudentInfo.GetAStudent)

	router.PUT("/Student/:id", StudentInfo.UpdateAStudent)
	router.DELETE("/Student/:id", StudentInfo.DeleteAStudent)

	router.Run("0.0.0.0:8080")
}

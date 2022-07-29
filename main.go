package main

//inport for use router
import (
	"github.com/gin-gonic/gin"
	"/usr/local/go/src/GOLang/StudentInfo"
	)


//main function which perform every action

func main() {
	router := gin.Default()
	router.GET("/Student", StudentInfo.GetStudent)
	router.POST("/Student", StudentInfo.PostStudent)
	router.GET("/Student/:id", StudentInfo.GetAStudent)
	
	router.PUT("/Student/:id", StudentInfo.UpdateAStudent)
	router.DELETE("/Student/:id", StudentInfo.DeleteAStudent)

	router.Run("localhost:8080")
}

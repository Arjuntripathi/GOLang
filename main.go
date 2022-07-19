package main

//inport for use router
import "github.com/gin-gonic/gin"


//main function which perform every action

func main() {
	router := gin.Default()
	router.GET("/Student", getStudent)
	router.POST("/Student", postStudent)
	router.GET("/Student/:id", getAStudent)
	
	router.PUT("/Student/:id", updateAStudent)
	router.DELETE("/Student/:id", DeleteAStudent)

	router.Run("localhost:3010")
}

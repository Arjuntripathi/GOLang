package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Students struct
type Student struct {
	Name   string  `json: "name"`
	ID     string  `json: "id"`
	Course string  `json: "course"`
	GPA    float64 `json: "gpa"`
}

//students info filled
var Students = []Student{
	{Name: "Arjun Tripathi", ID: "10", Course: "BCA", GPA: 7.4},
	{Name: "Rahul Durgapal", ID: "15", Course: "BCA", GPA: 6.9},
	{Name: "Kumar Aniket", ID: "24", Course: "BCA", GPA: 7.1},
	{Name: "Pratap Maurya", ID: "22", Course: "BCA", GPA: 6.8},
	{Name: "Mithilesh Singh", ID: "14", Course: "BBA", GPA: 6.9},
	{Name: "Manibhadra Singh", ID: "28", Course: "BCA", GPA: 7.2},
	{Name: "Saurabh Singh", ID: "21", Course: "BCA", GPA: 7.6},
}

//function to get all students info
/*
	curl http://localhost:8080/Student/' \
		--include \
		--header "Content-Type: application/json" \
		--request "GET" 

	______________EXAMPLE___________________

	curl http://localhost:8080/Student/ \
		--include \
		--header "Content-Type: application/json" \
		--request "GET"
*/
func getStudent(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Students)
}



//function to insert a student info
//using a commmand on terminal
/*
	curl http://localhost:8080/Student/ \
		--include \
		--header "Content-Type: application/json" \
		--request "POST" \
		--data '{what date you want to insert}'

	_________________Example________________
	
	curl http://localhost:8080/Student/ \
		--include \
		--header "Content-Type: application/json" \
		--request "POST" \
		--data '{"name": "TRISHA", "id": "4", "course": "BCA", "gpa": 9.5}'
*/
func postStudent(c *gin.Context) {
	var newStudent Student

	if err := c.BindJSON(&newStudent); err != nil {
		return
	}
	Students = append(Students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}



//function to get info of particular student by givine there id
/*
	curl http://localhost:8080/Student/'id' \
		--include \
		--header "Content-Type: application/json" \
		--request "GET" 

	______________EXAMPLE________________
	
	curl http://localhost:8080/Student/14 \
		--include \
		--header "Content-Type: application/json" \
		--request "GET"
*/
func getAStudent(c *gin.Context) {
	var id = c.Param("id")
	for _, Student := range Students {
		if Student.ID == id {
			c.IndentedJSON(http.StatusOK, Student)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"mesage": "Student are not avilable"})
}


//function to update particular student info by giving there id
//using a command on terminal
/*
	url http://localhost:3010/Student/14 \
		--include \
		--header "Content-Type: application/json" \
		--request "PUT" \
		--data '{write what to update}'

	____________EXAMPLE___________

	url http://localhost:3010/Student/14 \
		--include \
		--header "Content-Type: application/json" \
		--request "PUT" \
		--data '{"name": "TRISHA", "id": "4", "course": "BCA", "gpa": 9.5}'
*/
func updateAStudent(c *gin.Context){
	var id = c.Param("id")
	var tempStudent Student
	if err := c.BindJSON(&tempStudent); err != nil {
		return
	}
	Students = append(Students, tempStudent)
	for i, Student := range Students {
		if Student.ID == id {
			Students[i] = tempStudent;
			c.IndentedJSON(http.StatusNotFound, gin.H{"mesage": "Student info updated"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"mesage": "Student are not avilable"})
}


//function to deleted a particular student 
//using command on your terminal
/*
	curl http://localhost:3010/Student/'id' \
		--include \
		--header "Content-Type: application/json" \
		--request "DELETE"

	-------------------EXAMPLE--------------

	curl http://localhost:3010/Student/4 \
		--include \
		--header "Content-Type: application/json" \
		--request "DELETE"
*/
func DeleteAStudent(c *gin.Context){
	var id  = c.Param("id")
	for i, Student := range Students {
		if Student.ID == id {
			Students = RemoveIndex(Students,i)
			c.IndentedJSON(http.StatusNotFound, gin.H{"mesage": "Student info Deleted"}) //user define function
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"mesage": "Student are not avilable"})
	
}

//A function Which is created to delete a particular index and return new array
func RemoveIndex(tempStudents []Student, index int) []Student{
	return append(tempStudents[:index], tempStudents[index+1:]...)
}

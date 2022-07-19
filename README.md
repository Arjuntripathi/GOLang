# GOLang
Program based on Go language and using GIN api
<h3>
Design API endpoints
</h3>
When developing an API, you typically begin by designing the endpoints. Your API’s users will have more success if the endpoints are easy to understand.
<ul>
<li>GET – Get a list of all student, returned as JSON.</li>
<li>POST – Add a new student from request data sent as JSON./albums/:id</li>
<li>GET – Get an student by its ID, returning the album data as JSON.</li>
<li>PUT - put a update a of a student</li>
<li>DELETE - delete a student from student</li>
</ul>

<h4>Step 1:</h4>
<ul>
create a folder for your code
<pre>
$mkdir %temp_file%
$cd %temp_file%
</pre>
run a go mod init command
<pre>
$go mod init %temp_file%
</pre>
create a <b>main.go</b> file 
</ul>

<h4>Strp 2:</h4>
<ul>
Into main.go, at the top of the file, paste the following package declaration.
<pre>
package main

//inport for use router
import "github.com/gin-gonic/gin"
</pre>
it cause error, for solving it go on terminal and run it
<pre>
$ go get -u github.com/gin-gonic/gin
</pre>
and create a main function
<pre>
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
</pre>
</ul>

<h4>Step 3:</h4>
<ul>
create a new file <b>studentinfo.go</b> in folder and define method
<pre>
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
</pre>
<br><br>
<b>write a handeler to handle a api method</b><br><br><br>
this is for <b>GET</b> method
<pre>
//function to get all students info
func getStudent(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Students)
}
</pre>

this is for <b>POST</b> method
<pre>
//function to insert a student info
func postStudent(c *gin.Context) {
	var newStudent Student

	if err := c.BindJSON(&newStudent); err != nil {
		return
	}
	Students = append(Students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}
</pre>

this is for <b>GET/:id</b> method
<pre>
//function to get info of particular student by givine there id
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
</pre>

this is for <b>PUT</b> method
<pre>
//function to update particular student info by giving there id
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

</pre>

this is for <b>DELETE</b> method
<pre>
//function to deleted a particular student 
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
</pre>

</ul>

<h4>Step 4:</h4>
<ul>
Run your Program go on terminal and type
	<pre>
	$ go run main.go studentinfo.go
	</pre>
	and then open another terminal on your pc and run as <b>step 5</b>
</ul>

<h4>Step 5:</h4>
<ul>
<b>how to run on terminal</b>
<br>
install <pre>curl</pre> to run it, then on terminal<br>
for <b>get()</b> method
<pre>
$ curl http://localhost:8080/Student/ \
		--include \
		--header "Content-Type: application/json" \
		--request "GET"
</pre>

for <b>post()</b> method
<pre>
$ curl http://localhost:8080/Student/ \
		--include \
		--header "Content-Type: application/json" \
		--request "POST" \
		--data '{"name": "TRISHA", "id": "4", "course": "BCA", "gpa": 9.5}'
</pre>

for <b>get()/:id</b> method
<pre>
$ curl http://localhost:8080/Student/14 \
		--include \
		--header "Content-Type: application/json" \
		--request "GET"
</pre>

for <b>put()</b> method
<pre>
$ curl http://localhost:3010/Student/14 \
		--include \
		--header "Content-Type: application/json" \
		--request "PUT" \
		--data '{"name": "TRISHA", "id": "4", "course": "BCA", "gpa": 9.5}'
</pre>

for <b>delete()</b> method
<pre>
$ curl http://localhost:3010/Student/4 \
		--include \
		--header "Content-Type: application/json" \
		--request "DELETE"
</pre>

</ul>

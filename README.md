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
	download a gid repo in your dirctory
	<pre>
	$ git clone https://github.coom/Arjuntripathi/GOLang
	</pre>
</ul>
<h4>Step 2:</h4>
<ul>
	Create a new mod and sum file 
	<pre>
	$ go mod init GOLang
	$ go mod tidy
	<pre>
</ul>
<h4>Step 3:</h4>
<ul>
	Create your Docker image file
	<pre>
	$ docker build -t golang .
	<pre>
	it will take some time...
</ul>
<h4>Step 4:</h4>
<ul>
	Run your Docker file
	<pre>
	$ docker run --publish 8080:8080 golang
	</pre>
	Open a new terminal to chech it work or not..
</ul>
<h4>Step 5:</h4>
<ul>
<b>how to run on terminal</b>
<br>
install <pre>$ curl</pre> to run it, then on terminal<br>
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

# wsMessagesExampleGo
This is un example to implement a web service for sending a plain texte and storing it on any receptacle (at you to implement a database model).

Some basic refactoring of the main package has given 5 files :

1. main.go <br>
The name talks by itself. It used to launch a server on http://localhost:8080/.

2. message.go <br>
This is the model of the message

3. messageHandlers.go <br>
This file gathers all handlers that we can use

4. router.go <br>
Allows to catch the URL of all endpoints (URI) and to select the right route to follow for calling the appropriate function. This is a dispatcher or a router of the http URL.

5. routes.go <br>
Allows to define each route with their name, method and URI pattern.

6. mockDB.go <br>
Used for mockup the database with a starter datas set and to store incrementally the new message. It didn't used to persist the datas after each reboot of the server.      

# How to test (after installing the wsMessagesExampeGo package and the server launched)
1. To test mockup datas <br>
curl http://localhost:8080/messages/test

2. To list all messages stored <br>
curl http://localhost:8080/messages

3. To get the message from his ID <br>
curl http://localhost:8080/messages/{msgID}

4. To store new message <br>
curl http://localhost:8080/messages/ -d "my name is Mehdi and I like Golang"



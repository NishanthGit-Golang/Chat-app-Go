This is an ChatApp Application  for Multiple client joining in chatroom, sending messages in chatroom, retrieving message , leaving or exiting from chatroom .

Running commands  : 
    go mod tidy 
    go build
   ./chatApp for linux or ./chatApp.exe for windows

# RestAPI : 
1.JOIN-client:GET
url : http://localhost:8080/chatApp/join?id=user1
response:{
    "message": "Client joined"
}
2.sendMessage : GET
url : http://localhost:8080/chatApp/send?id=user1&message=Hello, World1!
response :{
    "message": "Message sent"
}
3.retrieve Message : GET
url :http://localhost:8080/chatApp/getMessage?id=user1
response : {
    "message": "user1: Hello, World1!"
}

4.Client-Leave :GET
 url : http://localhost:8080/chatApp/leave?id=user1
response : 
{
    "message": "Client left"
}




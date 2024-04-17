# golang_crud

Project API: 

URL: http://localhost:8080/api/data
Method: Get
Response:

{
    "message": "Hello, World!"
}


**Get Data from DB**
URL: http://localhost:8080/api/users
Method: Get
Response:
[
    {
        "id": 1,
        "name": "88828828",
        "email": "yosecond@gmai.com"
    },
    {
        "id": 2,
        "name": "hello world",
        "email": "hello@example.com"
    }
]

**Create new user**
URL:http://localhost:8080/api/users/create
Method: POST
Body Request:
{
    "name":"Hello World",
   "email": "helloworld@gmai.com"
}
Response:
{
    "message": "User created successfully"
}

**Update user**
URL: http://localhost:8080/api/users/update
Method: POST
Body Request:
{
    "name":"John Deo",
    "email": "johdeo@gmai.com"
}
Body Response:
{
    "message": "User updated successfully. Affected row: 1"
}

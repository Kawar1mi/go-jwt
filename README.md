# JWT Authentication

Authentication system by using Gorm, Gin, bcrypt, and jwt libraries.
Also using live reload tool - Air

# Implemented methods

## Sign up

### Request

`POST /signup`

    curl --location 'localhost:8080/signup' \
    --header 'Content-Type: application/json' \
    --data '{
        "email": "nomail@email.com",
        "password": "simplepass"
    }'

### Response

    HTTP/1.1 200 OK
    Content-Type: application/json; charset=utf-8
    Date: Fri, 24 Nov 2023 16:19:48 GMT
    Content-Length: 2
    
    {}
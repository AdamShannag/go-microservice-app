# A simple microservice app in Go


## How to run
from the project directory, run the folowing
`make up_build` to start the services or `make down` to stop the services, and wait for the services to start.

## Services:
* Broker Service
* User Service

## Access the user service through the broker:
1. [POST] `http://localhost:8080/handle`
1. [GET] `http://localhost:8080/handle`
1. [GET] `http://localhost:8080/handle?id=123`
1. [PUT] `http://localhost:8080/handle`
1. [DELETE] `http://localhost:8080/handle?id=123`

> Note: Always provide a body in json that contains a property named action, set it as user, to access the user service, see the following:

### **Example 1:** Create a user in the user service.
* [POST] `http://localhost:8080/handle`
* Body:
```
{
    "action": "user",
    "user": {
        "address": "333 Jordan",
        "name": "Yaseen"
    }
}
```

### **Example 2:** Get a user from the user service.
* [GET] `http://localhost:8080/handle?id=123`
* Body:
```
{
    "action": "user",
}
```

### **Example 3:** Get all users from the user service.
* [GET] `http://localhost:8080/handle`
* Body:
```
{
    "action": "user",
}
```

### **Example 4:** Update a user from the user service.
* [PUT] `http://localhost:8080/handle`
* Body:
```
{
    "action": "user",
    "user": {
        "address": "New Address",
        "id": "39bf1",
        "name": "Yaseen"
    }
}
```

### **Example 5:** Delete a user from the user service.
* [DELETE] `http://localhost:8080/handle?id=123`
* Body:
```
{
    "action": "user",
}
```
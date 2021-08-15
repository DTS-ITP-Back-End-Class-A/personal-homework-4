# IT Perbankan Class-A Coding Exercise

The goal of these exercises are to practive your proficiency in web frameworks [Fiber] that is related to the daily job. Follow the instructions below to complete the Practice.

## Setup
1. clone this repository
2. create new branch with format `class-A/{your name}`
3. initialize your project by creating a folder and then running go mod init `workspace`
4. type `go get -u github.com/gofiber/fiber/v2`
5. copy template code in [Quickstart]
6. type `go run main.go`
7. Hit the Server to test Health `curl localhost:3000/` , expect a string `Hello, World 👋!` and `200` response
8. let's Rock !! 🚀

### Tasks 
We define routes for handling create and read operations:

| Method        | Route                 | Action                                              |
|---------------|-----------------------|-----------------------------------------------------|
| POST          | /create-car           | Add new cars                                        |
| GET           | /get-car              | Get all cars                                        |

Access API via ```http://localhost:3000/{route}```

### Test The APIs

1. POST ```/create-car```

Request Body: 
` choose type raw json `
```
{
    "car_name": "honda",
    "car_color": "red",
    "car_type": "matic"
}
```

Response:
set status code : 200
```
{
    "message": "success",
    "status": "ok",
    "data": {
        "car_name": "honda",
        "car_color": "red",
        "car_type": "matic"
    }
}
```

2. GET ```/get-car```

Authorization: ```Basic Auth```

if username and password not match : ``` set status code: 401 ```
```
{
    "message": "username/ password is incorrect",
}
```

if true (data hardcode):
```
{
    "message": "success",
    "status": "ok",
    "data": {
        "car_name": "honda",
        "car_color": "red",
        "car_type": "matic"
    }
}
```

#### Challenge
1.  POST ```/create-car```

store data in memory:
```
type car struct {
    carName     string  `json:"car_name"`
    carColor    string  `json:"car_color"`
    carType     string  `json:"car_type"`
}
const cars = []car

cars = append(cars, requestBody)

return
{
    "message": "success",
    "status": "ok",
    "data": {
        "car_name": request Body,
        "car_color": request Body,
        "car_type": request Body
    }
}
status code success 200
```

2. Get ```/create-car```

get data from memory:
```
if cars empty return 
{
    "message": "success",
    "status": "ok",
    "data": []
}
status code success 200

if cars length > 0 
{
    "message": "success",
    "status": "ok",
    "data": [{
        "car_name": request Body,
        "car_color": request Body,
        "car_type": request Body
    }]
}
status code success 200

```

[Fiber]: <https://github.com/gofiber/fiber>
[Quickstart]: <https://github.com/gofiber/fiber#%EF%B8%8F-quickstart>
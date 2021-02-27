# GO FIBER STARTER PACK (USING GORM)
## Dependency
Default DB : postgresql
1. [Go Fiber](https://github.com/gofiber/fiber)
2. [Go Playground Validator v10](https://github.com/go-playground/validator)
3. [Go Beans](https://github.com/hsjsjsj009/go-beans)
4. [logrus](https://github.com/Sirupsen/logrus)
5. [godotenv](https://github.com/joho/godotenv)
6. [go-playground universal translator](https://github.com/go-playground/universal-translator)
7. [jwx](https://github.com/lestrrat/go-jwx)
8. [jwt](https://github.com/dgrijalva/jwt-go)
9. [gorm](https://gorm.io)
10. and many more

## How to Run the Server
### Prerequisites
1. Run the postgresql db using docker-compose.dev.yml
```shell
cd server
make run-container
```
### Run Server
```shell
cd server
make run
```
### Stop Container
```shell
cd server
make stop-container
```

## How to Create a Server Docker Image
```shell
docker build -t=<your-image-name> -f ./server/Dockerfile .
```
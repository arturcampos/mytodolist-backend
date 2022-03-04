# My ToDoList api ![Go](https://img.shields.io/badge/Go-darkblue?style=plastic&logo=go) ![Project last updated](https://img.shields.io/github/last-commit/arturcampos/mytodolist-backend/main?label=Last%20updated&style=plastic)  

This project is built using the following languages, frameworks and tools:
- ![Go](https://img.shields.io/badge/Go-v1.17-darkgreen?style=plastic&logo=go) with [gin-gonic](https://github.com/gin-gonic/gin) framework that implements functionalities from net/http (native from GoLang).
- ![Docker](https://img.shields.io/badge/Docker-gray?style=plastic&logo=docker) for build and configure our application.
- ![Docker-Compose](https://img.shields.io/badge/Docker%20Compose-gray?style=plastic&logo=docker-compose) for running the envronment starting a database and the application making it avaliable to be used.
- ![MongoDB](https://img.shields.io/badge/MongoDB-latest-darkgreen?style=plastic&logo=mongodb) as a NoSQL database (tag version `latest` from DockerHub).

## Functionalities to be added 
- [ ] Unit tests
- [ ] Integration Tests
- [ ] Postman Documentation

### Project structure

``` 
Project
├── mytodolist-backend
|   ├── main.go: GO main file
|   ├── Dockerfile
|   ├── docker-compose.yml
|     ├── database
|       ├── db.go: GO file responsible to start up data base connection and get database and collections from mongo
|     ├── service
|       ├── task_service.go: Go file where all business rules related to tasks are contained
|   ├── go.mod: init project mod from GO
```


### How to build and run:
- ![TODO](https://img.shields.io/badge/TODO-darkred?style=plastic)

### Testing:
- ![TODO](https://img.shields.io/badge/TODO-darkred?style=plastic)

### Validating:
- ![TODO](https://img.shields.io/badge/TODO-darkred?style=plastic)

### Contact
[![Gmail Badge](https://img.shields.io/badge/-arturcampos13@gmail.com-c14438?style=flat&logo=Gmail&logoColor=white)](mailto:arturcampos13@gmail.com "Connect via Email")
[![Linkedin Badge](https://img.shields.io/badge/-arturcamposrodrigues-0072b1?style=flat&logo=Linkedin&logoColor=white)](https://www.linkedin.com/in/arturcamposrodrigues/?locale=en_US/ "Connect on LinkedIn")
[![Twitter Badge](https://img.shields.io/badge/-@_artur_campos-00acee?style=flat&logo=Twitter&logoColor=white)](https://twitter.com/intent/follow?screen_name=_artur_campos "Follow on Twitter")

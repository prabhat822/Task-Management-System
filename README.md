# Tasks Management System using REST API written in Go
## Table of contents
## Clean Architecture

### Architecture Layers
![architecture](assets/architecture.png)
### Folder Structure
```
tasks-go-clean-architecture
├── bootstrap
│   └── env.go
├── cmd
│   └── main.go
├── database
│   └── mysql
│       ├── migrations
│       │   ├── 000_create-database.sql
│       │   ├── 001_create-schema.sql
│       │   └── 002_seeds-data.sql
│       └── mysql.go
├── go.mod
├── go.sum
├── internal
│   ├── delivery
│   │   └── http
│   │       ├── handler       -> outermost layer driving the business logic
│   │       │   ├── handler.go
│   │       │   └── interface.go
│   │       ├── middleware
│   │       │   └── jwt_auth_middleware.go
│   │       └── route
│   │           └── route.go
│   ├── domain
│   │   ├── entity            -> entity encapsulate business rules (struct with methods)
│   │   │   ├── task.go
│   │   │   └── user.go
│   │   └── interface
│   │       ├── repository    -> repository interface bridge repository and usecase layer
│   │       │   ├── task.go
│   │       │   └── user.go
│   │       └── usecase       -> usecase interface bridge usecase and handler layer
│   │           ├── login.go
│   │           ├── refresh_token.go
│   │           ├── signup.go
│   │           └── task.go
│   ├── repository            -> implementation of repository layer
│   │   ├── taskrepo
│   │   │   └── repository.go
│   │   └── userrepo
│   │       ├── model.go
│   │       └── repository.go
│   └── usecase               -> implementation of usecase layer
│       ├── login
│       │   └── login_usecase.go
│       ├── refreshtkn
│       │   └── refresh_token_usecase.go
│       ├── signup
│       │   └── signup_usecase.go
│       └── tasks
│           └── task_usecase.go
├── pkg                        -> utilities, helpers, ...
│   └── tokenutil
│       └── tokenutil.go
└── README.md
```
### Domain
#### Entities
- User
- Task
#### Interfaces
- Usecase Interfaces
- Repository Interfaces
### Usecase
### Repository
## API

- POST `/api/signup`: Sign up
- POST `/api/login`: Login
- GET `/api/task`: Fetch all tasks of user
- POST `/api/task/add`: Add a task
- PUT `/api/task/{id}`: Toggle complete status of a task
- DELETE `/api/task/{id}`: Delete a task

## Set Up
## References

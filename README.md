# Go Monorepo Boilerplate

The right way to implement clean architecture on golang using DDD

# Description

The Clean Architecture is a software architecture proposed by Robert C. Martin (better known as Uncle Bob). In this repository, the contents are list of examples of implementation of Clean Architecture in Golang. The examples will using real world scenario.

To understand more about Clean Architecture: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

# Dependency Rules

Based on Uncle Bob, there are 4 layers:

- Entity
- Use Cases
- Interface Adapters
- Frameworks and Drivers.
In this repository, we also using 4 layers (with modification) like this:

- Entity
- Use Cases (Implementation)
- Interface Adapters. Will be splitted into two:
    - Repository Interface. Bridging repository implementation  and use cases layer.
    - Use Case Interface. Bridging handler and use cases layer.
- Driver Layer
    - Handler
    - Repository implementation

# Folder Structure

```
.github
   |-- workflows
   |   |-- go.yaml
.gitignore
README.md
app.config.yaml
cmd
   |-- README.md
database
   |-- mysql
   |   |-- init.go
   |-- postgres
   |   |-- init.go
   |-- redis
   |   |-- init.go
exceptions
   |-- default_err.go
go.mod
go.sum
helpers
   |-- response.go
   |-- string.go
   |-- time.go
lib
   |-- date
   |   |-- date.go
   |-- jsonb
   |   |-- jsonb.go
middleware
   |-- auth.go
pull_request_template.md
scripts
   |-- README.md
services
   |-- auth-svc
   |   |-- .gitignore
   |   |-- Dockerfile
   |   |-- app
   |   |   |-- repository
   |   |   |   |-- user.go
   |   |   |-- usecase
   |   |   |   |-- user
   |   |   |   |   |-- implement.go
   |   |   |   |   |-- interface.go
   |   |-- config
   |   |   |-- config.go
   |   |-- entity
   |   |   |-- user.go
   |   |-- handler
   |   |   |-- handler.go
   |   |   |-- interface.go
   |   |-- main.go
   |   |-- repository
   |   |   |-- user
   |   |   |   |-- model.go
   |   |   |   |-- repository.go
   |   |-- service.yaml
test
   |-- auth
   |   |-- login_test.go
   |-- gin_test.go
   ```

# Guidelines

Step-by-step writing code using this pattern
- Setup skeleton of the microservices (including: main.go, migrations, config, pkg and handler folder)
- Defining the entities
- Defining usecase (interface & implementation) in folder app. We're gonna focus in this folder since the business logic will be written here.
- When the usecase need to communicate to the external agency (Database, other apis, etc) then write it to the repository interface
- After the usecase layer was done, now time to write repository implementation
- Put it up together + register to the main.go and handler folder

# Future Example

[] gRPC Implementation
[] Rest API service with background worker
[] Rest API service with external dependency
[] Rest API service with event driven system (pubsub)
[] Add Health Check
[] Implement Kubernetes
# Pet Clinic Restful API 

## Goals
Build a Pet Clinic restful API.

## Go Version Manager
* Simple go version manager [g](https://github.com/stefanmaric/g)
* Go Version Manager [GVM](https://github.com/moovweb/gvm)
I chose the first one, simple version.


## Libraries
* Dependency management: [go mod](https://blog.golang.org/using-go-modules)
* Viper; [viper](https://github.com/spf13/viper)
* Routing framework: [gin gonic](https://github.com/gin-gonic/gin)
* Gin Swagger: [gin-swagger](https://github.com/swaggo/gin-swagger)
* Okta JWT: [okta](https://www.okta.com/)
* Prometheus Client: [prometheus-client](https://github.com/prometheus/client_golang/prometheus/promhttp)
* GORM [GORM](https://gorm.io/)
* Postgresql [postgresql](gorm.io/driver/postgres)
* Data validation: [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
* Logging: [zap](https://github.com/uber-go/zap)
* Configuration: [viper](https://github.com/spf13/viper)
* Testing & Mock: [testify](https://github.com/stretchr/testify)


## Secure Endpoints:
| Path     | Method | Description                                          |
|:---------|:-------|:-----------------------------------------------------|
| /health  | GET    | Healthcheck checks dependent services/http endpoints |
| /info    | GET    | Show the app info (nSyH8x3v)                         |
| /metrics | GET    | Get out of the box metrics                           |


## API Endpoints:
| Path          | Method | Description |
|:--------------|:-------|:------------|
| /v1/vets/:id  | GET    |             |
| /v1/pets/:id  | GET    |             |
| /v1/owner/:id | GET    |             |


---
**Installing mockery**<br/>
brew install vektra/tap/mockery<br/>
go get github.com/vektra/mockery/v3<br/>

**Creating mock function/interface**<br/>
mockery --case=underscore --name=OwnerRepositorier  --inpackage<br/>
mockery --case=underscore --name=OwnerServicer  --inpackage<br/>


Okta Authentication & Authorization

The Client Credentials flow is recommended for use in machine-to-machine authentication. Your application will need to securely store its Client ID and Secret and pass those to Okta in exchange for an access token. At a high-level, the flow only has two steps:

* Your application passes its client credentials to your Okta authorization server.
* If the credentials are accurate, Okta responds with an access token.



## Project Layout

The template project layout:

```
.
├── app              app configuration 
├── api              application API
│     └── health     health check endpoint
│     └── info       info endpoint
│     └── metrics    metrics endpoint
├── config            configuration files for different environments

├── docs                 swagger docs files
├── ds                   data server library
├── errors               error types and handling
├── pkg              shared libraries
└── testdata             test data scripts
```

## Application Architecture
Explain the design patterns, the libraries, and the frameworks used in the application.

### Entry Point
The server.go is the application entrypoint. It comprises two init functions. The first init function responsible 
for loading the application configuration, logger, and gin engine. The second init function is responsible for 
setting up the application routes and middleware.

### Configuration
The app.LoadConfig function is responsible for loading the application configuration in YAML format. It reads and 
converts the yaml file into a struct using the viper library.  For example, it read this yaml file:

```yaml
appInfo:
  name: "Pet Clinic"
  version: "1.0.0"
  description: "Pet Clinic Restful"

database:
  postgres:
    driver: postgres
    dsn: "user=postgres password=mysecretpassword host=localhost port=5432 dbname=petclinic sslmode=disable"
  maxIdleConns: 0
  maxOpenConns: 5
  maxIdleTime: 60
```
and converts it into this struct:
```go
type AppInfoConfig struct {
    Name        string `mapstructure:"name"`
    Version     string `mapstructure:"version"`
    Description string `mapstructure:"description"`
}

type DatabaseConfig struct {
    Postgres     PostgresConfig
    MaxIdleConns int
    MaxOpenConns int
    MaxIdleTime  int
}

type PostgresConfig struct {
    Driver string
    Dsn    string
}
```
**By convention, I add postfix Config to the struct name to indicate it's a configuration struct.**

### Components Instantiation

### Inverting Dependencies

### Middleware

### Router Register

### Unit Tests

### Swagger



### A

### A

### A

### A


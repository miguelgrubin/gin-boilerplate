# Gin Boilerplate

This repository contains a golang API using hexagonal architecture.

## Features

- Decoupled components (using Ports and Adapters)
- Handmade Swagger documentation
- Mocks generation using Mockery
- Full suite of tests using Testify

## Toolkit

**Included on dependencies**

- [Gin Gonic](https://github.com/gin-gonic/gin) Web framework
- [Testify](https://github.com/stretchr/testify) Assert + Mocks
- [GORM](https://github.com/jinzhu/gorm) SQL ORM
- [Cobra](https://github.com/spf13/cobra) Command line framework
- [Viper](https://github.com/spf13/viper) Config files toolkit

**External tools**

- [Revive](https://github.com/mgechev/revive) Some linters
- [Secure Go](https://github.com/securego/gosec) Security Checker
- [Mockery](https://github.com/vektra/mockery) Generate Mocks
- [Air](https://github.com/cosmtrek/air) Live Reload
- [Delve](https://github.com/go-delve/delve) Debug

## Endpoints

| Method | URL                       |
| ------ | :------------------------ |
| GET    | /health                   |
| POST   | /v1/pets                  |
| GET    | /v1/pets                  |
| GET    | /v1/pet/{petId}           |
| PATCH  | /v1/pet/{petId}           |
| DELETE | /v1/pet/{petId}           |
| GET    | /v1/tags                  |
| POST   | /v1/categories            |
| GET    | /v1/categories            |
| GET    | /v1/category/{categoryId} |
| DELETE | /v1/category/{categoryId} |
| POST   | /v1/orders                |
| GET    | /v1/orders                |
| GET    | /v1/order/{orderId}       |

## CLI Usage

- app
- app serve
- app migrate
- app seed
- app create config

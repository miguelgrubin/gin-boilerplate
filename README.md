# Gin Boilerplate

This repository contains basic implementation of:
- View <-> Service <-> Repository (three tier architecture)
- Swagger documentation
- ORM
- DB Migrations


Using:
- [Gin Framework](https://github.com/gin-gonic/gin)
- [Testify](https://github.com/stretchr/testify)
- [GORM](https://github.com/jinzhu/gorm)
- [Migrate](https://github.com/golang-migrate/migrate)
- [Cobra](https://github.com/spf13/cobra)
- [Viper](https://github.com/spf13/viper)


Basic endpoints:
| Method | URL                        |
| ------ |:-------------------------- |
| GET    | /docs                      |
| GET    | /swagger.json              |
| GET    | /health                    |
| GET    | /v1/pets                   |
| POST   | /v1/pets                   |
| GET    | /v1/pet/{petId}            |
| PATCH  | /v1/pet/{petId}            |
| DELETE | /v1/pet/{petId}            |
| GET    | /v1/tags                   |
| POST   | /v1/tags                   |
| GET    | /v1/tag/{tagId}            |
| DELETE | /v1/tag/{tagId}            |
| GET    | /v1/categories             |
| POST   | /v1/categories             |
| GET    | /v1/category/{categoryId}  |
| DELETE | /v1/category/{categoryId}  |
| POST   | /v1/store/order            |
| GET    | /v1/store/order            |
| DELETE | /v1/store/order            |
| POST   | /v1/store/confirm-order    |


Basic cli commands:
- app
- app serve
- app migrate
- app seed
- app config create


Desired integrates:
- Docker
- Docker Compose
- Kubernetes
- Prometheus
- confd -> etcd

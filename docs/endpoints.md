# Endpoints

| Method | URL                       | Module  |
| ------ | :------------------------ | :------ |
| GET    | /health                   | shared  |
| POST   | /v1/pets                  | petshop |
| GET    | /v1/pets                  | petshop |
| GET    | /v1/pet/{petId}           | petshop |
| PATCH  | /v1/pet/{petId}           | petshop |
| DELETE | /v1/pet/{petId}           | petshop |
| GET    | /v1/store/inventory       | petshop |
| POST   | /v1/store/order           | petshop |
| GET    | /v1/store/order/{orderId} | petshop |
| DELETE | /v1/store/order/{orderId} | petshop |
| POST   | /v1/user                  | users   |
| GET    | /v1/user/{username}       | users   |
| PUT    | /v1/user/{username}       | users   |
| DELETE | /v1/user/{username}       | users   |
| POST   | /v1/user/login            | users   |
| POST   | /v1/user/logout           | users   |

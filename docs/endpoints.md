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
| POST   | /v1/users                 | users   |
| GET    | /v1/user/{username}       | users   |
| PATCH  | /v1/user/{username}       | users   |
| DELETE | /v1/user/{username}       | users   |
| POST   | /v1/auth/login            | users   |
| POST   | /v1/auth/refresh          | users   |
| POST   | /v1/auth/logout           | users   |

# Gin Boilerplate

This repository contains a golang API using simplified clean architecture.
Using factory pattern, interfaces and dependency inversion.

## Features

- Handmade OpenAPI 3.0 documentation
- Full suite of tests using Testify
- App config using Viper
- Command lines with batteries included
- Linters, security scanners, live reload and debugger
- Authentication using JWT with refresh tokens
- Passwork hashing with Argon2
- Redis integration
- Integration testing with Testcontainers

## Toolkit

**Main Dependencies**

- [Gin Gonic](https://github.com/gin-gonic/gin) Gin is a high-performance HTTP web framework written in Go
- [Testify](https://github.com/stretchr/testify) A toolkit with common assertions and mocks that plays nicely with the standard library
- [GORM](https://github.com/go-gorm/gorm) The fantastic ORM library for Golang, aims to be developer friendly
- [Cobra](https://github.com/spf13/cobra) A Commander for modern Go CLI interactions
- [Viper](https://github.com/spf13/viper) Go configuration with fangs
- [Redis](https://github.com/redis/go-redis) Redis Go client
- [Testcontainers](https://github.com/testcontainers/testcontainers-go) Run Docker containers as a testing setup

**External Tools**

- [Revive](https://github.com/mgechev/revive) 🔥 ~6x faster, stricter, configurable, extensible, and beautiful drop-in replacement for golint
- [Secure Go](https://github.com/securego/gosec) Go security checker
- [Air](https://github.com/cosmtrek/air) Live Reload
- [Delve](https://github.com/go-delve/delve) Debugger
- [Bruno](https://github.com/usebruno/bruno) Opensource IDE For Exploring and Testing API's

## CLI Usage

| Commands              |                           |
| :-------------------- | :------------------------ |
| `./app`               | Help                      |
| `./app serve`         | Runs HTTP server          |
| `./app migrate`       | Applies GORM automigrate  |
| `./app seed`          | Populates DB              |
| `./app create-config` | Creates default config    |
| `./app generate-keys` | Generates random RSA keys |

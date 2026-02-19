# Documentation Index

This documentation covers the Gin Boilerplate project - a Go API using simplified clean architecture with factory pattern, interfaces, and dependency inversion.

## Table of Contents

1. [Endpoints](01-endpoints.md) - API routes and HTTP endpoints
2. [Dev Tools](02-dev-tools.md) - Development utilities (GoSec, Air, Delve)
3. [Tasks](03-tasks.md) - Improvement tasks and backlog
4. [Architecture](04-architecture.md) - Design patterns and project structure

## Project Overview

### Architecture

The project follows a **modular monolith** architecture with three main modules:

| Module         | Description                              |
| -------------- | ---------------------------------------- |
| `sharedmodule` | Common services (DB, Redis, JWT, Config) |
| `petshop`      | Pet management domain                    |
| `users`        | User management and authentication       |

Each module follows a layered structure:

```
pkg/<module>/
├── domain/       # Business entities and rules
├── usecases/     # Application business logic
├── repositories/ # Data persistence
├── server/       # HTTP handlers and DTOs
├── mocks/        # Test doubles
└── factory.go    # Dependency injection
```

### Tech Stack

- **Framework**: Gin Gonic
- **ORM**: GORM (Postgres/MySQL/SQLite)
- **Testing**: Testify + Testcontainers
- **CLI**: Cobra + Viper
- **Auth**: JWT with RSA + Argon2 hashing
- **Cache**: Redis

### Quick Start

```bash
# Create config file
./app create-config

# Generate RSA keys for JWT
./app generate-keys

# Run migrations
./app migrate

# Seed sample data
./app seed

# Start server
./app serve
```

### CLI Commands

| Command               | Description                 |
| --------------------- | --------------------------- |
| `./app serve`         | Runs HTTP server            |
| `./app migrate`       | Applies GORM automigrate    |
| `./app seed`          | Populates DB with test data |
| `./app create-config` | Creates default config file |
| `./app generate-keys` | Generates RSA key pair      |

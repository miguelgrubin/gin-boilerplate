# Architecture

This document describes the architectural decisions and patterns used in the Gin Boilerplate project.

## Overview

The project follows a **modular monolith** architecture combining principles from:

- Clean Architecture (dependency inversion, separation of concerns)
- Domain-Driven Design (rich domain entities, domain errors)
- Factory Pattern (manual dependency injection)

## Project Structure

```
gin-boilerplate/
├── cmd/                    # CLI commands (Cobra)
├── pkg/                    # Application code
│   ├── app.go              # Application root, module composition
│   ├── sharedmodule/       # Shared infrastructure services
│   ├── petshop/            # Pet management module
│   └── users/              # User & authentication module
├── api/                    # OpenAPI specifications
├── docs/                   # Documentation
├── scripts/                # Utility scripts
└── test/                   # Integration test utilities
```

## Module Structure

Each business module follows a consistent layered structure:

```
pkg/<module>/
├── domain/           # Core business logic
│   ├── <entity>.go   # Domain entities with behavior
│   └── errors.go     # Domain-specific errors
├── usecases/         # Application services
│   └── <entity>_usecases.go
├── repositories/     # Data access layer
│   ├── <entity>_repository.go
│   ├── <entity>_entity.go      # Persistence model
│   └── <entity>_entity_mapper.go
├── server/           # HTTP layer
│   ├── <entity>_handlers.go
│   ├── <entity>_dtos.go
│   └── error_handler.go
├── mocks/            # Test doubles
└── factory.go        # Module composition & DI
```

## Layer Responsibilities

### Domain Layer

The innermost layer containing business logic with no external dependencies.

**Entities** (`domain/<entity>.go`):
- Rich domain models with behavior
- Factory functions for creation (`CreatePet`, `CreateUser`)
- Hydration functions for persistence reconstruction (`HydratePet`, `HydrateUser`)
- Domain event tracking via `EventRegistry`
- Validation logic

```go
// Domain entity with encapsulated business logic
type Pet struct {
    ID            string
    Name          string
    Status        string
    eventRegistry *sd.EventRegistry
}

func CreatePet(payload CreatePetParams) Pet {
    pet := Pet{ID: uuid.New().String(), ...}
    pet.eventRegistry.AddEvent(PetCreated)
    return pet
}
```

**Domain Errors** (`domain/errors.go`):
- Custom error types for domain-specific failures
- Implement the `error` interface

```go
type PetNotFound struct {
    ID string
}

func (e *PetNotFound) Error() string {
    return fmt.Sprintf("pet not found: %s", e.ID)
}
```

### Use Cases Layer

Application-specific business rules orchestrating domain entities.

**Responsibilities**:
- Coordinate domain entities and repositories
- Implement application workflows (CRUD operations)
- Define use case interfaces for dependency injection

```go
type PetUseCasesInterface interface {
    Creator(PetCreatorParams) (domain.Pet, error)
    Finder(PetFinderParams) ([]domain.Pet, error)
    Shower(string) (domain.Pet, error)
    Updater(string, PetUpdatersParams) (domain.Pet, error)
    Deleter(string) error
}
```

### Repositories Layer

Data persistence abstraction.

**Components**:
- **Interface** (`<entity>_repository.go`): Contract for data access
- **Entity** (`<entity>_entity.go`): GORM persistence model
- **Mapper** (`<entity>_entity_mapper.go`): Domain <-> Persistence conversion

```go
type PetRepository interface {
    FindOne(string) (*domain.Pet, error)
    FindAll() ([]domain.Pet, error)
    Save(domain.Pet) error
    Delete(string) error
}
```

**Entity Separation**: Domain entities are distinct from persistence entities to:
- Keep domain logic independent of ORM annotations
- Allow different persistence strategies
- Enable domain model evolution without migration concerns

### Server Layer

HTTP interface handling.

**Components**:
- **Handlers** (`<entity>_handlers.go`): HTTP request/response handling
- **DTOs** (`<entity>_dtos.go`): Request/response data structures
- **Error Handler** (`error_handler.go`): Domain error to HTTP status mapping

```go
type PetHandlers struct {
    pu usecases.PetUseCasesInterface
}

func (h PetHandlers) SetupRoutes(r *gin.RouterGroup) {
    r.POST("/pets", h.Creator)
    r.GET("/pets", h.Finder)
    // ...
}
```

## Dependency Flow

```
HTTP Request
    │
    ▼
┌─────────────────┐
│  Server Layer   │  Handlers, DTOs
└────────┬────────┘
         │ depends on
         ▼
┌─────────────────┐
│ Use Cases Layer │  Application logic
└────────┬────────┘
         │ depends on
         ▼
┌─────────────────┐
│  Domain Layer   │  Entities, business rules
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ Repository Layer│  Data access (implements interfaces)
└─────────────────┘
```

**Key Principle**: Dependencies point inward. Outer layers depend on inner layers via interfaces.

## Shared Module

The `sharedmodule` provides cross-cutting infrastructure services:

| Service         | Interface             | Implementation        |
| --------------- | --------------------- | --------------------- |
| Configuration   | `ConfigService`       | Viper-based           |
| Logging         | `LoggerService`       | Zap logger            |
| Database        | `DBService`           | GORM (Postgres/MySQL) |
| JWT             | `JWTService`          | RSA-based tokens      |
| Password Hash   | `HashService`         | Argon2                |
| Cache           | `RedisService`        | go-redis              |
| RSA Keys        | `RSAService`          | Key file management   |

## Module Composition

The `factory.go` in each module handles dependency injection:

```go
func NewPetShopModule(s sharedmodule.SharedModuleServices) PetShopModule {
    db := s.DBService.GetDB()
    
    // Build dependency graph
    pr := repositories.NewPetRepository(db)
    pu := usecases.NewPetUseCases(pr)
    ph := server.NewPetHandlers(&pu)
    
    return PetShopModule{
        Repositories: PetShopModuleRepositories{Pet: pr},
        UseCases:     PetShopModuleUseCases{Pet: &pu},
        Handlers:     PetShopModuleHandlers{Pet: ph},
    }
}
```

**Application Root** (`pkg/app.go`):

```go
func NewApp() (*App, error) {
    sharedServices := sharedmodule.NewSharedModuleServices()
    petShopModule := petshop.NewPetShopModule(sharedServices)
    usersModule := users.NewUsersModule(sharedServices)
    
    return &App{
        SharedServices: sharedServices,
        PetShopModule:  petShopModule,
        UsersModule:    usersModule,
    }, nil
}
```

## Testing Strategy

### Unit Tests

- Domain entities tested in isolation (`domain/*_test.go`)
- Use cases tested with mock repositories (`usecases/*_test.go`)
- Handlers tested with mock use cases (`server/*_test.go`)

### Integration Tests

- Repository tests use Testcontainers for real database instances
- Located in `repositories/*_test.go`

### Mocks

- Located in `mocks/` directory per module
- Implement the same interfaces as production code

## Key Design Decisions

1. **Manual DI over frameworks**: Explicit wiring provides clarity and compile-time safety
2. **Domain/Persistence entity separation**: Keeps domain pure, allows schema evolution
3. **Rich domain entities**: Business logic lives in domain, not use cases
4. **Interface-based dependencies**: Enables testing and future flexibility
5. **Module-based organization**: Clear boundaries, potential for future service extraction

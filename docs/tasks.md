# Improvement Tasks

This document outlines the tasks needed to improve the codebase quality based on the analysis of Clean Code, Clean Architecture, and Self-Similarity.

---

## Priority Legend

- **P0**: Critical - Must fix (bugs, typos affecting functionality)
- **P1**: High - Should fix soon (architectural improvements)
- **P2**: Medium - Nice to have (code quality improvements)
- **P3**: Low - Future enhancements

---

## 1. Clean Code Improvements

### 1.1 Fix Typos and Naming Inconsistencies (P0)

- [x] **Fix `Showher` typo in pet usecases**
  - File: `pkg/petshop/usecases/pet_usecases.go:52`
  - Change: `Showher` → `Shower`
  - Also update interface at line 30
  - Update all callers in `pkg/petshop/server/pet_handlers.go:52`
  - Update mocks if any

- [x] **Fix `PetReponseFromDomain` typo**
  - File: `pkg/petshop/server/pet_handlers.go:37,58,76`
  - Change: `PetReponseFromDomain` → `PetResponseFromDomain`
  - Also fix in `pkg/petshop/server/pet_dtos.go` (function definition)

- [x] **Fix wrong package comment in users server**
  - File: `pkg/users/server/error_handler.go:1`
  - Change: "petshop endpoints" → "users endpoints"

### 1.2 Add Constants for Magic Strings (P1)

- [x] **Create status constants in user domain**
  - File: `pkg/users/domain/user.go`
  - Add constants:
    ```go
    const (
        StatusActive   = "active"
        StatusInactive = "inactive"
        RoleUser       = "user"
        RoleAdmin      = "admin"
    )
    ```
  - Update `CreateUser` function to use constants

- [x] **Create status constants in pet domain**
  - File: `pkg/petshop/domain/pet.go`
  - Add constants for pet statuses (available, pending, sold, etc.)

### 1.3 Remove Debug Logging from Business Logic (P1)

- [x] **Remove log.Println from user usecases**
  - File: `pkg/users/usecases/user_usecases.go:107,112`
  - Remove or replace with structured logger passed via DI

### 1.4 Add Domain Validation (P2)

- [x] **Add email validation to User domain**
  - File: `pkg/users/domain/user.go`
  - Create `ValidateEmail(email string) error` function
  - Add validation error types to `pkg/users/domain/errors.go`

- [x] **Add username validation**
  - Min/max length
  - Allowed characters
  - Create `InvalidUsername` error type

- [x] **Add phone validation**
  - Basic format validation
  - Create `InvalidPhone` error type

### 1.5 Fix Unused/Incomplete Code (P2)

- [x] **Review `NewUser()` and `NewPet()` functions**
  - These create entities without IDs
  - Made private and added `HydrateUser` / `HydratePet` functions for persistence hydration
  - Updated mappers and tests to use proper factory functions

### 1.6 Add Structured Logging (P2)

- [x] **Integrate structured logging library**
  - Chose: zap
  - Add logger to `SharedModuleServices`
  - Replace all `log.Println` calls
  - Add request logging middleware with correlation IDs

---

## 2. Clean Architecture Improvements

### 2.1 Separate Authentication Module (P1)

- [ ] **Create new `auth` module**
  - Structure:
    ```
    pkg/auth/
    ├── factory.go
    ├── domain/
    │   ├── credentials.go
    │   ├── tokens.go
    │   └── errors.go
    ├── usecases/
    │   └── auth_usecases.go
    └── server/
        ├── auth_handlers.go
        ├── auth_dtos.go
        └── error_handler.go
    ```

- [ ] **Move authentication logic from users module**
  - Move `LoggerIn`, `LoggerOut`, `RefreshToken` to auth usecases
  - Keep user CRUD in users module
  - Auth module depends on user repository (read-only)

- [ ] **Update routes**
  - Move `/auth/*` routes to auth module
  - Keep `/users/*` routes in users module

### 2.2 Add Ports/Interfaces Layer (P1)

- [ ] **Create ports package in domain**
  - File: `pkg/users/domain/ports.go`
  - Define repository interface in domain layer (not in repositories package)
  - This follows proper Clean Architecture dependency direction

- [ ] **Create ports for petshop**
  - File: `pkg/petshop/domain/ports.go`
  - Same pattern as users

### 2.3 Implement Domain Event Publishing (P2)

- [ ] **Create event bus interface**
  - File: `pkg/sharedmodule/services/event_bus.go`
  - Define `EventBus` interface with `Publish(event string, payload interface{})` method

- [ ] **Implement in-memory event bus**
  - For local development and testing
  - Support event handlers registration

- [ ] **Add event publishing to repositories**
  - After successful save, publish events from entity's event registry
  - Clear event registry after publishing

- [ ] **Future: Add external event bus implementation**
  - Redis Pub/Sub, RabbitMQ, or Kafka adapter

---

## 3. Self-Similarity Improvements

### 3.1 Standardize Mock File Names (P1)

- [x] **Rename petshop mock files to match users pattern**
  - `pkg/petshop/mocks/repositories.go` → `pkg/petshop/mocks/pet_repository.go`
  - `pkg/petshop/mocks/usecases.go` → `pkg/petshop/mocks/pet_usecases.go`
  - Applied entity-prefixed naming convention consistently


### 3.3 Standardize Error Handler Pattern (P2)

- [x] **Create shared error handler base**
  - File: `pkg/sharedmodule/server/error_handler.go`
  - Define common HTTP error mapping with `ErrorClassifier` pattern
  - Module error handlers extend/compose with base via `HandleErrorWithClassifier`
  - Added helper functions: `SendError`, `SendNotFound`, `SendUnauthorized`, `SendBadRequest`, `SendInternalError`

- [x] **Add error response structure**
  - Consistent JSON error response format:
    ```go
    type ErrorResponse struct {
        Code    string `json:"code"`
        Message string `json:"message"`
        Details any    `json:"details,omitempty"`
    }
    ```

### 3.4 Standardize DTO Patterns (P2)

- [x] **Create base DTO types**
  - File: `pkg/sharedmodule/server/dtos.go`
  - Common patterns for pagination (`PaginationRequest`, `PaginationResponse`)
  - Timestamps (`TimestampsResponse`)
  - Generic list wrapper (`ListResponse[T]`)
  - ID request (`IDRequest`)

- [ ] **Add DTO validation tags**
  - Use `binding:"required"` consistently
  - Add custom validators if needed

---

## 4. Testing Improvements

### 4.1 Increase Test Coverage (P2)

- [ ] **Add handler tests for petshop**
  - Currently users has more thorough handler tests
  - Match the same coverage level

- [ ] **Add integration tests for full flows**
  - User registration → login → access protected route → logout
  - CRUD flows with real database (testcontainers)

### 4.2 Add Test Helpers (P2)

- [ ] **Create test fixtures package**
  - File: `pkg/testutil/fixtures.go`
  - Factory functions for creating test entities
  - Reduce duplication in test files

- [ ] **Create test database helpers**
  - File: `pkg/testutil/database.go`
  - Common setup/teardown for integration tests


---

## 5. Documentation Improvements

### 5.1 Add Architecture Decision Records (P2)

- [ ] **Create ADR directory**
  - `docs/adr/`
  - Document key architectural decisions

- [ ] **Write initial ADRs**
  - ADR-001: Modular monolith structure
  - ADR-002: Manual dependency injection
  - ADR-003: Domain entity vs persistence entity separation

### 5.2 Improve Code Documentation (P3)

- [ ] **Add GoDoc comments to all exported types**
  - Interfaces should document expected behavior
  - Functions should document parameters and return values

- [ ] **Update README with architecture overview**
  - Add diagram
  - Explain module structure
  - Quick start guide

---

## Implementation Order

### Phase 1: Quick Wins (1-2 days)
1. Fix all typos (1.1)
2. Add constants (1.2)
3. Remove debug logging (1.3)
4. Standardize mock file names (3.1)

### Phase 2: Core Improvements (3-5 days)
1. Add domain validation (1.4)
3. Standardize error handlers (3.3)
4. Add test helpers (4.2)

### Phase 3: Architectural Enhancements (1-2 weeks)
1. Separate authentication module (2.1)
2. Add ports layer (2.2)
3. Implement event publishing (2.3)

### Phase 4: Future Enhancements
1. Structured logging (1.6)
3. ADRs and documentation (5.x)

---

## Verification Checklist

After completing tasks, verify:

- [ ] All tests pass: `make test`
- [ ] Linting passes: `make lint`
- [ ] Build succeeds: `make build`
- [ ] Code coverage maintained/improved
- [ ] No regressions in existing functionality

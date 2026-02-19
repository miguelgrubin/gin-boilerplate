# Improvement Tasks

This document outlines the tasks needed to improve the codebase quality based on the analysis of Clean Code, Clean Architecture and Self-Similarity.

---

## Priority Legend

- **P0**: Critical - Must fix (bugs, typos affecting functionality)
- **P1**: High - Should fix soon (architectural improvements)
- **P2**: Medium - Nice to have (code quality improvements)
- **P3**: Low - Future enhancements

---

## 1. Documentation Improvements

### 1.1 Add Architecture Decision Records (P2)

- [ ] **Create ADR directory**
  - `docs/adr/`
  - Document key architectural decisions

- [ ] **Write initial ADRs**
  - ADR-001: Modular monolith structure
  - ADR-002: Manual dependency injection
  - ADR-003: Domain entity vs persistence entity separation

---

## 2. API First Approach (P1)

### 2.1 Upgrade to OpenAPI 3.x Specification

- [ ] **Migrate from Swagger 2.0 to OpenAPI 3.0+**
  - Update `api/swagger.yaml` → `api/openapi.yaml`
  - Document all existing endpoints (pets, users, auth, health)
  - Define all request/response schemas

- [ ] **Document complete API specification**
  - `GET/POST /v1/pets`, `GET/PUT/DELETE /v1/pets/{id}`
  - `GET/POST /v1/users`, `GET/PUT/DELETE /v1/users/{id}`
  - `POST /v1/auth/login`, `POST /v1/auth/refresh`
  - `GET /health`

---

## 3. OpenCode Commands for Code Quality (P3)

### 3.1 Create OpenCode Configuration

- [ ] **Initialize OpenCode config directory**
  - Create `.opencode/` directory structure
  - Create `opencode.json` with base configuration
  - Create `commands/` and `agents/` subdirectories

### 3.2 Quality Commands

- [ ] **Create `/lint` command**
  - Run `make lint` and suggest fixes
  - Integrate with revive output

- [ ] **Create `/test` command**
  - Run tests with coverage report
  - Analyze failures and suggest fixes

- [ ] **Create `/coverage` command**
  - Run `./scripts/coverage.sh`
  - Identify uncovered code paths

### 3.3 Architecture Commands

- [ ] **Create `/check-arch` command**
  - Verify clean architecture patterns
  - Check dependency direction (outer → inner layers)
  - Validate module structure consistency

### 3.4 Review Commands

- [ ] **Create `/review` command**
  - Code review with focus on quality, security, performance
  - Use plan mode (no changes)

### 3.5 Scaffold Commands

- [ ] **Create `/new-module` command**
  - Generate new module structure following existing patterns
  - Create domain/, usecases/, repositories/, server/, mocks/ directories
  - Generate factory.go boilerplate

### 3.6 Security Commands

- [ ] **Create `/security` command**
  - Run `make sec` (gosec)
  - Check for OWASP vulnerabilities
  - Analyze and suggest remediations

### 3.7 Custom Agents

- [ ] **Create `arch-reviewer` agent**
  - Validate hexagonal/clean architecture patterns
  - Read-only mode (no file changes)
  - Focus on dependency inversion, layer separation

- [ ] **Create `security-auditor` agent**
  - Check OWASP Top 10 vulnerabilities
  - Read-only mode
  - Focus on input validation, auth, data exposure

---

## Verification Checklist

After completing tasks, verify:

- [ ] All tests pass: `make test`
- [ ] Linting passes: `make lint`
- [ ] Build succeeds: `make build`
- [ ] Code coverage maintained/improved
- [ ] No regressions in existing functionality

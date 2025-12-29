# Unit Tests

Comprehensive unit tests for the Order Service following Domain-Driven Design (DDD) and CQRS patterns.

## Test Organization

Tests are organized by DDD architectural layers:

```
tests/unit/
├── domain/                   # Domain Layer Tests
│   └── entity_test.go        # Domain entities (Base, Order, Orderitem)
│
├── application/              # Application Layer Tests
│   ├── command_test.go       # CQRS commands (Create, Update, Delete)
│   ├── queries_test.go       # CQRS queries (GetByID, GetAll, List, Search)
│   ├── handler_test.go       # Command & Query handlers
│   └── dto_test.go           # Data Transfer Object conversions
│
├── infrastructure/           # Infrastructure Layer Tests
│   ├── middleware_test.go    # HTTP middleware (Auth, RateLimit)
│   ├── config_test.go        # Configuration loading
│   └── http_handler_test.go  # HTTP endpoint handlers
│
├── pkg/                      # Shared Package Tests
│   ├── validator_test.go     # Request validation
│   └── response_test.go      # HTTP response helpers
│
└── telemetry/                # Observability Tests
    └── telemetry_test.go     # TelemetryFlow SDK integration
```

## Running Tests

### Run All Unit Tests

```bash
go test ./tests/unit/... -v
```

### Run Tests by Layer

```bash
# Domain layer
go test ./tests/unit/domain/... -v

# Application layer
go test ./tests/unit/application/... -v

# Infrastructure layer
go test ./tests/unit/infrastructure/... -v

# Package tests
go test ./tests/unit/pkg/... -v

# Telemetry tests
go test ./tests/unit/telemetry/... -v
```

### Run with Coverage

```bash
# Generate coverage report
go test ./tests/unit/... -cover -coverprofile=coverage.out

# View coverage in browser
go tool cover -html=coverage.out

# View coverage summary
go tool cover -func=coverage.out
```

### Run Benchmarks

```bash
# Run all benchmarks
go test ./tests/unit/... -bench=. -benchmem

# Run specific benchmark
go test ./tests/unit/domain/... -bench=BenchmarkNewOrder -benchmem
```

### Run Specific Test

```bash
# Run a specific test function
go test ./tests/unit/domain/... -run TestNewOrder -v

# Run tests matching a pattern
go test ./tests/unit/... -run ".*Validate.*" -v
```

## Test Patterns

All tests follow these conventions:

### Table-Driven Tests

```go
func TestExample(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"valid input", "test", "TEST"},
        {"empty input", "", ""},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Transform(tt.input)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

### Arrange-Act-Assert Structure

```go
func TestOrderCreate(t *testing.T) {
    // Arrange
    customerID := uuid.New()
    total := 100.0

    // Act
    order := entity.NewOrder(customerID, total, "pending")

    // Assert
    assert.NotNil(t, order)
    assert.Equal(t, customerID, order.CustomerID)
}
```

### Mock-Based Testing

```go
func TestHandler(t *testing.T) {
    // Setup mock
    mockRepo := new(MockOrderRepository)
    mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*entity.Order")).Return(nil)

    // Execute
    handler := NewHandler(mockRepo)
    err := handler.Create(ctx, cmd)

    // Verify
    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}
```

## Test Coverage by Component

| Layer          | Package                    | Coverage Focus                           |
|----------------|----------------------------|------------------------------------------|
| Domain         | `entity`                   | Entity creation, validation, soft delete |
| Application    | `command`                  | Command validation, entity conversion    |
| Application    | `query`                    | Query validation, pagination defaults    |
| Application    | `handler`                  | CQRS handler orchestration               |
| Application    | `dto`                      | Entity-to-DTO conversion                 |
| Infrastructure | `middleware`               | JWT auth, rate limiting, RBAC            |
| Infrastructure | `config`                   | Env var loading, defaults                |
| Infrastructure | `http/handler`             | HTTP request/response handling           |
| Pkg            | `validator`                | Struct tag validation                    |
| Pkg            | `response`                 | Standardized API responses               |
| Telemetry      | `telemetry`                | SDK initialization, graceful degradation |

## Dependencies

Tests use the following packages:

- `github.com/stretchr/testify/assert` - Assertions
- `github.com/stretchr/testify/require` - Fatal assertions
- `github.com/stretchr/testify/mock` - Mocking framework
- `github.com/labstack/echo/v4` - HTTP testing
- `net/http/httptest` - HTTP test utilities

## Best Practices

1. **Isolation**: Each test should be independent and not rely on external state
2. **Naming**: Use descriptive test names that explain the scenario
3. **Mocking**: Mock external dependencies to focus on unit under test
4. **Coverage**: Aim for >80% coverage on business logic
5. **Benchmarks**: Include benchmarks for performance-critical code
6. **Edge Cases**: Test boundary conditions and error scenarios

---

Generated by TelemetryFlow RESTful API Generator
Copyright (c) 2024-2026 DevOpsCorner Indonesia. All rights reserved.

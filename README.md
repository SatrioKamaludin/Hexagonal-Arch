CRUD-Go-Hexa-MongoDB
├─ go.mod
├─ go.sum
├─ internal
│  ├─ adapters
│  │  ├─ handlers
│  │  │  └─ product_handler.go
│  │  └─ repository
│  │     ├─ mongo
│  │     │  └─ profiling_repo.go
│  │     └─ postgresql
│  │        └─ product_repo.go
│  ├─ app
│  │  └─ app.go
│  ├─ domain
│  │  ├─ models
│  │  │  ├─ product.go
│  │  │  └─ profiling.go
│  │  └─ services
│  │     ├─ product_service.go
│  │     └─ profiling_service.go
│  ├─ ports
│  │  ├─ IProductRepository.go
│  │  ├─ IProfilingRepository.go
│  │  └─ IService.go
│  └─ utils
│     └─ response.go
├─ main.go
├─ pkg
│  └─ config
│     └─ config.go
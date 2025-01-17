myapp/
├── cmd/ # Entry point for the application
│ └── myapp/ # Main application folder
│ └── main.go # Main file
├── internal/ # Private application code (not exported)
│ ├── config/ # Configuration files and logic
│ │ └── config.go
│ ├── domain/ # Business logic and domain entities
│ │ └── user.go # Example domain entity
│ ├── repository/ # Data access layer (repositories)
│ │ └── user_repo.go # Example repository interface/implementation
│ ├── service/ # Application logic layer
│ │ └── user_service.go # Example service logic
│ ├── transport/ # API-related logic
│ │ ├── http/ # HTTP transport logic
│ │ │ ├── user_handler.go # Example HTTP handler
│ │ │ ├── middleware.go # HTTP middleware
│ │ │ └── router.go # HTTP router setup
│ │ └── grpc/ # gRPC transport logic (optional)
├── pkg/ # Shared or reusable packages (optional)
│ └── logger/ # Logging utility
├── migrations/ # Database migrations
├── docs/ # Documentation (Swagger, Postman, etc.)
├── go.mod # Dependency management
└── go.sum # Dependency lockfile

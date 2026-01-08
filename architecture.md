library-core-service/
│
├─ cmd/
│ └─ server/
│ └─ server.go # App bootstrap (config → DB → middleware → routes)
│
├─ internal/
│
│ ├─ model/ # 🔹 CENTRAL DOMAIN + ORM MODELS
│ │ ├─ entities.go
│ │ ├─ domain.go
│ │ └─ enums.go
│
│ ├─ validator/ # 🔹 CENTRAL REQUEST VALIDATION
│ │ ├─ validator.go
│ │ └─ errors.go
│
│ ├─ mailer/ # 🔹 EMAIL DELIVERY (HOW emails are sent)
│ │ ├─ mailer.go # Mailer interface
│ │ ├─ zeptomail.go # ZeptoMail implementation
│ │ └─ types.go # Email payload structs
│
│ ├─ renderer/ # 🔹 EMAIL RENDERING (WHAT emails look like)
│ │ ├─ renderer.go # html/template loader + executor
│ │ └─ html/
│ │ ├─ forgot_password.html
│ │ ├─ welcome.html
│ │ └─ verify_email.html
│
│ ├─ routes/ # 🔹 CENTRAL ROUTES REGISTRATION
│ │ └─ routes.go
│
│ ├─ auth/
│ │ ├─ handler/
│ │ │ ├─ handler.go
│ │ │ └─ routes.go
│ │ ├─ service/
│ │ │ └─ service.go # Calls renderer + mailer
│ │ └─ repository/
│ │ ├─ repository.go
│ │ └─ gorm.go
│
│ ├─ books/
│ │ ├─ handler/
│ │ ├─ service/
│ │ └─ repository/
│
│ ├─ users/
│ │ ├─ handler/
│ │ ├─ service/
│ │ └─ repository/
│
│ ├─ loans/
│ │ ├─ handler/
│ │ ├─ service/
│ │ └─ repository/
│
│ ├─ middleware/
│ │ ├─ auth.go
│ │ ├─ logging.go
│ │ └─ not_found.go
│
│ ├─ database/
│ │ └─ postgres.go
│
│ └─ config/
│ ├─ types.go
│ └─ config.go
│
├─ migrations/ # SQL migrations
│
├─ .env # Local environment variables
├─ .gitignore # Git ignore rules
├─ go.mod
├─ main.go
└─ README.md

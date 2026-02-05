library-core-service/
│
├─ cmd/
│ └─ server/
│ └─ server.go # App bootstrap (config → DB → middleware → routes)
│
├─ internal/
│
│ ├─ model/ # 🔹 CENTRAL DOMAIN + ORM MODELS
│ │ ├─ entities.go # GORM entities (User, Book, Order, Payment, etc)
│ │ ├─ domain.go # Request / response DTOs
│ │ └─ enums.go # Status enums (LoanStatus, OrderStatus, PaymentStatus)
│
│ ├─ utils/ # 🔹 SHARED HELPERS
│ │ ├─ time.go
│ │ ├─ pagination.go
│ │ ├─ strings.go
│ │ ├─ pointers.go
│ │ └─ ids.go
│
│ ├─ validator/
│ │ ├─ validator.go
│ │ └─ errors.go
│
│ ├─ mailer/ # 🔹 EMAIL DELIVERY
│ │ ├─ mailer.go
│ │ ├─ zeptomail.go
│ │ └─ types.go
│
│ ├─ renderer/ # 🔹 EMAIL TEMPLATES
│ │ ├─ renderer.go
│ │ └─ html/
│ │ ├─ forgot_password.html
│ │ ├─ welcome.html
│ │ └─ verify_email.html
│
│ ├─ routes/
│ │ └─ routes.go
│
│ ├─ auth/
│ │ ├─ handler/
│ │ ├─ service/
│ │ └─ repository/
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
│ ├─ orders/ # 🆕 ORDERS MODULE
│ │ ├─ handler/
│ │ ├─ service/
│ │ └─ repository/
│
│ ├─ payments/ # 🆕 PAYMENTS MODULE
│ │ ├─ handler/
│ │ ├─ service/
│ │ └─ repository/
│
│ ├─ webhooks/ # 🆕 PAYMENT WEBHOOK LISTENERS
│ │ └─ handler/
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
├─ migrations/
│
├─ .env
├─ .gitignore
├─ go.mod
├─ main.go
└─ README.md

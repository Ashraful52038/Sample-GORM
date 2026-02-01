Go Backend Project Part-1:PostgreSQL and GORM
📋 Overview

A Go backend application demonstrating database operations with PostgreSQL using GORM ORM. Features include database connection management, CRUD operations, request/response handling, and structured API responses.

🚀 Quick Start
Prerequisites

    Go 1.19+

    PostgreSQL 12+

    Git

Installation
bash

# Clone the repository
git clone <repository-url>
cd go-backend

# Initialize Go module
go mod init github.com/ashraful/go-backend

# Install dependencies
go get gorm.io/gorm
go get gorm.io/driver/postgres

# Set up PostgreSQL database
sudo -u postgres psql -c "CREATE DATABASE sampledb;"

Configuration

Update database connection in main.go:
go

dsn := "host=localhost user=postgres dbname=sampledb port=5432 sslmode=disable timezone=Asia/Dhaka"

Run the Application
bash

go run main.go

📦 Core Components
1. Database Connection (main.go)
go

func connectDB() *gorm.DB {
    dsn := "host=localhost user=postgres dbname=sampledb port=5432 sslmode=disable timezone=Asia/Dhaka"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    // ... error handling
    return db
}

2. Data Model (models/user.go)
go

type User struct {
    gorm.Model
    Name  string
    Email string
}

3. Request Structures

Product Request:
go

type ProductRequest struct {
    Name        string  `json:"name" validate:"required"`
    Price       float64 `json:"price" validate:"required,gt=0"`
    Stock       int     `json:"stock" validate:"gte=0"`
    Description string  `json:"description"`
}

User Request:
go

type UserRequest struct {
    Name  string `json:"name" validate:"required,min=2,max=100"`
    Email string `json:"email" validate:"required,email"`
}

4. API Response Format (response/api_response.go)
go

type APIResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

🔧 API Endpoints (Conceptual)
User Management

    POST /api/users - Create new user

    GET /api/users - List all users

    GET /api/users/:id - Get user by ID

    PUT /api/users/:id - Update user

    DELETE /api/users/:id - Delete user

Product Management

    POST /api/products - Create product

    GET /api/products - List products with filters

    GET /api/products/search - Search products

    PUT /api/products/:id - Update product

    DELETE /api/products/:id - Delete product

📊 Database Operations
Migration
go

db.AutoMigrate(&models.User{})

Create
go

user := models.User{Name: "John", Email: "john@example.com"}
db.Create(&user)

Read
go

var users []models.User
db.Find(&users)

Update
go

db.Model(&user).Update("Email", "new@example.com")

Delete
go

db.Delete(&user)

🛠️ Development
Dependencies
bash

# Install all dependencies
go mod tidy

# Check for updates
go list -u -m all

Testing
bash

# Run tests
go test ./...

# Test with coverage
go test -cover ./...

Building
bash

# Build for current OS
go build -o backend-app main.go

# Cross-compile for Linux
GOOS=linux GOARCH=amd64 go build -o backend-linux main.go

⚙️ Configuration
Environment Variables

Create .env file:
env

DB_HOST=localhost
DB_PORT=5432
DB_NAME=sampledb
DB_USER=postgres
DB_PASSWORD=your_password
DB_TIMEZONE=Asia/Dhaka

Database Configuration
go

func getDSN() string {
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    dbname := os.Getenv("DB_NAME")
    return fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable timezone=%s",
        host, user, dbname, port, timezone)
}

🔍 Code Examples
Complete CRUD Example
go

// Create
func createUser(db *gorm.DB, name, email string) error {
    user := models.User{Name: name, Email: email}
    result := db.Create(&user)
    return result.Error
}

// Read with conditions
func getUsersByEmail(db *gorm.DB, domain string) ([]models.User, error) {
    var users []models.User
    result := db.Where("email LIKE ?", "%@"+domain).Find(&users)
    return users, result.Error
}

Transaction Example
go

func createUserWithTransaction(db *gorm.DB, user models.User) error {
    return db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(&user).Error; err != nil {
            return err
        }
        // Additional operations...
        return nil
    })
}

🐛 Common Issues & Solutions
1. Database Connection Failed

Issue: failed to connect database
Solution:

    Verify PostgreSQL is running: sudo systemctl status postgresql

    Check credentials in DSN

    Ensure database exists: psql -l

2. Migration Errors

Issue: Error creating table
Solution:

    Check model struct tags

    Verify database permissions

    Drop and recreate database if needed

3. Data Not Persisting

Issue: Data disappears after restart
Solution:

    Check for db.AutoMigrate() issues

    Verify transaction commits

    Check error handling after db.Create()

📈 Performance Tips

    Connection Pooling: Configure GORM connection pool

    Indexing: Add indexes for frequently queried columns

    Batch Operations: Use batch inserts for multiple records

    Selective Loading: Use Select() to load only needed fields

🤝 Contributing

    Fork the repository

    Create feature branch: git checkout -b feature-name

    Commit changes: git commit -m 'Add feature'

    Push to branch: git push origin feature-name

    Submit pull request

📄 License

This project is licensed under the MIT License.
🔗 Useful Links

    GORM Documentation

    PostgreSQL Go Driver

    Go Standard Library

    Project Repository

👥 Authors

    Ashraful - Initial work

🙏 Acknowledgments

    GORM team for the excellent ORM library

    PostgreSQL community

    Go developers community

Note: This is a demonstration project. For production use, add proper error handling, logging, security measures, and environment-based configuration.

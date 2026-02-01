Go Backend Project Documentation
📁 Project Structure Overview

1st Part:

This project is a Go backend application with PostgreSQL database integration using GORM ORM.
📦 Main Components
1. main.go - Application Entry Point

The main application file that coordinates database operations.
Features:

    Database connection setup

    Table creation and migration

    Data insertion and retrieval

    Data cleanup operations

Key Functions:

connectDB() - Database Connection
go

// Establishes connection to PostgreSQL
// Uses DSN: host=localhost, database=sampledb, user=postgres
// Timezone: Asia/Dhaka

clearOldData(db) - Data Cleanup
go

// Clears existing users from the database
// Executes: DELETE FROM users

createTable(db) - Table Migration
go

// Automatically creates User table using GORM AutoMigrate
// Based on models.User struct definition

insertData(db) - Sample Data Insertion
go

// Inserts 3 sample user records:
// 1. John Doe (xyz@mail.com)
// 2. Jane Smith (smith@mail.com)
// 3. Alice Johnson (alice@mail.com)

readData(db) - Data Retrieval
go

// Reads all users from database using raw SQL query
// Displays: ID, Name, Email for each user

2. models/user.go - Database Model Definition
go

// User model with GORM conventions
// Fields: ID, CreatedAt, UpdatedAt, DeletedAt (automatic from gorm.Model)
// Custom fields: Name, Email (both strings)

3. request/ - Request Structures Package
ProductRequest

Used for product-related operations (CRUD)
go

// Fields with validation tags:
// - Name: Required
// - Price: Required, must be > 0
// - Stock: Must be ≥ 0
// - Description & Category: Optional

SearchProductRequest

Used for product search with filters
go

// Search parameters:
// - Query: Text search
// - Category: Filter by category
// - Price Range: MinPrice, MaxPrice
// - Tags: Multiple tags filter
// - Pagination: Page, Limit

User-related Requests

    UserRequest: For creating new users (with validation)

    UpdateUserRequest: For updating users (partial updates allowed)

    LoginRequest: For user authentication

4. response/ - Standardized API Response Package
APIResponse Struct

Standard response format for all API endpoints
go

type APIResponse struct {
    Success bool        // Operation status
    Message string      // Human-readable message
    Data    interface{} // Response data (optional)
    Error   string      // Error details (optional)
}

Helper Functions:

    SuccessResponse(): Creates success responses

    ErrorResponse(): Creates error responses

🛠️ Technical Stack

    Language: Go (Golang)

    ORM: GORM v2

    Database: PostgreSQL

    Project Structure: Modular (models, request, response packages)

🚀 How to Run

    Prerequisites:

        Go installed

        PostgreSQL running on localhost

        Database named "sampledb" created

    Setup:

bash

# Install dependencies
go mod init github.com/ashraful/go-backend
go get gorm.io/gorm
go get gorm.io/driver/postgres

# Run the application
go run main.go

🔄 Workflow

When main.go executes:

    Connects to PostgreSQL database

    Clears existing user data (if any)

    Creates/Updates User table structure

    Inserts 3 sample user records

    Reads and displays all users from database

📝 Notes

    The code uses both raw SQL queries and GORM's ORM features

    Models follow GORM conventions with automatic timestamps

    Request validation uses struct tags (validate:"...")

    Standardized response format ensures consistent API responses

    Timezone is set to Asia/Dhaka for database operations

⚠️ Important Observations

    Email Bug: First user's email has leading space (" xyz@mail.com")

    Raw SQL Usage: readData() uses raw SQL instead of GORM's query builder

    Context Usage: Only raw query uses context, other operations don't

    No Error Handling: Missing proper error handling in some places

    Hardcoded DSN: Database credentials are hardcoded (should use env vars)

🔧 Potential Improvements

    Use environment variables for database configuration

    Add proper error handling throughout

    Use GORM's query builder instead of raw SQL

    Implement repository pattern for database operations

    Add structured logging

    Implement database transactions for data integrity


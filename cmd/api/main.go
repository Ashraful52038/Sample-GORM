package main

import (
	"context"
	"fmt"

	"github.com/ashraful/go-backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	db := connectDB()

	clearOldData(db)

	createTable(db)

	insertData(db)

	readData(db)

}

func connectDB() *gorm.DB {
	dsn := "host=localhost user=postgres dbname=sampledb port=5432 sslmode=disable timezone=Asia/Dhaka"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect database")
	}

	fmt.Println("Database connected successfully:", db)

	return db
}

func clearOldData(db *gorm.DB) {
	db.Exec("Delete from users")
	fmt.Println("Old data cleared successfully")
}

func createTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}
	fmt.Println("Table created successfully")
}

func insertData(db *gorm.DB) {

	users := []models.User{
		{Name: "John Doe", Email: " xyz@mail.com"},
		{Name: "Jane Smith", Email: "smith@mail.com"},
		{Name: "Alice Johnson", Email: "alice@mail.com"},
	}

	db.Create(&users)
	fmt.Println("3 user added")

}

func readData(db *gorm.DB) {

	type User struct {
		ID    int
		Name  string
		Email string
	}

	users, err := gorm.G[User](db).
		Raw("SELECT id, name, email from users").
		Find(context.Background())

	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	if len(users) == 0 {
		fmt.Println("No users found in the database")
		return
	}

	fmt.Println("\n All users of database")

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}
}

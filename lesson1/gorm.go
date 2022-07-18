package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// go get -u gorm.io/gorm
//

type User struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "jinzhu", Age: 18})

	// Read
	var user User
	db.First(&user, 1) // find user with id 1
	fmt.Println(user.Name)

	// Update - update user's name to "jinzhu"
	db.Model(&user).Update("Name", "jinzhu")

	// // Delete - delete user with id 1
	// db.Delete(&user, 1)
}

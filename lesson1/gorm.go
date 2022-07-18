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
	db.First(&user, 1)                    // find user with id 1
	db.First(&user, "name = ?", "jinzhu") // find user with name jinzhu
	fmt.Println(user.Name)

	// Update - update user's name to "jinzhu"
	db.Model(&user).Update("Name", "yinzhu")

	// Update - update user's name to "jinzhu" and age to 20
	db.Model(&user).Update("Name", "yyz").Update("Age", 20)

	// update - update multiple fields
	db.Model(&user).Updates(map[string]interface{}{"Name": "multiple", "Age": 20})

	// update - update multiple fields with struct
	db.Model(&user).Updates(&User{Name: "struct", Age: 20})

	// // Delete - delete user with id 1
	// db.Delete(&user, 1)
}

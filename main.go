package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"vuegolang/user"
)

func main() {
	dbName := "vuegolang"
	dsn := "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("successfully connected with database", dbName)

	var users []user.User

	db.Find(&users)

	fmt.Println(len(users))
	fmt.Println(len("-----"))

	for i, user := range users {
		fmt.Println(user.Name, i)
	}

	router := gin.Default()
	router.GET("/", handler)
	router.POST("/user", createUser)
	router.Run()

}

func handler(c *gin.Context) {
	dbName := "vuegolang"
	dsn := "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"user":    users,
	})
	//input
	//handler mapping input ke struct
	//service mapping ke struct User
	//repository save struct User ke db
	//db
}

func createUser(c *gin.Context) {
	dbName := "vuegolang"
	dsn := "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepository := user.NewRepository(db)
	user := user.User{
		Name:         "setiawan",
		Occupation:   "bantul",
		PasswordHash: "ashfkjahsdfjkahsdfk",
		Role:         "admin",
		Avatar:       "tes.png",
		Token:        "sadlkfjalk;sjfl;kds",
		CreatedAt:    1670746838,
		UpdatedAt:    1670746838,
	}
	userRepository.Save(user)

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"user":    user,
	})
}

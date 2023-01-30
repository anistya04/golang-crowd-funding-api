package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"vuegolang/handler"
	"vuegolang/user"
)

func main() {
	router := gin.Default()
	router.GET("/", handler2)
	router.POST("/user", createUser)
	router.POST("/login", login)
	router.POST("/user/check-email-availability", checkEmailAvailability)
	router.Run("127.0.0.1:8080")

}

func handler2(c *gin.Context) {
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
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	response := userHandler.RegisterUser(c)

	c.JSON(http.StatusOK, response)
}

func login(c *gin.Context) {
	dbName := "vuegolang"
	dsn := "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	response := userHandler.Login(c)

	c.JSON(http.StatusOK, response)
}

func checkEmailAvailability(c *gin.Context) {
	dbName := "vuegolang"
	dsn := "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	response := userHandler.CheckEmailAvailability(c)

	c.JSON(http.StatusOK, response)
}

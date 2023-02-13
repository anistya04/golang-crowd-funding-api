package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"vuegolang/auth"
	"vuegolang/campaign"
	"vuegolang/handler"
	"vuegolang/helper"
	"vuegolang/user"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	router := gin.Default()
	router.GET("/", handler2)
	router.POST("/user", createUser)
	router.POST("/login", login)
	router.POST("/user/check-email-availability", checkEmailAvailability)
	router.POST("/user/upload-avatar", authMiddleware, uploadAvatar)

	// campaigns
	router.GET("/campaigns", authMiddleware, getAllCampaigns)

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

// USER
func createUser(c *gin.Context) {
	dbName := "vuegolang"
	dsn := "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewJwtService()
	userHandler := handler.NewUserHandler(userService, authService)

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
	authService := auth.NewJwtService()
	userHandler := handler.NewUserHandler(userService, authService)

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
	authService := auth.NewJwtService()
	userHandler := handler.NewUserHandler(userService, authService)

	response := userHandler.CheckEmailAvailability(c)

	c.JSON(http.StatusOK, response)
}

func uploadAvatar(c *gin.Context) {

	dbName := "vuegolang"
	dsn := "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewJwtService()
	userHandler := handler.NewUserHandler(userService, authService)

	userHandler.UploadAvatar(c)
}

func authMiddleware(c *gin.Context) {
	dbName := "vuegolang"
	dsn := "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepository := user.NewRepository(db)
	authHeader := c.GetHeader("authorization")
	if strings.Contains(authHeader, "bearer") {
		response := helper.ApiResponse("unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}
	// validate token
	arrayToken := strings.Split(authHeader, " ")
	authService := auth.NewJwtService()
	token, er := authService.ValidateToken(arrayToken[1])

	if er != nil {
		response := helper.ApiResponse("unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	claim, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		response := helper.ApiResponse("unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	userId := int(claim["userId"].(float64))
	userService := user.NewService(userRepository)

	userData, e := userService.GetByID(userId)

	if e != nil {
		response := helper.ApiResponse("unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	c.Set("currentUser", userData)
}

// CAMPAIGNS
func getAllCampaigns(c *gin.Context) {
	dbName := "vuegolang"
	dsn := "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	campaignHandler.GetCampaigns(c)
}

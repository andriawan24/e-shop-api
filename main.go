package main

import (
	"e-shop/src/auth"
	"e-shop/src/carts"
	"e-shop/src/handler"
	"e-shop/src/products"
	"e-shop/src/users"
	"e-shop/src/utils"
	"log"
	"net/http"
	"os"
	"strings"

	_ "e-shop/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// @title E-Shop API
// @version 1.0
// @description API ini dibuat untuk memenuhi salah satu syarat dalam proses perekrutan internship di synapsis.id

// @contact.name Naufal Fawwaz Andriawan
// @contact.email fawaznaufal23@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	productRepository := products.NewRepository(db)
	productService := products.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	cartRepository := carts.NewRepository(db)
	cartService := carts.NewService(cartRepository)
	cartHandler := handler.NewCartsHandler(cartService)

	if err != nil {
		log.Fatal("Error while connecting to SQL " + err.Error())
	}

	router := gin.Default()
	router.Use(cors.Default())

	secretKey, _ := utils.GetSecretKey()
	cookieStore := cookie.NewStore([]byte(secretKey))
	router.Use(sessions.Sessions("e-shop", cookieStore))

	api := router.Group("/api/v1")

	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping success",
		})
	})

	// Auth
	api.POST("/sign-up", userHandler.RegisterUser)
	api.POST("/sign-in", userHandler.Login)
	api.POST("/me", userHandler.FetchUser)
	api.PUT("/update-profile", authMiddleware(authService, userService), userHandler.UpdateUser)
	api.GET("/me", authMiddleware(authService, userService), userHandler.FetchUser)

	// Products
	api.GET("/categories", productHandler.GetCategories)
	api.GET("/products", productHandler.GetProducts)

	// Carts
	api.GET("/carts", authMiddleware(authService, userService), cartHandler.GetUserCart)
	api.POST("/carts", authMiddleware(authService, userService), cartHandler.SaveOrUpdateCart)
	api.DELETE("/carts", authMiddleware(authService, userService), cartHandler.RemoveCart)
	api.DELETE("/carts/product", authMiddleware(authService, userService), cartHandler.RemoveProduct)

	// Documentation URL
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}

func authMiddleware(authService auth.Service, userService users.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := utils.APIResponse(
				"Unauthorized",
				http.StatusUnauthorized,
				"error",
				nil,
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Split between Bearer and token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := utils.APIResponse(
				"Unauthorized",
				http.StatusUnauthorized,
				"error",
				nil,
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := utils.APIResponse(
				"Unauthorized",
				http.StatusUnauthorized,
				"error",
				nil,
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["users_id"].(float64))

		user, err := userService.GetUserByID(userId)

		if err != nil {
			response := utils.APIResponse(
				"Unauthorized",
				http.StatusUnauthorized,
				"error",
				nil,
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

package main

import (
	"github.com/Kawar1mi/go-jwt/internals/controllers"
	"github.com/Kawar1mi/go-jwt/internals/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}

package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Kawar1mi/go-jwt/internals/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println(err)
	}

	expTime, err := claims.GetExpirationTime()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if time.Now().Unix() > expTime.Unix() {
		log.Println("token expired")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	userID, err := claims.GetSubject()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	user := &models.User{}
	user.GetUserByID(userID)

	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Set("user", user)

	c.Next()
}

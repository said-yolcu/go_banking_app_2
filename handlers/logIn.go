package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/said-yolcu/go_banking_app/initializers"
	"github.com/said-yolcu/go_banking_app/models"
	"gorm.io/gorm"
)

var expireDuration = time.Minute * 3

func LogIn(ctx *gin.Context) {
	var credentials models.Credentials

	// Bind the post values to the user variable
	if err := ctx.BindJSON(&credentials); err != nil {
		fmt.Println("Mistake: Wrong format entered")
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Check that state id and password are entered
	if credentials.StateID == "" || credentials.Password == "" {
		fmt.Println("Both state id and password must be entered")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Get the user with that state id
	var user models.User
	if err := initializers.DB.Table("users").Where("State_id = ?",
		credentials.StateID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("No such user")
			ctx.AbortWithError(http.StatusNotFound, err)
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Check the password
	if user.Password != credentials.Password {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Create expiration time for the jwt token
	expirationTime := time.Now().Add(expireDuration)

	// Create a new claims instance for use while creating token
	claims := &models.Claims{
		StateID: credentials.StateID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Get the token string
	tokenString, err := token.SignedString(models.JwtKey)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Set the cookie
	http.SetCookie(ctx.Writer,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

	fmt.Println("? Have set the cookie")
	fmt.Println("? Logged in")
}

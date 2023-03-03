package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/said-yolcu/go_banking_app/initializers"
	"github.com/said-yolcu/go_banking_app/models"
)

// Marshals and returns the user entry, given its name surname
func GetUser(ctx *gin.Context) {
	// The requested user
	var requested models.GetUser
	var user models.User

	// If cannot bind context to requested struct, abort
	if err := ctx.BindJSON(&requested); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// The request must specify name and surname
	if requested.Name == "" || requested.Surname == "" {
		fmt.Println("No name or surname")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Find the first instance of the name, surname. Put the entry into construct user
	initializers.DB.Table("users").Where("Name = ?", requested.Name).
		Where("Surname = ?", requested.Surname).First(&user)

	// Return accepted status and user in json
	ctx.JSON(http.StatusOK, &user)
}

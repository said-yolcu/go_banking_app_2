package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/said-yolcu/go_banking_app/initializers"
	"github.com/said-yolcu/go_banking_app/models"
)

// This returns an array containing the users in json format
func GetAllUsers(ctx *gin.Context) {
	var userList []models.User

	if err := initializers.DB.Table("users").Find(&userList).Error; err != nil {

		fmt.Println("Error: internal server error")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &userList)
}

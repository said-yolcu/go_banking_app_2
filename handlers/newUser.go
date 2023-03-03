package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/said-yolcu/go_banking_app/initializers"
	"github.com/said-yolcu/go_banking_app/models"
)

// Creates a new user and adds to the database
func NewUser(ctx *gin.Context) {
	var user models.User

	// Bind the post values to the user variable
	if err := ctx.BindJSON(&user); err != nil {
		fmt.Println("Could not bind JSON")
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Printf("User is %v \n", user)

	// If name and surname are required. If any of them is empty, abort
	if user.Name == "" || user.Surname == "" || user.Email == "" ||
		user.StateID == "" || user.Phone == "" || user.Password == "" ||
		user.Balance == 0 {
		fmt.Println("Mistake: One of the required fields is empty")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Must check if the user with the same name exists or not
	var exists bool
	// Must specify the table in the database to the gorm
	if err := initializers.DB.Table("users").Select("count(*) > 0").
		Where("Name = ?", user.Name).Where("Surname = ?", user.Surname).
		Find(&exists).Error; err != nil {

		fmt.Println("Error: internal server error")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if exists {
		fmt.Println("Mistake: User already exists")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Must check if the state id is a valid Turkish Identification Number
	_, err := strconv.Atoi(user.StateID)
	if len(user.StateID) != 11 || err != nil {
		fmt.Println("Mistake: State ID must be a 11-digit integer")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Create a new user and handle any creation errors
	if err := initializers.DB.Table("users").Create(&user).Error; err != nil {
		fmt.Println("Failed to create the user")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	fmt.Println("? Created the user")

	ctx.Writer.WriteHeader(http.StatusCreated)
}

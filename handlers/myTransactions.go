package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/said-yolcu/go_banking_app/initializers"
	"github.com/said-yolcu/go_banking_app/models"
)

func MyTransactions(ctx *gin.Context, transaction *models.Transaction) {

	// User performing the transaction
	stateId := ctx.Value("stateId")
	fmt.Printf("State id from the cookie is %d\n", stateId)

	// Check if the user with the specified id exist
	var userExists bool
	if err := initializers.DB.Table("users").Select("count(*) > 0").
		Where("State_id = ?", stateId).Find(&userExists).Error; err != nil {

		fmt.Println("? Error: internal server error")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if !userExists {
		fmt.Println("? Error: user not found")
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Get the user with the state id
	var user models.User
	if err := initializers.DB.Table("users").Where("State_id = ?", stateId).
		First(&user).Error; err != nil {

		fmt.Println("? Error: internal server error")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// There are no user ids any more
	// userId := user.UserID
	// fmt.Printf("@myTransactions: user id is %s\n", userId)

	// Connect to the other user
	var transactions []models.Transaction
	if err := initializers.DB.Table("transactions").Where("User_id = ?", stateId).
		Find(&transactions).Error; err != nil {

		fmt.Printf("? Failed to retrieve the transactions ")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Print out the transactions
	fmt.Println(transactions)

	ctx.JSON(http.StatusOK, transactions)
}

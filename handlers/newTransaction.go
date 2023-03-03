package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/said-yolcu/go_banking_app/initializers"
	"github.com/said-yolcu/go_banking_app/models"
)

func NewTransaction(ctx *gin.Context, transaction *models.Transaction) {

	// var transaction models.Transaction

	// if err := ctx.BindJSON(&transaction); err != nil {
	// 	fmt.Println("Cannot bind with json new transaction")
	// 	ctx.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// User performing the transaction
	userId := transaction.UserId

	// Check if the user with the specified id exist
	var userExists bool
	// Must specify the table in the database to the gorm
	if err := initializers.DB.Table("users").Select("count(*) > 0").
		Where("State_id = ?", userId).Find(&userExists).Error; err != nil {

		fmt.Println("Error: internal server error")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Ifthe users do not exist, abort
	if !userExists {
		fmt.Println("Mistake: One or both of the users do not exist")
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	// The other end of the transaction
	otherId := transaction.OtherId

	// Check if the other user exists
	var otherExists bool
	if err := initializers.DB.Table("users").Select("count(*) > 0").
		Where("State_id = ?", otherId).Find(&otherExists).Error; err != nil {

		fmt.Println("Error: internal server error")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// If the other user do not exist, abort
	if !otherExists {
		fmt.Println("Mistake: One or both of the users do not exist")
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	// Connect to the user
	var user models.User
	if err := initializers.DB.Table("users").Where("State_id = ?", userId).
		Find(&user).Error; err != nil {

		fmt.Printf("? Failed to retrieve the user with id = %v", userId)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Connect to the other user
	var other models.User
	if err := initializers.DB.Table("users").Where("State_id = ?", otherId).
		Find(&other).Error; err != nil {

		fmt.Printf("? Failed to retrieve the user with id = %v", otherId)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Create the new transaction
	if err := initializers.DB.Table("transactions").Create(&transaction).
		Error; err != nil {

		fmt.Println("? Failed to create transaction")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Printf("? User : %d\n? Other : %d\n? Amount : %d\n",
		user.Balance, other.Balance, transaction.Amount)

	// Check and update the balance
	if user.Balance-transaction.Amount >= 0 &&
		other.Balance+transaction.Amount >= 0 {

		user.Balance -= transaction.Amount
		other.Balance += transaction.Amount
		fmt.Printf("? User balance: %d\n? Other balance: %d\n",
			user.Balance, other.Balance)
	} else {
		// If balances are not suitable, delete the transaction
		if err := initializers.DB.Table("transactions").Delete(&transaction).
			Error; err != nil {
			fmt.Println("Failed to delete the transaction")
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		fmt.Println("? One of the users does not have enough money to carry " +
			"this transaction")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Update the user
	if err := initializers.DB.Table("users").Where("State_id = ?", userId).
		Update("balance", user.Balance).Error; err != nil {
		fmt.Println("Failed to update the user")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Update the other user
	if err := initializers.DB.Table("users").Where("State_id = ?", otherId).
		Update("balance", other.Balance).Error; err != nil {
		fmt.Println("Failed to update the other user")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

}

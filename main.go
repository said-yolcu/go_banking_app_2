package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/said-yolcu/go_banking_app/handlers"
	"github.com/said-yolcu/go_banking_app/initializers"
	"github.com/said-yolcu/go_banking_app/middlewares"
	"github.com/said-yolcu/go_banking_app/models"
)

var (
	server *gin.Engine
	config initializers.Config
	err    error
)

func init() {
	config, err = initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	// Returns a default engine
	server = gin.Default()

	fmt.Println("Initialized the main file")
}

func main() {

	// Create users table
	initializers.DB.Table("users").AutoMigrate(&models.User{})
	fmt.Println("? Users migration complete")

	// Create transactions table
	initializers.DB.AutoMigrate(&models.Transaction{})
	fmt.Println("? Transactions migration complete")

	// Create a server group for users
	userRouter := server.Group("/user")

	// Add handler functions for user
	//userRouter.GET("/get_user", handlers.GetUser)
	//userRouter.GET("/get_all_users", handlers.GetAllUsers)
	userRouter.POST("/sign_up", handlers.NewUser)
	userRouter.POST("/log_in", handlers.LogIn)

	// Create a server group for transactions
	//trxRouter := server.Group("/transaction")

	// Add handler functions for transactions
	//	trxRouter.GET("/get_transaction", handlers.GetTransaction)
	//	trxRouter.GET("/get_all_transactions", handlers.GetAllTransactions)

	// I had to transfer the /new_transaction to the /user middleware. Otherwise
	// cookies are inaccessible
	userRouter.POST("/new_transaction", middlewares.Authenticate(handlers.NewTransaction))
	userRouter.GET("/my_transactions", middlewares.Authenticate(handlers.MyTransactions))

	// Run the server and print out any errors
	log.Fatal(server.Run(":" + config.ServerPort))
}

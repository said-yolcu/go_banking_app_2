package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/said-yolcu/go_banking_app/initializers"
	"github.com/said-yolcu/go_banking_app/models"
)

// Returns an authenticater function, that authenticates the user before
// calling the handler function
func Authenticate(originalHandler func(*gin.Context, *models.Transaction)) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		fmt.Println("? Running authentication function")

		// Retrieve the cookie
		cookie, err := ctx.Request.Cookie("token")
		// fmt.Println(ctx.Request)
		// fmt.Println(ctx.Request.Cookies())

		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Println("? No cookies found")
				ctx.Writer.WriteHeader(http.StatusUnauthorized)
				return
			}
			fmt.Println("? Cannot retrieve cookie")
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			return
		}

		// Destructure the cookie
		tokenStr := cookie.Value

		claims := &models.Claims{}

		// Retrieve the token
		tkn, err := jwt.ParseWithClaims(tokenStr, claims,
			func(t *jwt.Token) (interface{}, error) {
				return models.JwtKey, nil
			})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				fmt.Println("? Trying to access unauthorizedly")
				ctx.Writer.WriteHeader(http.StatusUnauthorized)
				return
			}
			ctx.Writer.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			fmt.Println("? Token is not valid")
			ctx.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Get the path of the router
		path := ctx.FullPath()
		paths := strings.Split(path, "/")
		lastPath := paths[len(paths)-1]

		fmt.Printf("last path: %s\n", lastPath)
		fmt.Printf("path: %s\n", path)

		// New transaction
		var transaction models.Transaction

		if lastPath == "new_transaction" {

			if err := ctx.BindJSON(&transaction); err != nil {
				fmt.Println("Cannot bind with json new transaction")
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			// Get the request body
			fmt.Printf("body: %v\n", transaction)

			// Get the user with the specified id
			var user models.User
			if err := initializers.DB.Table("users").Where("State_id = ?", transaction.UserId).
				Find(&user).Error; err != nil {
				fmt.Println("? Cannot get the user with that id")
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			// This becomes unnecessary as we use state id to
			// identify a user

			// // Check if the user state id and claimed state id are equal
			// if user.StateID != claims.StateID {
			// 	fmt.Printf("user: %s\nclaims: %s\n", user.StateID, claims.StateID)
			// 	fmt.Println("? Different user: Trying to access unauthorizedly")
			// 	ctx.Writer.WriteHeader(http.StatusUnauthorized)
			// 	return
			// }

			fmt.Println("? Calling the original handler")

			// Call the original handler
			originalHandler(ctx, &transaction)
		} else if lastPath == "my_transactions" {
			ctx.Set("stateId", claims.StateID)
			fmt.Printf("@authenticate Claimed state id %s\n", claims.StateID)

			fmt.Println("? Calling the original handler")

			// Call the original handler
			originalHandler(ctx, &models.Transaction{})
		}
	})
}

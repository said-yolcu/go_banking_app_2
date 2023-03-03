# go_banking_app

In this project I will try to build a banking app using go. 

This project is composed of two parts: back-end and front-end. This repo implements the first part. In this repo a server is built with Golang that is connected to a Postgresql database. The front-end is implemented with Node.js and is a terminal interface that prompts user for requests and accordingly makes requests to the server. Then prints out the response in a user-friendly manner. Front-end is here: https://github.com/said-yolcu/banking_app_node

0. docker-compose.yml
    With docker-compose file we specify the properties of the docker container we will open with docker. We open the database in a server and connect to that server via port 6500.

1. initializers package
    Implements functionality to take configuration specifications from the app.env file with LoadConfig(), and connects to the specified database with ConnectDB.

2. middlewares package
    Authenticate() function takes in a ginHandler function. It retrieves the cookie, then it compares the values in the cookie with parameter values depending on the parameter handler func (for example it compares the transaction's user id with the user id in cookie for new_transaction path). Then if the user is successfully authenticated, it authorizes the user to continue and calls the parameter handler function.

3. handlers package
    1. getAllUsers.go file
        It finds and returns all the users in the database. This function is not called by the front-end.
    2. getUser.go file
        Queries the DB and returns the user with the given name and surname. Not called by front-end either.
    3. logIn.go file
        Logs in to the system as the user with the specified state id and password. Then creates a cookie with an expiration date to store this login info.
    4. myTransactions.go file
        Prints out the transactions of the user on the cookie.
    5. newTransactions.go file
        Makes a new transaction for the user with the specified user_id. This function must be changed in the future to not require the user_id of the user that is making the transaction. It should just require the state_id of the user on the other end of the transaction.
    6. newUser.go file
        Creates a new user.

4. models package
    Most used structs are User and Transaction structs. The are are some unnecessary structs in that package. Unnecessary struct must be weeded out.

5. main.go file
    Firstly, we connect to the DB and create a gin engine (server) in init() function.

    In main() function, we migrate (create) one table for users and one for transactions. Then we create a router and add handlers for sign-up, log-in, making a new transaction, seeing logged-in user's all transactions. At the end we run the server and log the messages to the terminal.

    Further handlers must be added here. Two that come to mind are log-out and sign-out handlers.

# go_banking_app_2

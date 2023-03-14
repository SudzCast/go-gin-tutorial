package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var users = []User{
	{ID: 1, FirstName: "Person", LastName: "One"},
	{ID: 2, FirstName: "Person", LastName: "Two"},
	{ID: 3, FirstName: "Person", LastName: "Three"},
}

func main() {
	router := gin.Default()
	router.GET("/users", getAllUsers)
	router.POST("/users", addNewUser)
	router.GET("/users/:id", getUserById)
	router.DELETE("/users/:id", deleteUser)
	router.Run("localhost:8080")
}

func getAllUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, users)
}

func addNewUser(context *gin.Context) {
	var newUser User
	if err := context.BindJSON(&newUser); err != nil {
		fmt.Println(err.Error())
		return
	}

	newUser.ID = users[len(users)-1].ID + 1

	users = append(users, newUser)
	context.IndentedJSON(http.StatusCreated, newUser)
}

func getUserById(context *gin.Context) {
	userId, _ := strconv.Atoi(context.Param("id"))
	for _, user := range users {
		if user.ID == userId {
			context.IndentedJSON(http.StatusFound, user)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, "No user found")
}

func deleteUser(context *gin.Context) {
	userId, _ := strconv.Atoi(context.Param("id"))
	recordsDeleted := false
	for i, user := range users {
		if user.ID == userId {
			users[i] = users[len(users)-1]
			users = users[:len(users)-1]
			recordsDeleted = true
		}
	}
	if recordsDeleted {
		context.IndentedJSON(http.StatusOK, users)
	} else {
		context.IndentedJSON(http.StatusNotFound, "No user with matching id")
	}
}

package main

import (
	"awesomeProject1/controllers"
	"awesomeProject1/initializers"
	"awesomeProject1/repositories"
	"fmt"
	"net/http"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	createUser()
}

func main() {
	http.HandleFunc("/jwt/create", controllers.Create)
	http.HandleFunc("/jwt/refresh", controllers.Refresh)
	http.ListenAndServe(":8080", nil)
}

func createUser() {
	_, err := repositories.CreateUser()
	if err != nil {
		fmt.Println(err)
	}
}

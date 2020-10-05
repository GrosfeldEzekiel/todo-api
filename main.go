package main

import (
	"LearningGo/section4/todo-api/controller"
	"LearningGo/section4/todo-api/model"
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()

	db := model.Connect()
	defer db.Close() // Run the code at the end

	fmt.Println("Serving...")

	log.Fatal(http.ListenAndServe(":3000", mux))
}

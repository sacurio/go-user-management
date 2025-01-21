package main

import (
	"fmt"
	"log"
	h "net/http"
	"user-management/internal/db"
	"user-management/internal/handler/http"
	"user-management/internal/repository"
	"user-management/internal/service"
)

func main() {
	fmt.Println("Starting User Management Service...")

	// userHandler := handler.NewUserHandler()

	// userHandler.CreateUser("John Doe", "john.doe@mail.com")
	// userHandler.CreateUser("Patrick Doe", "patrick.doe@mail.com")

	// fmt.Println("List of users:")
	// userHandler.ListUsers()

	database, err := db.InitDatabase()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer database.Close()

	userRepo := repository.NewRepositoryDB(database)
	userService := service.NewUserService(userRepo)

	userHandler := http.NewUserHTTPHandler(userService)

	h.HandleFunc("/users", func(w h.ResponseWriter, r *h.Request) {
		if r.Method == h.MethodGet {
			userHandler.ListUsers(w, r)
		} else if r.Method == h.MethodPost {
			userHandler.CreateUser(w, r)
		} else {
			h.Error(w, "Method not allowed", h.StatusMethodNotAllowed)
		}
	})

	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(h.ListenAndServe(port, nil))
}

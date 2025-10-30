package main

import (
	"log"
	"net/http"
	"server/dbconfig"
	"server/handler"
	"server/middleware"
	"server/routes/manage-users"

	"github.com/gorilla/mux"
)


func main() {
	dbconfig.ConnectDB()
	defer dbconfig.Database.Close()

	router := mux.NewRouter()
	protected := router.PathPrefix("/protected").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	/* Authentication Routes */
	router.HandleFunc("/api/signin", handler.Signin).Methods("POST")

	/* User Management Routes */
	protected.HandleFunc("/api/users/add", manageusers.AddUserHandler).Methods("POST")
	protected.HandleFunc("/api/users", manageusers.GetUserHandler).Methods("GET")
	protected.HandleFunc("/api/users/get", manageusers.GetIndividualUserHandler).Methods("POST")
	protected.HandleFunc("/api/users/update-role", manageusers.UpdateUserRoleHandler).Methods("PUT")
	protected.HandleFunc("/api/users/update-acc-status", manageusers.UpdateUserStatusHandler).Methods("PUT")
	protected.HandleFunc("/api/users/update", manageusers.UpdateUserDetailsHandler).Methods("PUT")

	/* Manager And Labour Association Routes */

	/* Start Server */
	port := ":8080"
	log.Printf("Server starting at http://localhost%s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Server failed:", err)
	}
}

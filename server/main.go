package main

import (
	"log"
	"net/http"
	"server/dbconfig"
	"server/handlefile"
	"server/handlebills"
	"server/handler"
	"server/middleware"
	manageexpense "server/routes/manage-expense"
	managelabourassociation "server/routes/manage-labour-association"
	managelabourattendance "server/routes/manage-labour-attendance"
	manageprojects "server/routes/manage-projects"
	managestocks "server/routes/manage-stocks"
	manageusers "server/routes/manage-users"
	materialtracking "server/routes/material-tracking"

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
	protected.HandleFunc("/api/association", managelabourassociation.GetAssociationHandler).Methods("GET")
	protected.HandleFunc("/api/association/add", managelabourassociation.AddManagerLabourAssociationHandler).Methods("POST")
	protected.HandleFunc("/api/association/update", managelabourassociation.UpdateAssociationHandler).Methods("PUT")
	protected.HandleFunc("/api/association/delete", managelabourassociation.DeleteAssociationHandler).Methods("DELETE")

	/* Manage Labour Attendance Routes */
	protected.HandleFunc("/api/attendance/update", managelabourattendance.UpdateAttendanceHandler).Methods("PUT")
	protected.HandleFunc("/api/attendance/add", managelabourattendance.AddAttendanceHandler).Methods("POST")
	protected.HandleFunc("/api/attendance", managelabourattendance.GetAttendanceHandler).Methods("GET")

	/* Manage Stock Routes */
	protected.HandleFunc("/api/stocks/add", managestocks.AddUserHandler).Methods("POST")
	protected.HandleFunc("/api/stocks", managestocks.GetStocksHandler).Methods("GET")
	protected.HandleFunc("/api/stocks/update", managestocks.UpdateStockHandler).Methods("PUT")
	protected.HandleFunc("/api/stocks", managestocks.GetStocksHandler).Methods("DELETE")

	/* Manage Project Routes */
	protected.HandleFunc("/api/projects/add",manageprojects.AddNewProjectHandler).Methods("POST")
	router.HandleFunc("/api/project/docs",handlefile.ServePDFHandler).Methods("GET")
    protected.HandleFunc("/api/project/getdetails",manageprojects.GetProjectDetails).Methods("GET")
	protected.HandleFunc("/api/project/update-details",manageprojects.UpdateProjectHandler).Methods("PUT")
	protected.HandleFunc("/api/project/delete",manageprojects.DeleteProjectDetails).Methods("DELETE")

	/* Manage Expense Routes */
	protected.HandleFunc("/api/expense/add",manageexpense.AddExpense).Methods("POST")
	protected.HandleFunc("/api/expense/update",manageexpense.UpdateExpense).Methods("PUT")
	protected.HandleFunc("/api/expense/delete",manageexpense.DeleteExpense).Methods("DELETE")
	protected.HandleFunc("/api/expense",manageexpense.GetAllExpenses).Methods("GET")
	router.HandleFunc("/api/getexpensebills",handlebills.ServeExpenseHandler).Methods("GET")
	//http://localhost:8080/api/getexpensebills?file=1762604165_image.png

	/* Material Tracking Routes */
	protected.HandleFunc("/api/material-tracking/add", materialtracking.AddTrackingRecord).Methods("POST")
	protected.HandleFunc("/api/material-tracking", materialtracking.GetTrackingRecords).Methods("GET")
	protected.HandleFunc("/api/material-tracking/update", materialtracking.UpdateTrackingRecord).Methods("PUT")
	protected.HandleFunc("/api/material-tracking/delete", materialtracking.DeleteTrackingRecord).Methods("DELETE")

	/* Start Server */
	port := ":8080"
	log.Printf("Server starting at http://localhost%s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Server failed:", err)
	}
}

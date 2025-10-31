package managelabourattendance

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"server/models"
	"time"
)

func AddAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Only POST is allowed.", http.StatusMethodNotAllowed)
		return
	}

	var attendance models.LabourAttendanceModel
	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if attendance.LabourID == "" || attendance.ManagerID == "" || attendance.AttendanceStatus == "" {
		http.Error(w, "Missing required fields: labour_id, manager_id, or attendance_status", http.StatusBadRequest)
		return
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	attendance.CreatedAt = currentTime
	attendance.UpdatedAt = currentTime

	query := `
		INSERT INTO manage_labours (labour_id, manager_id, attendance_status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING manage_labour_id
	`

	err := dbconfig.Database.QueryRow(
		query,
		attendance.LabourID,
		attendance.ManagerID,
		attendance.AttendanceStatus,
		attendance.CreatedAt,
		attendance.UpdatedAt,
	).Scan(&attendance.ManageLabourID)

	if err != nil {
		http.Error(w, "Database insert error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":    true,
		"message":    "Attendance added successfully",
		"attendance": attendance,
	})
}

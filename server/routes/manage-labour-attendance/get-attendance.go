package managelabourattendance

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"server/models"
)

func GetAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method. Only GET is allowed.", http.StatusMethodNotAllowed)
		return
	}

	query := `
		SELECT 
			la.manage_labour_id,
			la.labour_id,
			labour.user_name AS labour_name,
			la.manager_id,
			manager.user_name AS manager_name,
			la.attendance_status,
			la.created_at,
			la.updated_at
		FROM manage_labours la
		JOIN users labour ON la.labour_id::int = labour.user_id
		JOIN users manager ON la.manager_id::int = manager.user_id
		ORDER BY la.created_at DESC
	`

	rows, err := dbconfig.Database.Query(query)
	if err != nil {
		http.Error(w, `{"success": false, "message": "Database query error: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var attendances []models.AttendanceWithNames

	for rows.Next() {
		var a models.AttendanceWithNames
		if err := rows.Scan(
			&a.ManageLabourID,
			&a.LabourID,
			&a.LabourName,
			&a.ManagerID,
			&a.ManagerName,
			&a.AttendanceStatus,
			&a.CreatedAt,
			&a.UpdatedAt,
		); err != nil {
			http.Error(w, `{"success": false, "message": "Error scanning result: `+err.Error()+`"}`, http.StatusInternalServerError)
			return
		}
		attendances = append(attendances, a)
	}

	if len(attendances) == 0 {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "No attendance records found",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":     true,
		"message":     "Attendance records fetched successfully",
		"attendances": attendances,
	})
}

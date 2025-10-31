package managelabourattendance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/dbconfig"
	"strings"
	"time"
)

type UpdateAttendanceInput struct {
	ManageLabourID   int     `json:"manage_labour_id"`
	LabourID         *string `json:"labour_id,omitempty"`
	ManagerID        *string `json:"manager_id,omitempty"`
	AttendanceStatus *string `json:"attendance_status,omitempty"`
}

func UpdateAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, `{"success": false, "message": "Invalid request method. Use POST or PUT."}`, http.StatusMethodNotAllowed)
		return
	}

	var input UpdateAttendanceInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"success": false, "message": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if input.ManageLabourID <= 0 {
		http.Error(w, `{"success": false, "message": "manage_labour_id is required and must be valid"}`, http.StatusBadRequest)
		return
	}

	setClauses := []string{}
	params := []interface{}{}
	paramIndex := 1

	if input.LabourID != nil {
		setClauses = append(setClauses, fmt.Sprintf("labour_id = $%d", paramIndex))
		params = append(params, *input.LabourID)
		paramIndex++
	}
	if input.ManagerID != nil {
		setClauses = append(setClauses, fmt.Sprintf("manager_id = $%d", paramIndex))
		params = append(params, *input.ManagerID)
		paramIndex++
	}
	if input.AttendanceStatus != nil {
		setClauses = append(setClauses, fmt.Sprintf("attendance_status = $%d", paramIndex))
		params = append(params, *input.AttendanceStatus)
		paramIndex++
	}

	if len(setClauses) == 0 {
		http.Error(w, `{"success": false, "message": "No fields provided for update"}`, http.StatusBadRequest)
		return
	}

	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", paramIndex))
	params = append(params, time.Now())
	paramIndex++

	query := fmt.Sprintf(`UPDATE manage_labours SET %s WHERE manage_labour_id = $%d`,
		strings.Join(setClauses, ", "), paramIndex)
	params = append(params, input.ManageLabourID)

	result, err := dbconfig.Database.Exec(query, params...)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"success": false, "message": "Database update error", "error": "%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, `{"success": false, "message": "Error fetching update result"}`, http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, `{"success": false, "message": "No record found with the given manage_labour_id"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Attendance record updated successfully",
	})
}

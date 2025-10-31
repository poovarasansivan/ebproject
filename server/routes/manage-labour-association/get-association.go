package managelabourassociation

import (
	"encoding/json"
	"net/http"
	"server/dbconfig"
	"server/models"
)

func GetAssociationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method. Only GET is allowed.", http.StatusMethodNotAllowed)
		return
	}

	query := `
	SELECT 
			la.association_id,
			la.manager_id,
			manager.user_name AS manager_name,
			manager.user_email AS manager_email,
			manager.phone_no AS manager_phone,
			manager.role AS manager_role,
			manager.acc_status AS manager_status,
			la.labour_id,
			labour.user_name AS labour_name,
			labour.user_email AS labour_email,
			labour.phone_no AS labour_phone,
			labour.role AS labour_role,
			labour.acc_status AS labour_status,
			la.created_at,
			la.updated_at
		FROM manager_labour_association la
		JOIN users manager ON la.manager_id::int = manager.user_id
		JOIN users labour ON la.labour_id::int = labour.user_id
		ORDER BY la.created_at DESC
	`

	rows, err := dbconfig.Database.Query(query)
	if err != nil {
		http.Error(w, "Database query error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var associations []models.AssociationWithDetails

	for rows.Next() {
		var a models.AssociationWithDetails
		if err := rows.Scan(
			&a.AssociationID,
			&a.ManagerID,
			&a.ManagerName,
			&a.ManagerEmail,
			&a.ManagerPhone,
			&a.ManagerRole,
			&a.ManagerStatus,
			&a.LabourID,
			&a.LabourName,
			&a.LabourEmail,
			&a.LabourPhone,
			&a.LabourRole,
			&a.LabourStatus,
			&a.CreatedAt,
			&a.UpdatedAt,
		); err != nil {
			http.Error(w, "Error scanning result: "+err.Error(), http.StatusInternalServerError)
			return
		}
		associations = append(associations, a)
	}

	if len(associations) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "No labour associations found",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":      true,
		"message":      "Labour associations fetched successfully",
		"associations": associations,
	})
}

package models

type LabourAssociationModel struct {
	AssociationID int    `json:"association_id"`
	ManagerID     string `json:"manager_id"`
	LabourID      string `json:"labour_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type AssociationWithDetails struct {
	AssociationID int    `json:"association_id"`
	ManagerID     string `json:"manager_id"`
	ManagerName   string `json:"manager_name"`
	ManagerEmail  string `json:"manager_email"`
	ManagerPhone  string `json:"manager_phone"`
	ManagerRole   string `json:"manager_role"`
	ManagerStatus string `json:"manager_status"`
	LabourID      string `json:"labour_id"`
	LabourName    string `json:"labour_name"`
	LabourEmail   string `json:"labour_email"`
	LabourPhone   string `json:"labour_phone"`
	LabourRole    string `json:"labour_role"`
	LabourStatus  string `json:"labour_status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

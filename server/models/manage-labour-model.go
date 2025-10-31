package models

type LabourAttendanceModel struct {
	ManageLabourID   int    `json:"manage_labour_id"`
	LabourID         string `json:"labour_id"`
	ManagerID        string `json:"manager_id"`
	AttendanceStatus string `json:"attendance_status"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type AttendanceWithNames struct {
	ManageLabourID   int    `json:"manage_labour_id"`
	LabourID         string `json:"labour_id"`
	LabourName       string `json:"labour_name"`
	ManagerID        string `json:"manager_id"`
	ManagerName      string `json:"manager_name"`
	AttendanceStatus string `json:"attendance_status"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

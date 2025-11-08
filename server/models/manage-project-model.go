package models

type ProjectModel struct {
	ProjectID          int    `json:"project_id"`
	FDNumber           string `json:"fd_number"`
	ProjectName        string `json:"project_name"`
	ProjectDescription string `json:"project_description"`
	ProjectBudget      float64 `json:"project_amount"`
	Status             string `json:"status"`
	AgreementDocs      string `json:"agreement_docs"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

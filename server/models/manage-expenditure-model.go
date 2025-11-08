package models

type ManageExpenditure struct {
	ExpenseID          int    `json:"expense_id"`
	LabourID           int    `json:"labour_id"`
	ManagerID          string `json:"manager_id"`
	ExpenseDescription string `json:"expense_description"`
	TotalExpense       string `json:"total_expense"`
	BillDocs           string `json:"bill_docs"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}

// type ManageExpenditureResponse struct {
// 	ExpenseID          int    `json:"expense_id"`
// 	LabourID           int    `json:"labour_id"`
// 	LabourName         string `json:"user_name"`
// 	ManagerID          string `json:"manager_id"`
// 	ManagerNames       string `json:"user_name"`
// 	ExpenseDescription string `json:"expense_description"`
// 	TotalExpense       string `json:"total_expense"`
// 	BillDocs           string `json:"bill_docs"`
// 	CreatedAt          string `json:"created_at"`
// 	UpdatedAt          string `json:"updated_at"`
// }

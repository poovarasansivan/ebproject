package models

type MaterialTrackingModel struct {
	TrackingID     int    `json:"tracking_id"`
	StockID        int    `json:"stock_id"`
	ProjectID      int    `json:"project_id"`
	NoOfStocks     int    `json:"no_of_stocks_used"`
	ImageProof     string `json:"image_proof"`
	MachineRunTime string `json:"machines_running_time"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type TrackingRecordResponse struct {
	TrackingID     int    `json:"tracking_id"`
	StockID        int    `json:"stock_id"`
	StockName      string `json:"stock_name"`
	ProjectID      int    `json:"project_id"`
	ProjectName    string `json:"project_name"`
	NoOfStocks     int    `json:"no_of_stocks_used"`
	ImageProof     string `json:"image_proof"`
	MachineRunTime string `json:"machines_running_time"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
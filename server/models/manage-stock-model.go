package models

type ManageStocksModel struct {
	StockID        int    `json:"stock_id"`
	StockName      string `json:"stock_name"`
	StockType      string `json:"stock_type"`
	AvailableStock int    `json:"available_stock"`
	StockConsumed  int    `json:"stock_delivered_out"`
	ReturnedStock  int    `json:"returned_stock"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

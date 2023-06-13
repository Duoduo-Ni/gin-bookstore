package model

// Data 结构
type Data struct {
	Amount      float64 `json:"amount"`
	TotalAmount float64 `json:"totalAmount"`
	TotalCount  int64   `json:"totalCount"`
}

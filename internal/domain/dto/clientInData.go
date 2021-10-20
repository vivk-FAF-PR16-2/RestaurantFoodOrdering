package dto

type ClientInData struct {
	OrderId int           `json:"order_id"`
	Orders  []OrderInData `json:"orders"`
}

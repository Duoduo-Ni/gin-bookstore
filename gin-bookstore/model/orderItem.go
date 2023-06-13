package model

// OrderItem 结构
type OrderItem struct {
	OrderItemID int64   `json:"orderItemID"` //订单项的id
	Count       int64   `json:"count"`       //订单项中图书的数量
	Amount      float64 `json:"amount"`      //订单项中图书的金额小计
	Title       string  `json:"title"`       //订单项中图书的书名
	Author      string  `json:"author"`      //订单项中图书的作者
	Price       float64 `json:"price"`       //订单项中图书的价格
	ImgPath     string  `json:"imgPath"`     //订单项中图书的封面
	OrderID     string  `json:"orderID"`     //订单行所属的订单
}

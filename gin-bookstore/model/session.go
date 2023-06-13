package model

// 用户每次重新登录或浏览器重启，都会生成新的sessionID（即新的cookie）
// Session 结构
type Session struct {
	SessionID string   `json:"sessionID"`
	UserName  string   `json:"userName"`
	UserID    int      `json:"userID"`
	Cart      *Cart    `json:"cart"`
	OrderID   string   `json:"orderID"`
	Orders    []*Order `json:"orders"`
}

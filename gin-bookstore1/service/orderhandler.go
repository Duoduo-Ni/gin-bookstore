package service

import (
	"gin-bookstore1/dao"
	//"gin-bookstore1/model"
	//"gin-bookstore1/utils"
	"html/template"
	"net/http"
	//"time"
)

// GetOrders 获取所有订单
func GetOrders(w http.ResponseWriter, r *http.Request) {
	//调用dao中获取所有订单的函数
	orders, _ := dao.GetOrders()
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	//执行
	t.Execute(w, orders)
}

// GetOrderInfo 获取订单对应的订单项
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	//获取订单号
	orderID := r.FormValue("orderId")
	//根据订单号调用dao中获取所有订单项的函数
	orderItems, _ := dao.GetOrderItemsByOrderID(orderID)
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	//执行
	t.Execute(w, orderItems)
}

// SendOrder 发货
func SendOrder(w http.ResponseWriter, r *http.Request) {
	//获取要发货的订单号
	orderID := r.FormValue("orderId")
	//调用dao中的更新订单状态的函数
	dao.UpdateOrderState(orderID, 1)
	//调用GetOrders函数再次查询一下所有的订单
	GetOrders(w, r)
}

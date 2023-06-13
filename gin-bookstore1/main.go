package main

import (
	"gin-bookstore1/service"
	"net/http"
)

// 管理员部分接口
func main() {
	//设置处理静态资源，如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	//获取带分页的图书信息
	http.HandleFunc("/getPageBooks", service.GetPageBooks)
	//去更新或添加图书的页面
	http.HandleFunc("/toUpdateBookPage", service.ToUpdateBookPage)
	//更新或添加图书的提交按钮
	http.HandleFunc("/updateOraddBook", service.UpdateOraddBook)
	//删除图书
	http.HandleFunc("/deleteBook", service.DeleteBook)

	//获取所有订单
	http.HandleFunc("/getOrders", service.GetOrders)
	//获取订单详情，即订单所对应的所有的订单项
	http.HandleFunc("/getOrderInfo", service.GetOrderInfo)
	//发货
	http.HandleFunc("/sendOrder", service.SendOrder)

	http.ListenAndServe(":8082", nil)

}

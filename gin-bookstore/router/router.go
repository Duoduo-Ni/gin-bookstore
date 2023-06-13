package router

import (
	middlewares "gin-bookstore/cors"
	"gin-bookstore/service"
	"github.com/gin-gonic/gin"
)

// 登录和用户部分接口
func Router() *gin.Engine {

	////设置处理静态资源，如css和js文件
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	////直接去html页面
	//http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	//因为r.run内置了http.ListenAndServe，用框架和原生库的服务想同时启动，同一个端口会被占用，所以将使用http部分的接口（管理员部分）分开在gin-bookstore1项目中

	r := gin.Default()

	r.Use(middlewares.Cors())

	//首页
	r.POST("/main", service.GetPageBooksByPrice)
	//图书封面
	r.POST("getBookCover", service.GetBookCover)
	//登录
	r.POST("/login", service.Login)
	//退出登录
	r.GET("/logout", service.Logout)
	//注册
	r.POST("/reGist", service.Regist)
	//通过Ajax请求验证用户名是否可用
	//r.POST("/checkUserName", service.CheckUserName)
	//管理员登录
	r.POST("/adminLogin", service.AdminLogin)

	////获取带分页的图书信息（图书管理页面）
	//r.POST("/getPageBooks", service.GetPageBooks)
	////添加图书
	//r.POST("/addBook", service.AddBook)
	////去更新或添加图书的页面
	//r.POST("/toUpdateOrAddBookPage", service.ToUpdateOrAddBookPage)
	////更新或添加图书的提交按钮
	//r.POST("/updateOrAddBookButton", service.UpdateOrAddBookButton)
	////删除图书
	//r.POST("/deleteBook", service.DeleteBook)

	//获取带分页的图书信息
	//http.HandleFunc("/getPageBooks", service.GetPageBooks)
	////去更新或添加图书的页面
	//http.HandleFunc("/toUpdateOrAddBookPage", service.ToUpdateOrAddBookPage)
	////更新或添加图书的提交按钮
	//http.HandleFunc("/updateOrAddBookButton", service.UpdateOrAddBookButton)
	////删除图书
	//http.HandleFunc("/deleteBook", service.DeleteBook)

	//获取所有订单
	//r.GET("/getOrders", service.GetOrders)
	//获取订单详情，即订单所对应的所有的订单项
	//r.POST("/getOrderInfo", service.GetOrderInfo)
	//发货
	//r.POST("/sendOrder", service.SendOrder)

	////获取所有订单
	//http.HandleFunc("/getOrders", service.GetOrders)
	////获取订单详情，即订单所对应的所有的订单项
	//http.HandleFunc("/getOrderInfo", service.GetOrderInfo)
	////发货
	//http.HandleFunc("/sendOrder", service.SendOrder)

	//添加图书到购物车中
	r.POST("/addBookToCar", service.AddBookToCar)
	//获取购物车信息
	r.GET("/getCarInfo", service.GetCarInfo)
	//清空购物车
	r.POST("/deleteCar", service.DeleteCar)
	//删除购物项
	r.POST("/deleteCarItem", service.DeleteCartItem)
	//更新购物项
	r.POST("/updateCarItem", service.UpdateCartItem)
	//去结账
	r.GET("/checkOut", service.Checkout)
	//获取我的订单
	r.GET("/getMyOrder", service.GetMyOrders)
	//获取订单详情，即订单所对应的所有的订单项
	r.POST("/getOrderInfo1", service.GetOrderInfo1)
	//确认收货
	r.POST("/takeOrder", service.TakeOrder)

	return r

}

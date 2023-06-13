package main

import (
	"gin-bookstore/router"
)

func main() {
	r := router.Router()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

//package main

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"io/ioutil"
//)
//
//func main() {
//	r := gin.Default()
//	r.POST("getImage", func(c *gin.Context) {
//		getImage(c)
//	})
//	r.Run(":8088") // 监听并在 0.0.0.0:8088 上启动服务
//}
//
//func getImage(c *gin.Context) { //显示图片的方法
//	imageName := c.Query("imageName") //截取get请求参数，也就是图片的路径，可是使用绝对路径，也可使用相对路径
//	//imageName:="F:\F\ruangongkehse\gin-bookstore\views\static\img\C语言入门经典.jpg"
//	fmt.Println(imageName)
//	c.JSON(200, gin.H{
//		"iamgeName": imageName,
//	})
//	file, _ := ioutil.ReadFile(imageName) //把要显示的图片读取到变量中
//	c.Writer.WriteString(string(file))    //关键一步，写给前端
//}

package service

import (
	"gin-bookstore/dao"
	_ "gin-bookstore/model"
	_ "gin-bookstore/utils"
	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	////判断管理员是否已经登录
	//state := dao.AdminIsLogin()
	//if flag {
	//	//已经登录
	//	//去首页
	//	GetPageBooksByPrice(c)
	//} else {
	//	//未登录
	//获取用户名和密码
	adminname := c.PostForm("adminname")
	adminpassword := c.PostForm("adminpassword")
	//调用userdao中验证用户名和密码的方法
	admin, _ := dao.CheckAdminNameAndPassword(adminname, adminpassword)
	//fmt.Println(admin.AdminID, admin.AdminName, admin.AdminPassword)
	if admin.AdminID > 0 {
		//用户名和密码正确
		////生成UUID作为Session的id
		//uuid := utils.CreateUUID()
		////创建一个Session
		//sess := &model.Session{
		//	SessionID: uuid,
		//	UserName:  user.Username,
		//	UserID:    user.ID,
		//}
		////将Session保存到数据库中
		//dao.AddSession(sess)
		////创建一个Cookie，让它与Session相关联
		//cookie := http.Cookie{
		//	Name:     "user",
		//	Value:    uuid,
		//	HttpOnly: true,
		//}
		////将cookie发送给浏览器
		//http.SetCookie(c.Writer, &cookie)
		////t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		////t.Execute(w, user)
		//c.JSON(http.StatusOK, "管理员登录成功") //登录成功界面
		c.JSON(200, gin.H{
			"code": 200,
			"message": map[string]interface{}{
				"res": "管理员登录成功",
			},
		})
	} else {
		//用户名或密码不正确
		//t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		//t.Execute(w, "用户名或密码不正确！")
		//c.JSON(403, "管理员用户名或密码不正确！") //登录界面
		c.JSON(0, gin.H{
			"code": 0,
			"message": map[string]interface{}{
				"res": "管理员用户名或密码不正确！",
			},
		})
	}
}

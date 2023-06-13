package service

import (
	"gin-bookstore/dao"
	"gin-bookstore/model"
	"gin-bookstore/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Logout //处理用户注销的函数
func Logout(c *gin.Context) {
	//获取Cookie
	cookie, _ := c.Request.Cookie("user")
	if cookie != nil {
		//获取cookie的value值
		cookieValue := cookie.Value
		//删除数据库中与之对应的Session
		dao.DeleteSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1
		//将修改之后的cookie发送给浏览器
		http.SetCookie(c.Writer, cookie)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"message": map[string]interface{}{
			"res": "用户退出登录成功",
		},
	})
	//去首页
	//GetPageBooksByPrice(c)
}

func Login(c *gin.Context) {
	//判断是否已经登录
	flag, _ := dao.IsLogin(c.Request)
	//fmt.Println(flag)
	if flag {
		//已经登录
		//去首页
		//GetPageBooksByPrice(c)
		c.JSON(0, gin.H{
			"code": 0,
			"message": map[string]interface{}{
				"res": "已登录",
			},
		})
	} else {
		//未登录
		//获取用户名和密码
		username := c.PostForm("username")
		password := c.PostForm("password")
		//调用userdao中验证用户名和密码的方法
		user, _ := dao.CheckUserNameAndPassword(username, password)
		if user.ID > 0 {
			//用户名和密码正确
			//生成UUID作为Session的id
			uuid := utils.CreateUUID()
			//创建一个Session
			sess := &model.Session{
				SessionID: uuid,
				UserName:  user.Username,
				UserID:    user.ID,
			}
			//将Session保存到数据库中
			dao.AddSession(sess)
			//创建一个Cookie，让它与Session相关联
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode, //设置允许所有第三方网站发送cookie
				Secure:   true,
			}
			//将cookie发送给浏览器
			http.SetCookie(c.Writer, &cookie)
			//t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			//t.Execute(w, user)
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"message": map[string]interface{}{
					"res": "登录成功",
				},
			}) //登录成功界面
		} else {
			//用户名或密码不正确
			//t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			//t.Execute(w, "用户名或密码不正确！")
			//c.JSON(403, "用户名或密码不正确！") //登录界面
			c.JSON(0, gin.H{
				"code": 0,
				"message": map[string]interface{}{
					"res": "用户名或密码不正确",
				},
			}) //登录成功界面
		}
	}
}

// Regist 处理用户的函注册数
func Regist(c *gin.Context) {
	//获取用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已存在
		//t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		//t.Execute(w, "用户名已存在！")
		//c.JSON(401, "用户名已存在") //注册页面
		c.JSON(0, gin.H{
			"code": 0,
			"message": map[string]interface{}{
				"res": "用户名已存在",
			},
		})
	} else {
		//用户名可用，将用户信息保存到数据库中
		dao.SaveUser(username, password, email)
		//用户名和密码正确
		//t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		//t.Execute(w, "")
		//c.JSON(http.StatusOK, "注册成功") //注册成功页面
		c.JSON(200, gin.H{
			"code": 200,
			"message": map[string]interface{}{
				"res": "注册成功",
			},
		})
	}
}

// CheckUserName 通过发送Ajax验证用户名是否可用
func CheckUserName(c *gin.Context) {
	//获取用户输入的用户名
	username := c.PostForm("username")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已存在
		c.Writer.Write([]byte("用户名已存在！"))
	} else {
		//用户名可用
		c.Writer.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}

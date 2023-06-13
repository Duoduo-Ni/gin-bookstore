package service

import (
	"gin-bookstore/dao"
	"gin-bookstore/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func GetBookCover(c *gin.Context) {
	imgpath := c.PostForm("imgpath")

	file, _ := ioutil.ReadFile(imgpath) //把要显示的图片读取到变量中
	c.Writer.WriteString(string(file))  //关键一步，写给前端
}

// Ge222tPageBooksByPrice 获取带分页和价格范围的图书（首页）
func GetPageBooksByPrice(c *gin.Context) {

	//获取页码
	pageNo := c.PostForm("pageNo")
	//获取价格范围
	minPrice := c.PostForm("min")
	maxPrice := c.PostForm("max")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		//调用bookdao中获取带分页的图书的函数
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		//调用bookdao中获取带分页和价格范围的图书的函数
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		//将价格范围设置到page中
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	//调用IsLogin函数判断是否已经登录
	flag, session := dao.IsLogin(c.Request)

	if flag {
		//已经登录，设置page中的IsLogin字段和Username的字段值
		page.IsLogin = true
		page.Username = session.UserName
	}
	////解析模板文件
	//t := template.Must(template.ParseFiles("views/index.html"))
	////执行
	//t.Execute(w, page)
	//page.Code = 200
	//c.JSON(http.StatusOK, page) //首页
	c.JSON(200, gin.H{
		"code": 200,
		"message": map[string]interface{}{
			"res": page,
		},
	})

	//imgpaths := []*string{}
	//for _, book := range page.Books {
	//	//fmt.Println(imgpath)
	//	//fmt.Println("111")
	//	imgpath := (*book).ImgPath
	//	fmt.Println(imgpath)
	//	imgpaths = append(imgpaths, &imgpath)
	//}
	//fmt.Println(imgpaths)
	//file, _ := ioutil.ReadFile(imgpaths) //把要显示的图片读取到变量中
	////service\bookhanldler.go:67:29: cannot use imgpaths (variable of type []*string) as type string in argument to ioutil.ReadFile
	////报错，ioutil包里没找到读写[]string的方法，所以这里改成了用单独的接口/getBookCover，一个一个传递图片
	//c.Writer.WriteString(string(file))   //关键一步，写给前端
}

//// GetPageBooks 获取带分页的图书（图书管理页面）
//func GetPageBooks(c *gin.Context) {
//	//获取页码
//	pageNo := c.PostForm("pageNo")
//	if pageNo == "" {
//		pageNo = "1"
//	}
//	//调用bookdao中获取带分页的图书的函数
//	page, _ := dao.GetPageBooks(pageNo)
//	////解析模板文件
//	//t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
//	////执行
//	//t.Execute(w, page)
//	//c.JSON(http.StatusOK, page) //图书管理页面
//	c.JSON(200, gin.H{
//		"code": 200,
//		"message": map[string]interface{}{
//			"res": page,
//		},
//	})
//}

//// GetPageBooks 获取带分页的图书
//func GetPageBooks(w http.ResponseWriter, r *http.Request) {
//	//获取页码
//	pageNo := r.FormValue("pageNo")
//	if pageNo == "" {
//		pageNo = "1"
//	}
//	//调用bookdao中获取带分页的图书的函数
//	page, _ := dao.GetPageBooks(pageNo)
//	//解析模板文件
//	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
//	//执行
//	t.Execute(w, page)
//}

// AddBook 添加图书
//func AddBook(c *gin.Context) {
//	//获取图书信息
//	title := c.PostForm("title")
//	author := c.PostForm("author")
//	price := c.PostForm("price")
//	sales := c.PostForm("sales")
//	stock := c.PostForm("stock")
//	//将价格、销量和库存进行转换
//	fPrice, _ := strconv.ParseFloat(price, 64)
//	iSales, _ := strconv.ParseInt(sales, 10, 0)
//	iStock, _ := strconv.ParseInt(stock, 10, 0)
//	//创建Book
//	book := &model.Book{
//		Title:   title,
//		Author:  author,
//		Price:   fPrice,
//		Sales:   int(iSales),
//		Stock:   int(iStock),
//		ImgPath: "/static/img/default.jpg",
//	}
//	//调用bookdao中添加图书的函数
//	dao.AddBook(book)
//	c.JSON(http.StatusOK, "添加成功")
//}

//// DeleteBook 删除图书
//func DeleteBook(c *gin.Context) {
//	//获取要删除的图书的id
//	bookID := c.PostForm("bookId")
//	//调用bookdao中获取图书的函数
//	book, _ := dao.GetBookByID(bookID)
//	if book.ID > 0 {
//		//调用bookdao中删除图书的函数
//		dao.DeleteBook(bookID)
//		c.JSON(200, gin.H{
//			"code": 200,
//			"message": map[string]interface{}{
//				"res": "删除图书成功",
//			},
//		})
//	} else {
//		c.JSON(0, gin.H{
//			"code": 0,
//			"message": map[string]interface{}{
//				"res": "输入bookId无效，删除失败",
//			},
//		})
//	}
//	//调用GetBooks处理器函数再次查询一次数据库
//	//GetPageBooks(c) //跳转到图书管理页面
//}

//// DeleteBook 删除图书
//func DeleteBook(w http.ResponseWriter, r *http.Request) {
//	//获取要删除的图书的id
//	bookID := r.FormValue("bookId")
//	//调用bookdao中删除图书的函数
//	dao.DeleteBook(bookID)
//	//调用GetBooks处理器函数再次查询一次数据库
//	GetPageBooks(w, r)
//}

//// ToUpdateOrAddBookPage 去更新或者添加图书的页面
//func ToUpdateOrAddBookPage(c *gin.Context) {
//	//获取要更新的图书的id
//	bookID := c.PostForm("bookId")
//	//调用bookdao中获取图书的函数
//	book, _ := dao.GetBookByID(bookID)
//	if book.ID > 0 {
//		//在更新图书
//		////解析模板
//		//t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
//		////执行
//		//t.Execute(w, book)
//		//c.JSON(http.StatusOK, book) //更新图书
//		c.JSON(200, gin.H{
//			"code": 200,
//			"message": map[string]interface{}{
//				"res1": "输入bookId有效，更新图书",
//				"res2": book,
//			},
//		})
//	} else {
//		//在添加图书
//		////解析模板
//		//t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
//		////执行
//		//t.Execute(w, "")
//		//c.JSON(http.StatusOK, "") //添加图书
//		c.JSON(201, gin.H{
//			"code": 201,
//			"message": map[string]interface{}{
//				"res": "输入bookId为空或无效，添加图书",
//			},
//		})
//	}
//}

//// ToUpdateOrAddBookPage 去更新或者添加图书的页面
//func ToUpdateOrAddBookPage(w http.ResponseWriter, r *http.Request) {
//	//获取要更新的图书的id
//	bookID := r.FormValue("bookId")
//	//调用bookdao中获取图书的函数
//	book, _ := dao.GetBookByID(bookID)
//	if book.ID > 0 {
//		//在更新图书
//		//解析模板
//		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
//		//执行
//		t.Execute(w, book)
//	} else {
//		//在添加图书
//		//解析模板
//		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
//		//执行
//		t.Execute(w, "")
//	}
//}

//// UpdateOrAddBookButton //更新或添加图书的提交按钮
//func UpdateOrAddBookButton(c *gin.Context) {
//	//获取图书信息
//	bookID := c.PostForm("bookId")
//	title := c.PostForm("title")
//	author := c.PostForm("author")
//	price := c.PostForm("price")
//	sales := c.PostForm("sales")
//	stock := c.PostForm("stock")
//	//将价格、销量和库存进行转换
//	fPrice, _ := strconv.ParseFloat(price, 64)
//	iSales, _ := strconv.ParseInt(sales, 10, 0)
//	iStock, _ := strconv.ParseInt(stock, 10, 0)
//	ibookID, _ := strconv.ParseInt(bookID, 10, 0)
//	//创建Book
//	book := &model.Book{
//		ID:      int(ibookID),
//		Title:   title,
//		Author:  author,
//		Price:   fPrice,
//		Sales:   int(iSales),
//		Stock:   int(iStock),
//		ImgPath: "/static/img/default.jpg",
//	}
//	if book.ID > 0 {
//		//在更新图书
//		//调用bookdao中更新图书的函数
//		dao.UpdateBook(book)
//		c.JSON(200, gin.H{
//			"code": 200,
//			"message": map[string]interface{}{
//				"res": "更新图书成功",
//			},
//		})
//	} else {
//		//在添加图书
//		//调用bookdao中添加图书的函数
//		dao.AddBook(book)
//		c.JSON(200, gin.H{
//			"code": 200,
//			"message": map[string]interface{}{
//				"res": "输入bookId为空，添加图书成功",
//			},
//		})
//	}
//	//调用GetBooks处理器函数再次查询一次数据库
//	//GetPageBooks(c) //刷新图书管理页面
//}

//// UpdateOrAddBookButton //更新或添加图书的提交按钮
//func UpdateOrAddBookButton(w http.ResponseWriter, r *http.Request) {
//	//获取图书信息
//	bookID := r.PostFormValue("bookId")
//	title := r.PostFormValue("title")
//	author := r.PostFormValue("author")
//	price := r.PostFormValue("price")
//	sales := r.PostFormValue("sales")
//	stock := r.PostFormValue("stock")
//	//将价格、销量和库存进行转换
//	fPrice, _ := strconv.ParseFloat(price, 64)
//	iSales, _ := strconv.ParseInt(sales, 10, 0)
//	iStock, _ := strconv.ParseInt(stock, 10, 0)
//	ibookID, _ := strconv.ParseInt(bookID, 10, 0)
//	//创建Book
//	book := &model.Book{
//		ID:      int(ibookID),
//		Title:   title,
//		Author:  author,
//		Price:   fPrice,
//		Sales:   int(iSales),
//		Stock:   int(iStock),
//		ImgPath: "/static/img/default.jpg",
//	}
//	if book.ID > 0 {
//		//在更新图书
//		//调用bookdao中更新图书的函数
//		dao.UpdateBook(book)
//	} else {
//		//在添加图书
//		//调用bookdao中添加图书的函数
//		dao.AddBook(book)
//	}
//	//调用GetBooks处理器函数再次查询一次数据库
//	GetPageBooks(w, r)
//}

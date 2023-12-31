package model

// Page 结构
type Page struct {
	//Code        int64   `json:"code"`
	Books       []*Book `json:"books"`       //每页查询出来的图书存放的切片
	PageNo      int64   `json:"pages"`       //当前页
	PageSize    int64   `json:"pageSize"`    //每页显示的条数
	TotalPageNo int64   `json:"totalPageNo"` //总页数，通过计算得到
	TotalRecord int64   `json:"totalRecord"` //总记录数，通过查询数据库得到
	MinPrice    string  `json:"minPrice"`
	MaxPrice    string  `json:"maxPrice"`
	IsLogin     bool    `json:"isLogin"`
	Username    string  `json:"username"`
}

// IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

// IsHasNext 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

// GetPrevPageNo 获取上一页
func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	}
	return 1
}

// GetNextPageNo 获取下一页
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	}
	return p.TotalPageNo

}

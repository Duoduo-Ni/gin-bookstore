package dao

import (
	"gin-bookstore1/model"
	"gin-bookstore1/utils"
)

// CheckUserNameAndPassword 根据用户名和密码从数据库中查询一条记录
func CheckAdminNameAndPassword(adminname string, adminpassword string) (*model.Admin, error) {
	//写sql语句
	sqlStr := "select adminid,adminname,adminpassword from admin where adminname = ? and adminpassword = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, adminname, adminpassword)
	admin := &model.Admin{}
	row.Scan(&admin.AdminID, &admin.AdminName, &admin.AdminPassword)
	return admin, nil
}

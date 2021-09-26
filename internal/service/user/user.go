package userservice

//UserService user service interface
type UserService interface{
	GetUserInfo()
	//...
}

//UserDO User DO
type UserDO struct {
	UserID int
	Username string
	Name string
	// VIP *VIPDO
	//...
}

//GetUserInfo 获取用户信息
func (user *UserDO) GetUserInfo() *UserDO {
	//do something
	//调用UserModel方法获取用户信息
	//把UserModel字段映射到UserDO对象上
	//这里也可以把UserDo转为UserDTO对象再返回，直接供UserBiz相关方法调用
	return user
}

//...
package usermodel

//UserModel user model
type UserModel struct{
	UserID int
	Username string
	Name string
	//...
}

//GetUserByID 示例方法
func (user *UserModel) GetUserByID() *UserModel {
	//从数据库中获取数据
	return user
}

// ...
package userbiz


//UserDTO 用户对象DTO
type UserDTO struct {
	UserID int
	Username string
	Name string
	// ...
}

//RegisterUser 注册用户
func (user *UserDTO) RegisterUser() *UserDTO {
	//do something
	//一般来说，调用UserService的相关方法
	return user
}

//其他的一些方法
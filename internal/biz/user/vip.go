package userbiz


//VIPDTO 会员DTO
type VIPDTO struct{
	User *UserDTO
	VIPNo string
	VIPLevel int
	// ...
}

//GetVIPInfo 获取会员信息，会员等级，会员头像...
func (vip *VIPDTO) GetVIPInfo() *VIPDTO {
	//do something
	//一般来说调用VIPService相关方法
	return vip
}


//...其他的方法
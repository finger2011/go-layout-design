package userservice

//VIPService VIP service interface
type VIPService interface{
	GetVIPInfo()
	//...
}

//VIPDO VIP PO
type VIPDO struct{
	VIPID int
	UserID int
	VIPLevel int
	//...
}

//GetVIPInfo 获取会员信息
func (vip *VIPDO) GetVIPInfo() *VIPDO {
	//do something
	//这个方法可能是在UserDO的方法中被调用
	//从数据库中获取会员信息
	//也可以转化为VIP DTO再返回，根据业务需求来即可
	return vip
}
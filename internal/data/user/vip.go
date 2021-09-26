package usermodel


//VIPModel VIP model
type VIPModel struct {
	VIPID int
	UserID int
	VIPLevel int
	//...
}

//TestVIP 示例方法
func (vip *VIPModel) TestVIP() *VIPModel {
	//从DB中取数据
	return vip
}
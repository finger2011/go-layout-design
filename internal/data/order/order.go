package ordermodel


//OrderModel order model对象
type OrderModel struct {
	OrderID int
	OrderNo string
	UserID int
	Address string
	//...
}

//GetOrderByID 根据order id获取订单信息
func (order *OrderModel) GetOrderByID() *OrderModel {
	//从db中根据order id查询数据 balabala
	return order
}

//其他增删改方法
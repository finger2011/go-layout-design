package orderbiz

//OrderDTO order dto对象
type OrderDTO struct {
	OrderID  int
	OrderNo string
	Username string
	Address  string
	//... 其他相关信息
}

//CreateOrder 创建订单
func (order *OrderDTO) CreateOrder() *OrderDTO {
	//do something
	//一般来说，调用OrderService的相关方法
	return order
}

//GetOrder 获取订单信息
func (order *OrderDTO) GetOrder() *OrderDTO {
	//do something
	//一般来说，调用OrderService的相关方法
	return order
}

//UpdateOrder 更新订单信息
func (order *OrderDTO) UpdateOrder() *OrderDTO {
	//do something
	//一般来说，调用OrderService的相关方法
	return order
}

//DeleteOrder 删除订单
func (order *OrderDTO) DeleteOrder() bool {
	//do something
	//一般来说，调用OrderService的相关方法
	return true
}

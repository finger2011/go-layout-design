package orderservice

//OrderService order service interface
type OrderService interface {
	CreateOrder()
	//...
}

//OrderDO order DO对象
type OrderDO struct {
	OrderID int
	OrderNo string
	Username string
	Address  string
	// ...
}

//CreateOrder 创建订单
func (order *OrderDO) CreateOrder() *OrderDO {
	//do something
	//风险控制，库存检测...
	// 调用OrderModel插入表
	// 后续不影响流程业务过程丢到队列中处理
	return order
}

//...
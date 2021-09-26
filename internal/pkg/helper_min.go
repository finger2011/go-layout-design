package internelpkg

//Min 实例，pkg下存的应该是跟业务无关的公共代码
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
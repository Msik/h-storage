package storage

type User struct {
	id    int
	order Order
}

func (user *User) SetOrder(order Order) {
	user.order = order
}

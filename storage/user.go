package storage

type User struct {
	id    int
	order Order
}

func NewUser(id int) *User {
	return &User{id: id}
}

func (user *User) SetOrder(order Order) {
	user.order = order
}

func (user *User) GetOrder() Order {
	return user.order
}

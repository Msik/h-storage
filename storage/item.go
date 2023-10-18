package storage

type Item struct {
	id    int
	price int
	count int
	name  string
}

func NewItem(name string, id, price, count int) *Item {
	return &Item{id: id, price: price, count: count, name: name}
}

func (item *Item) setCount(count int) {
	item.count = count
}

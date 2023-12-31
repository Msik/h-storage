package storage

type Status int

const (
	New Status = iota
	Paid
	Finished
)

func (s Status) isValid() bool {
	switch s {
	case New, Paid, Finished:
		return true
	}

	return false
}

type Order struct {
	items  []Item
	status Status
}

func NewOrder() *Order {
	return &Order{items: make([]Item, 0), status: New}
}

func (order *Order) SetStatus(status Status) (bool, string) {
	if status.isValid() {
		return false, "the status is invalid"
	}

	order.status = status
	return true, ""
}

func (order *Order) GetItems() []Item {
	return order.items
}

func (order *Order) SetItems(items []Item) {
	order.items = items
}

func (order *Order) GetSum() (sum int) {
	for _, item := range order.items {
		sum += item.price * item.count
	}

	return
}

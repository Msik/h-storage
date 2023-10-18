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

func (order *Order) SetStatus(status Status) (bool, string) {
	if status.isValid() {
		return false, "the status is invalid"
	}

	order.status = status
	return true, ""
}

func (order *Order) SetItems(items []Item) {
	order.items = items
}

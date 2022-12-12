package chapter8

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

type Employee struct {
	Name   string
	Salary float64
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Example struct {
	string
}

func (e Example) Text() string {
	return e.string
}

package customer

type Customer struct {
	Name string
}

func New(name string) *Customer {
	return &Customer{Name: name}
}

func (c *Customer) String() string {
	return c.Name
}

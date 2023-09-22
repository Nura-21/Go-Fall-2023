package courier

type Courier struct {
	position string
	salary   float64
	address  string
}

func (c *Courier) SetPosition(position string) {
	c.position = position
}

func (c *Courier) Position() string {
	return c.position
}

func (c *Courier) SetSalary(salary float64) {
	c.salary = salary
}

func (c *Courier) Salary() float64 {
	return c.salary
}

func (c *Courier) SetAddress(address string) {
	c.address = address
}

func (c *Courier) Address() string {
	return c.address
}

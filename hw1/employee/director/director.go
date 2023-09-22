package director

type Director struct {
	position string
	salary   float64
	address  string
}

func (d *Director) SetPosition(position string) {
	d.position = position
}

func (d *Director) Position() string {
	return d.position
}

func (d *Director) SetSalary(salary float64) {
	d.salary = salary
}

func (d *Director) Salary() float64 {
	return d.salary
}

func (d *Director) SetAddress(address string) {
	d.address = address
}

func (d *Director) Address() string {
	return d.address
}

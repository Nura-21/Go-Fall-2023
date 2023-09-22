package developer

type Developer struct {
	position string
	salary   float64
	address  string
}

func (d *Developer) SetPosition(position string) {
	d.position = position
}

func (d *Developer) Position() string {
	return d.position
}

func (d *Developer) SetSalary(salary float64) {
	d.salary = salary
}

func (d *Developer) Salary() float64 {
	return d.salary
}

func (d *Developer) SetAddress(address string) {
	d.address = address
}

func (d *Developer) Address() string {
	return d.address
}

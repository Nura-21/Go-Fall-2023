package janitor

type Janitor struct {
	position string
	salary   float64
	address  string
}

func (j *Janitor) SetPosition(position string) {
	j.position = position
}

func (j *Janitor) Position() string {
	return j.position
}

func (j *Janitor) SetSalary(salary float64) {
	j.salary = salary
}

func (j *Janitor) Salary() float64 {
	return j.salary
}

func (j *Janitor) SetAddress(address string) {
	j.address = address
}

func (j *Janitor) Address() string {
	return j.address
}

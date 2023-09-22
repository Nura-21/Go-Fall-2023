package employee

type Employee interface {
	SetSalary(salary float64)
	Salary() float64
	SetPosition(position string)
	Position() string
	SetAddress(address string)
	Address() string
}

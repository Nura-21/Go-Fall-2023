package manager

type Manager struct {
	position string
	salary   float64
	address  string
}

func (m *Manager) SetPosition(position string) {
	m.position = position
}

func (m *Manager) Position() string {
	return m.position
}

func (m *Manager) SetSalary(salary float64) {
	m.salary = salary
}

func (m *Manager) Salary() float64 {
	return m.salary
}

func (m *Manager) SetAddress(address string) {
	m.address = address
}

func (m *Manager) Address() string {
	return m.address
}

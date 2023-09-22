package student

type Student struct {
	name    string
	surname string
	id      int
}

func (s *Student) GetId() int {
	return s.id
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) GetUserType() string {
	return "Student"
}

func (s *Student) SetName(name string) {
	s.name = name
}

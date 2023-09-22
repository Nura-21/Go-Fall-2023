package user

type User interface {
	GetId() int
	GetUserType() string
	GetName() string
	SetName(name string)
}

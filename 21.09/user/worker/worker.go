package worker

type Worker struct {
	name    string
	surname string
	id      int
}

func (w *Worker) GetId() int {
	return w.id
}

func (w *Worker) GetName() string {
	return w.name
}

func (w *Worker) GetUserType() string {
	return "Worker"
}

func (w *Worker) SetName(name string) {
	w.name = name
}

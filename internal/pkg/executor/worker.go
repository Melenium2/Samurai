package executor

type Worker interface {
	Work() error
	Done()
}

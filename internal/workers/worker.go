package workers

type Worker interface {
	Process() error
}

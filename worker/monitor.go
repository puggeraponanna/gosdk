package worker

type WorkerPool[T any, U any] interface {
}

type SimpleWorkerPool[T any, U any] struct {
	inputChannel  chan Input[T]
	outputChannel chan Output[U]
}

func NewSimpleWorkerPool[T any, U any](poolSize int, work func(T) U) WorkerPool[T, U] {
	return &SimpleWorkerPool[T, U]{}
}

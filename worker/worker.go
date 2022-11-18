package worker

type Input[T any] struct {
	id   int
	data T
}

type Output[T any] struct {
	data T
	err  error
}

type Worker[T any, U any] interface {
	Work(T) U
}

type SimpleWorker[T any, U any] struct {
	workerFunc func(T) U
}

func NewSimpleWorker[T any, U any](workerFunc func(T) U) Worker[T, U] {
	return &SimpleWorker[T, U]{
		workerFunc: workerFunc,
	}
}

func (sw *SimpleWorker[T, U]) Work(inp T) U {
	return sw.workerFunc(inp)
}

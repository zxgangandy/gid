package worker

type IdAssigner interface {
	AssignWorkerId() int64

	AssignFakeWorkerId() int64
}

type DisposableWorkerIdAssigner struct {
}

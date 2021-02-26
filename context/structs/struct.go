package main

import "context"

type Worker struct {
	ctx context.Context
}

func New(ctx context.Context) *Worker {
	return &Worker{
		ctx: ctx,
	}
}

func (w *Worker) Process() {
	_ = w.ctx
}

func (w *Worker) Fetch() {
	_ = w.ctx
}

func main() {
	w := New(context.Background())
	w.Fetch()
	w.Process()
}

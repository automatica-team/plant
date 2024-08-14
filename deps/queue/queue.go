package queue

import (
	"automatica.team/plant"
	"github.com/golang-queue/queue"
	"github.com/golang-queue/queue/job"
)

func (*Queue) Name() string {
	return "plant/queue"
}

type Queue struct {
	*queue.Queue
}

func New() *Queue {

	return &Queue{queue.NewPool(10)}
}

func (q *Queue) Import(v plant.V) error {
	v.SetDefault("size", 10)

	return nil
}

func (q *Queue) Task(taskFunc job.TaskFunc) error {
	return q.QueueTask(taskFunc)
}

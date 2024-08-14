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
	return &Queue{}
}

func (q *Queue) Import(m plant.M) error {
	q.Queue = queue.NewPool(m["size"].(int))
	return nil
}

func (q *Queue) Task(taskFunc job.TaskFunc) error {
	return q.QueueTask(taskFunc)
}

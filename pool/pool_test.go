package pool

import "testing"

func TestPool(t *testing.T) {
	pool := Pool{
		WorkerNum:   5,
		JobChannel:  make(chan Job),
		WorkerQueue: make(chan chan Job, 5),
		Quit:        make(chan bool),
		QuitQueue:   make(chan chan bool, 5),
	}

	pool.Start()

	for i := 0; i < 5; i++ {
		job := Job{ID: i}
		pool.AddJob(job)
	}

	pool.Stop()
}

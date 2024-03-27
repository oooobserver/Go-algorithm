package pool

import (
	"fmt"
	"sync"
)

type Job struct {
	ID int
}

type Worker struct {
	ID         int
	JobChannel chan Job
	Quit       chan bool
}

type Pool struct {
	WorkerNum   int
	JobChannel  chan Job
	WorkerQueue chan chan Job
	Quit        chan bool
	QuitQueue   chan chan bool
	wg          sync.WaitGroup
}

func NewWorker(id int) Worker {
	return Worker{
		ID:         id,
		JobChannel: make(chan Job),
		Quit:       make(chan bool),
	}
}

func (w Worker) Start(workerQueue chan chan Job, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {
			workerQueue <- w.JobChannel
			fmt.Printf("%d: add\n", w.ID)
			select {
			case job := <-w.JobChannel:
				fmt.Printf("Worker %d started job %d\n", w.ID, job.ID)
				fmt.Printf("Worker %d finished job %d\n", w.ID, job.ID)
			case <-w.Quit:
				fmt.Printf("Worker %d quit\n", w.ID)
				close(w.JobChannel)
				return
			}
		}
	}()
}

func (p *Pool) Start() {
	for i := 0; i < p.WorkerNum; i++ {
		worker := NewWorker(i)
		p.wg.Add(1)
		p.QuitQueue <- worker.Quit
		worker.Start(p.WorkerQueue, &p.wg)
	}

	go func() {
		defer close(p.JobChannel)
		for {
			select {
			case job := <-p.JobChannel:
				worker := <-p.WorkerQueue
				worker <- job
			case <-p.Quit:
				for i := 0; i < p.WorkerNum; i++ {
					q := <-p.QuitQueue
					q <- true
					close(q)
				}
				return
			}
		}
	}()
}

func (p *Pool) AddJob(job Job) {
	p.JobChannel <- job
}

func (p *Pool) Stop() {
	p.Quit <- true
	close(p.Quit)
	p.wg.Wait()
}

package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type Error struct {
	lock  sync.Mutex
	count int
}

type ErrorInterface interface {
	getCount() int
	increment()
}

func (e *Error) getCount() int {
	e.lock.Lock()
	defer e.lock.Unlock()
	return e.count
}

func (e *Error) increment() {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.count++
}

func newError() ErrorInterface {
	e := new(Error)
	e.count = 0
	return e
}

func Run(tasks []Task, n int, m int) error {
	lenTasks := len(tasks)
	if lenTasks == 0 {
		return ErrErrorsLimitExceeded
	}

	var wg sync.WaitGroup
	chanTasks := make(chan Task)
	consumer := make(chan Task)
	err := newError()

	wg.Add(n)
	defer wg.Wait()
	defer close(consumer)

	for i := 0; i < n; i++ {
		go task(chanTasks, consumer, err, &wg)
	}

	if err := checkTask(tasks, err, m, chanTasks); err != nil {
		return err
	}

	return nil
}

func checkTask(tasks []Task, err ErrorInterface, m int, chanTasks chan Task) error {
	for _, task := range tasks {
		v := err.getCount()
		if v >= m {
			return ErrErrorsLimitExceeded
		}
		chanTasks <- task
	}
	return nil
}

func task(chanTasks chan Task, consumer chan Task, error ErrorInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-consumer:
			return
		case task := <-chanTasks:
			err := task()
			if err != nil {
				error.increment()
			}
		}
	}
}

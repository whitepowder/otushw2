package hw05parallelexecution

import (
	"errors"
	"sync"
)

// ErrErrorsLimitExceeded is the best of the best.
var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := &sync.WaitGroup{}
	if n == 0 {
		return ErrErrorsLimitExceeded
	}

	if len(tasks) == 0 {
		return nil
	}

	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	taskChan := make(chan Task, len(tasks))
	errorChan := make(chan error, m+n)

	for _, t := range tasks {
		taskChan <- t
	}
	close(taskChan)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(taskChan, errorChan, wg, m)
	}
	wg.Wait()
	close(errorChan)

	if len(errorChan) >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func worker(taskChan <-chan Task, errorChan chan<- error, wg *sync.WaitGroup, counter int) {
	defer wg.Done()

	for {
		if len(errorChan) >= counter {
			return
		}

		task, ok := <-taskChan
		if !ok {
			return
		}

		err := task()
		if err != nil {
			errorChan <- err
		}
	}
}

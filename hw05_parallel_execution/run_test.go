package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

type test struct {
	n int
	m int
}

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("Zero N and M", func(t *testing.T) {
		tasks:= make([]Task,0)
		result := Run(tasks, 0, 0)
		require.Equal(t, result, ErrErrorsLimitExceeded)
	})

	t.Run("Different N and M", func(t *testing.T) {
		for _, tst :=range [...]test{
			{10,23},
			{0,23},
			{0,1},
			{5,1},
		} {
			tasksCount := 50
			tasks := make([]Task, 0, tasksCount)

			var runTasksCount int32
			var sumTime time.Duration

			for i := 0; i < tasksCount; i++ {
				taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
				sumTime += taskSleep

				tasks = append(tasks, func() error {
					time.Sleep(taskSleep)
					atomic.AddInt32(&runTasksCount, 1)
					return nil
				})
			}

			workersCount := tst.n
			maxErrorsCount := tst.m

			start := time.Now()
			result := Run(tasks, workersCount, maxErrorsCount)
			elapsedTime := time.Since(start)
			require.Nil(t, result)

			require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
			require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")

		}
	})
}

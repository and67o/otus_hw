package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	res := in
	for _, f := range stages {
		res = worker(res, done, f)
	}
	return res
}

func worker(in In, done In, f Stage) Out {
	result := make(Bi)
	go func() {
		defer close(result)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				result <- v
			}
		}
	}()
	return f(result)
}

package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	for _, n := range stages {
		input := make(Bi)
		output(done, input, out)
		out = n(input)
	}
	return out
}

func output(done Out, input Bi, out In) {
	go func() {
		defer close(input)
		for {
			select {
			case <-done:
				return
			case v, ok := <-out:
				if !ok {
					return
				}
				select {
				case <-done:
					return
				case input <- v:
				}
			}
		}
	}()
}

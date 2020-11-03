package run

import (
	"bytes"
	"math/rand"
	"sync"
	"time"
)

type Runner struct {
	buffer bytes.Buffer
	fs     []func()
}

func (r *Runner) Print(s string) {
	r.buffer.WriteString(s)
}

func (r *Runner) String() string {
	return r.buffer.String()
}

func (r *Runner) Register(f ...func()) {
	r.fs = append(r.fs, f...)
}

func (r *Runner) RunTest() {
	r.run(true)
}

func (r *Runner) RunBenchmark() {
	r.run(false)
}

func (r *Runner) ResetBuffer() {
	r.buffer.Reset()
}

func (r *Runner) run(startRandomly bool) {
	var wg sync.WaitGroup
	wg.Add(len(r.fs))
	for i := 0; i < len(r.fs); i++ {
		f := r.fs[i]
		go func() {
			if startRandomly {
				sleepRandomly()
			}
			f()
			wg.Done()
		}()
	}
	wg.Wait()
}

func sleepRandomly() {
	sleepTime := rand.Int63n(100)
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
}

package cake

import ()

// HandlerFunc func
type HandlerFunc func(interface{})

// Func func
type Func func(HandlerFunc) HandlerFunc

// Runner struct
type Runner struct {
	handlers []Func
}

// New inits new Wall instance
func New() *Runner {
	runner := &Runner{
		handlers: []Func{},
	}

	return runner
}

// Use adds handle to a chain
func (runner *Runner) Use(h Func) {
	runner.handlers = append(runner.handlers, h)
}

// Call runs the through chain of handlers
func (runner *Runner) Call(env interface{}) {
	var f HandlerFunc

	if len(runner.handlers) == 1 {
		runner.handlers[0](runner.final)(env)
		return
	}

	for i := len(runner.handlers) - 1; i >= 1; i-- {
		if i == len(runner.handlers)-1 { // find first
			f = runner.handlers[i](runner.final)
		} else {
			f = runner.handlers[i](f)
		}
	}

	runner.handlers[0](f)(env)
}

func (runner *Runner) final(env interface{}) {
}

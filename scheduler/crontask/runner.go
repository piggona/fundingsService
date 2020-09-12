package crontask

import "log"

type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       DataChan
	dataSize   int
	ExecPair
}

func NewRunner(size int, pair ExecPair) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		dataSize:   size,
		ExecPair:   pair,
	}
}

func (r *Runner) startDispatch() {
	for {
		select {
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				log.Printf("Dispatch")
				if err == FinishErr {
					r.Reset()
					r.Error <- CLOSE
				} else if err != nil {
					log.Printf("dispatch error: %s\n")
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			}
			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err == FinishErr {
					r.Reset()
					r.Error <- CLOSE
				} else if err != nil {
					log.Printf("execute error: %s\n")
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		}
	}
}

func (r *Runner) StartAll() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}

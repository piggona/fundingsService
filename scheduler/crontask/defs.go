package crontask

import "fmt"

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "e"
	CLOSE             = "c"
)

type DataChan chan interface{}
type controlChan chan string

var (
	FinishErr = fmt.Errorf("task finished")
)

type ExecPair interface {
	Dispatcher(dc DataChan) error
	Executor(dc DataChan) error
	Reset()
}

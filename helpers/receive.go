package helpers

import (
	"fmt"
	"reflect"

	"github.com/cmd-stream/core-go"
)

type ReceiveFn func(results <-chan core.AsyncResult) (
	asyncResult core.AsyncResult, err error)

type WantReceive struct {
	AsyncResult core.AsyncResult
	Err         error
}

func Receive[T any](results <-chan core.AsyncResult,
	receiveFn ReceiveFn, wantReceive WantReceive,
) (err error) {
	asyncResult, err := receiveFn(results)
	if !reflect.DeepEqual(asyncResult, wantReceive.AsyncResult) {
		return fmt.Errorf("unexpected asyncResult, want %v got %v",
			wantReceive.AsyncResult, asyncResult)
	}
	if err != wantReceive.Err {
		return fmt.Errorf("unexpected receive err, want %v got %v", wantReceive.Err,
			err)
	}
	return nil
}

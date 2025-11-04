package helpers

import (
	"fmt"

	"github.com/cmd-stream/core-go"
)

func Exchange[T any](cmd core.Cmd[T], sendFn SendFn[T], receiveFn ReceiveFn,
	wantSend WantSend, wantReceive WantReceive,
) (err error) {
	results := make(chan core.AsyncResult, 1)
	err = Send(cmd, results, sendFn, wantSend)
	if err != nil {
		err = fmt.Errorf("send: %w", err)
		return
	}
	err = Receive[T](results, receiveFn, wantReceive)
	if err != nil {
		err = fmt.Errorf("receive: %w", err)
	}
	return
}

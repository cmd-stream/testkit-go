package helpers

import (
	"fmt"

	"github.com/cmd-stream/core-go"
)

type SendFn[T any] func(cmd core.Cmd[T], results chan<- core.AsyncResult) (
	seq core.Seq, n int, err error)

type WantSend struct {
	Seq core.Seq
	N   int
	Err error
}

func Send[T any](cmd core.Cmd[T], results chan core.AsyncResult,
	sendFn SendFn[T], wantSend WantSend,
) (err error) {
	seq, n, err := sendFn(cmd, results)
	if seq != wantSend.Seq {
		return fmt.Errorf("unexpected seq, want %v got %v", wantSend.Seq, seq)
	}
	if n != wantSend.N {
		return fmt.Errorf("unexpected n, want %v got %v", wantSend.N, n)
	}
	if err != wantSend.Err {
		return fmt.Errorf("unexpected send err, want %v got %v", wantSend.Err, err)
	}
	return nil
}

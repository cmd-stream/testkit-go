package helpers

import (
	"fmt"

	grp "github.com/cmd-stream/cmd-stream-go/group"
	"github.com/cmd-stream/core-go"
)

type SendGrpFn[T any] func(cmd core.Cmd[T], results chan<- core.AsyncResult) (
	seq core.Seq, clientID grp.ClientID, n int, err error)

type WantSendGrp struct {
	Seq      core.Seq
	ClientID grp.ClientID
	N        int
	Err      error
}

func SendGrp[T any](cmd core.Cmd[T], results chan core.AsyncResult,
	sendFn SendGrpFn[T], wantSend WantSendGrp,
) (err error) {
	seq, clientID, n, err := sendFn(cmd, results)
	if seq != wantSend.Seq {
		return fmt.Errorf("unexpected seq, want %v got %v", wantSend.Seq, seq)
	}
	if clientID != wantSend.ClientID {
		return fmt.Errorf("unexpected clientID, want %v got %v", wantSend.ClientID,
			clientID)
	}
	if n != wantSend.N {
		return fmt.Errorf("unexpected n, want %v got %v", wantSend.N, n)
	}
	if err != wantSend.Err {
		return fmt.Errorf("unexpected send err, want %v got %v", wantSend.Err, err)
	}
	return nil
}

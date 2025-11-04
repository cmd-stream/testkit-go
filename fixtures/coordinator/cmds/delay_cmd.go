package cmds

import (
	"context"
	"fmt"
	"time"

	"github.com/cmd-stream/coordinator-go"
	coordman "github.com/cmd-stream/coordinator-go/manager"
	"github.com/cmd-stream/core-go"
	rcvr "github.com/cmd-stream/testkit-go/fixtures/coordinator/receiver"
	"github.com/cmd-stream/testkit-go/fixtures/coordinator/results"
)

type DelayCmd struct {
	workflow    *coordman.Workflow[*rcvr.Receiver] `json:"-"`
	DelaysCount int                                `json:"delaysCount"`
	BlockOn     int                                `json:"blockOn"`
}

func (c DelayCmd) Exec(ctx context.Context, seq core.Seq, at time.Time,
	receiver *rcvr.Receiver,
	proxy core.Proxy,
) (err error) {
	i := len(c.Workflow().Outcomes())
	if i < c.DelaysCount {
		err = c.Workflow().AppendOutcome("service ID", Outcome(i))
		if err != nil {
			return
		}
		if c.BlockOn != 0 && i == c.BlockOn-1 {
			fmt.Println("block")
			return coordinator.ErrCmdBlocked
		}
		return coordinator.ErrCmdDelayed
	}
	_, err = proxy.Send(seq, results.Result(i))
	if err != nil {
		return
	}
	receiver.Done()
	return
}

func (c DelayCmd) SetWorkflow(
	workflow *coordman.Workflow[*rcvr.Receiver],
) coordman.Cmd[*rcvr.Receiver] {
	c.workflow = workflow
	return c
}

func (c DelayCmd) Workflow() *coordman.Workflow[*rcvr.Receiver] {
	return c.workflow
}

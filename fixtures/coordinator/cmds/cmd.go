package cmds

import (
	"context"
	"time"

	coordman "github.com/cmd-stream/coordinator-go/manager"
	"github.com/cmd-stream/core-go"
	"github.com/cmd-stream/testkit-go/fixtures/coordinator/receiver"
	"github.com/cmd-stream/testkit-go/fixtures/coordinator/results"
)

type Cmd struct {
	workflow      *coordman.Workflow[*receiver.Receiver] `json:"-"`
	OutcomesCount int                                    `json:"outcomesCount"`
}

func (c Cmd) Exec(ctx context.Context, seq core.Seq, at time.Time,
	r *receiver.Receiver,
	proxy core.Proxy,
) (err error) {
	var i int
	for i = range c.OutcomesCount {
		err = c.Workflow().AppendOutcome(coordman.ServiceID("service ID"), Outcome(i))
		if err != nil {
			return
		}
	}
	_, err = proxy.Send(seq, results.Result(i+1))
	if err != nil {
		return
	}
	r.Done()
	return
}

func (c Cmd) SetWorkflow(
	workflow *coordman.Workflow[*receiver.Receiver],
) coordman.Cmd[*receiver.Receiver] {
	c.workflow = workflow
	return c
}

func (c Cmd) Workflow() *coordman.Workflow[*receiver.Receiver] {
	return c.workflow
}

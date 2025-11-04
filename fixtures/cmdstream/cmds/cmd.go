package cmds

import (
	"context"
	"time"

	"github.com/cmd-stream/core-go"
	rcvr "github.com/cmd-stream/testkit-go/fixtures/cmdstream/receiver"
	"github.com/cmd-stream/testkit-go/fixtures/cmdstream/results"
)

type Cmd struct {
	ExecTime time.Duration `json:"execTime"`
}

func (c Cmd) Exec(ctx context.Context, seq core.Seq, at time.Time,
	receiver rcvr.Receiver,
	proxy core.Proxy,
) (err error) {
	time.Sleep(c.ExecTime)
	_, err = proxy.Send(seq, results.Result{LastOneFlag: true})
	return
}

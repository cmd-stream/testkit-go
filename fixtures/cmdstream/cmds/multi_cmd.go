package cmds

import (
	"context"
	"time"

	"github.com/cmd-stream/core-go"
	rcvr "github.com/cmd-stream/testkit-go/fixtures/cmdstream/receiver"
	"github.com/cmd-stream/testkit-go/fixtures/cmdstream/results"
)

type MultiCmd struct {
	ResultsCount int
	ExecTime     time.Duration
}

func (c MultiCmd) Exec(ctx context.Context, seq core.Seq, at time.Time,
	receiver rcvr.Receiver,
	proxy core.Proxy,
) (err error) {
	for i := range c.ResultsCount {
		time.Sleep(c.ExecTime)
		_, err = proxy.Send(seq, results.Result{LastOneFlag: i == c.ResultsCount-1})
		if err != nil {
			return
		}
	}
	return
}

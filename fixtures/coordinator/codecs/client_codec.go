package codecs

import (
	"encoding/json"
	"fmt"

	"github.com/cmd-stream/core-go"
	"github.com/cmd-stream/testkit-go/fixtures/coordinator/cmds"
	rcvr "github.com/cmd-stream/testkit-go/fixtures/coordinator/receiver"
	"github.com/cmd-stream/testkit-go/fixtures/coordinator/results"
	"github.com/cmd-stream/transport-go"
	"github.com/mus-format/dts-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
)

type ClientCodec struct{}

func (c ClientCodec) Encode(cmd core.Cmd[*rcvr.Receiver],
	w transport.Writer,
) (n int, err error) {
	var (
		bs []byte
		n1 int
	)
	switch c := cmd.(type) {
	case cmds.Cmd:
		n, err = dts.DTMSer.Marshal(cmds.CmdDTM, w)
		if err != nil {
			return
		}
		bs, err = json.Marshal(c)
		if err != nil {
			return
		}
		n1, err = ord.ByteSlice.Marshal(bs, w)
		n += n1
		return
	case cmds.DelayCmd:
		n, err = dts.DTMSer.Marshal(cmds.DelayCmdDTM, w)
		if err != nil {
			return
		}
		bs, err = json.Marshal(c)
		if err != nil {
			return
		}
		n1, err = ord.ByteSlice.Marshal(bs, w)
		n += n1
		return
	default:
		panic(fmt.Sprintf("unknown cmd: %T", cmd))
	}
}

func (c ClientCodec) Decode(r transport.Reader) (result core.Result, n int,
	err error,
) {
	bs, n, err := ord.ByteSlice.Unmarshal(r)
	if err != nil {
		return
	}
	res := results.Result(0)
	err = json.Unmarshal(bs, &res)
	if err != nil {
		return
	}
	result = res
	return
}

package codecs

import (
	"encoding/json"
	"fmt"

	"github.com/cmd-stream/core-go"
	"github.com/cmd-stream/testkit-go/fixtures/cmdstream/cmds"
	rcvr "github.com/cmd-stream/testkit-go/fixtures/cmdstream/receiver"
	"github.com/cmd-stream/transport-go"
	"github.com/mus-format/dts-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
)

type ServerCodec struct{}

func (c ServerCodec) Encode(result core.Result, w transport.Writer) (
	n int, err error,
) {
	bs, err := json.Marshal(result)
	if err != nil {
		return
	}
	return ord.ByteSlice.Marshal(bs, w)
}

func (c ServerCodec) Decode(r transport.Reader) (cmd core.Cmd[rcvr.Receiver],
	n int, err error,
) {
	dtm, n, err := dts.DTMSer.Unmarshal(r)
	if err != nil {
		return
	}
	var (
		bs []byte
		n1 int
	)
	switch dtm {
	case cmds.CmdDTM:
		bs, n1, err = ord.ByteSlice.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		c := cmds.Cmd{}
		err = json.Unmarshal(bs, &c)
		if err != nil {
			return
		}
		cmd = c
		return
	case cmds.MultiCmdDTM:
		bs, n, err = ord.ByteSlice.Unmarshal(r)
		n += n1
		if err != nil {
			return
		}
		c := cmds.MultiCmd{}
		err = json.Unmarshal(bs, &c)
		if err != nil {
			return
		}
		cmd = c
		return
	default:
		panic(fmt.Sprintf("unknown dtm: %d", dtm))
	}
}

package codecs

import (
	"encoding/json"

	cmdstream_codec "github.com/cmd-stream/cmd-stream-go/codec"
	"github.com/cmd-stream/core-go"
	"github.com/cmd-stream/testkit-go/fixtures/coordinator/cmds"
	com "github.com/mus-format/common-go"
	"github.com/mus-format/dts-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
)

func AsyncResult(seq core.Seq, result core.Result) core.AsyncResult {
	return core.AsyncResult{
		Seq:       seq,
		BytesRead: ResultSize(seq, result),
		Result:    result,
	}
}

func AsyncErrResult(seq core.Seq, err error) core.AsyncResult {
	return core.AsyncResult{
		Seq:   seq,
		Error: err,
	}
}

func CmdSize(seq core.Seq, cmd cmds.Cmd) (size int) {
	return cmdSize(seq, cmds.CmdDTM, cmd)
}

func DelayCmdSize(seq core.Seq, cmd cmds.DelayCmd) (size int) {
	return cmdSize(seq, cmds.DelayCmdDTM, cmd)
}

func ResultSize(seq core.Seq, result core.Result) (size int) {
	size = cmdstream_codec.SeqMUS.Size(seq)
	bs, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	return size + ord.ByteSlice.Size(bs)
}

func cmdSize[T any](seq core.Seq, dtm com.DTM, t T) (size int) {
	size = cmdstream_codec.SeqMUS.Size(seq)
	size += dts.DTMSer.Size(dtm)
	bs, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return size + ord.ByteSlice.Size(bs)
}

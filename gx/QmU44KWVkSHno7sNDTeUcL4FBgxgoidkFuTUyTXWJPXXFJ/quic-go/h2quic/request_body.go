package h2quic

import (
	"io"

	quic "mbfs/go-mbfs/gx/QmU44KWVkSHno7sNDTeUcL4FBgxgoidkFuTUyTXWJPXXFJ/quic-go"
)

type requestBody struct {
	requestRead bool
	dataStream  quic.Stream
}

// make sure the requestBody can be used as a http.Request.Body
var _ io.ReadCloser = &requestBody{}

func newRequestBody(stream quic.Stream) *requestBody {
	return &requestBody{dataStream: stream}
}

func (b *requestBody) Read(p []byte) (int, error) {
	b.requestRead = true
	return b.dataStream.Read(p)
}

func (b *requestBody) Close() error {
	// stream's Close() closes the write side, not the read side
	return nil
}
package utils

import (
	"sync"

	"github.com/lucas-clemente/quic-go/internal/protocol"
)

var bufferPool sync.Pool

func GetPacketBuffer() []byte {
	return bufferPool.Get().([]byte)
}

func PutPacketBuffer(buf []byte) {
	if cap(buf) != int(protocol.MaxReceivePacketSize) {
		panic("PutPacketBuffer called with packet of wrong size!")
	}
	bufferPool.Put(buf[:0])
}

func init() {
	bufferPool.New = func() interface{} {
		return make([]byte, 0, protocol.MaxReceivePacketSize)
	}
}

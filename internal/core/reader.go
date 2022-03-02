package core

import (
	"github.com/pion/rtcp"
)

// reader is an entity that can read a stream.
type reader interface {
	close()
	onReaderAccepted()
	onReaderPacketRTP(int, *data)
	onReaderPacketRTCP(int, rtcp.Packet)
	onReaderAPIDescribe() interface{}
}

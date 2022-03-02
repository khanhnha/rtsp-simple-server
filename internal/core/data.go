package core

import (
	"time"

	"github.com/pion/rtp"
)

type data struct {
	rtp   *rtp.Packet
	nalus [][]byte
	pts   time.Duration
}

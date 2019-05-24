package contact

import (
	"fmt"

	"github.com/nats-io/go-nats"
)

var (
	url           = "nats://mke.systems:4222"
	subjectGlobal = "global"
)

// ConnectNats is
func ConnectNats() {
	nc, _ := nats.Connect(url)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	c.Subscribe(subjectGlobal, ReciveGlobal)

	fmt.Scanln()
}

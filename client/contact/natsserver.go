package contact

import (
	"bytes"
	"fmt"
	"net"

	"github.com/KMACEL/ASPAR/client/databasecenter"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nats-io/go-nats"
)

var (
	url           = "nats://mke.systems:4222"
	subjectGlobal = "global"
	subjectMotor  = "motor"
)

// ConnectNats is
func ConnectNats() {
	nc, _ := nats.Connect(url)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	c.Subscribe(subjectGlobal, ReciveMessage)
	c.Subscribe(subjectMotor, ReciveMessage)

	fmt.Scanln()
}

// FirstInstallation is
func FirstInstallation() {
	var d databasecenter.DB
	db := d.Open("Client.db")
	d.CreateTable(db)
}

// GetMacAddr is
func GetMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}

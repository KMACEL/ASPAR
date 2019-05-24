package contact

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/KMACEL/ASPAR/common/databasecenter"
	"github.com/nats-io/go-nats"
)

var (
	dataBase databasecenter.DB
	db       *sql.DB
	client   nats.EncodedConn
)

const (
	dataBaseName = "test.db"
	tableName    = "user"
)

var (
	url           = "nats://mke.systems:4222"
	subjectGlobal = "global"
	subjectMotor  = "motor"
)

// ConnectNats is
func ConnectNats() {
	nc, _ := nats.Connect(url)
	client, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer client.Close()

	client.Subscribe(subjectGlobal, ReciveMessage)
	client.Subscribe(subjectMotor, ReciveMessage)

	go databaseConnection()

	fmt.Scanln()
}

func databaseConnection() {
	if _, err := os.Stat(dataBaseName); !os.IsNotExist(err) {
		db = dataBase.Open(dataBaseName)
		defer dataBase.Close(db)
		log.Println("DB is Open...")
	} else {
		db = dataBase.Open(dataBaseName)
		defer dataBase.Close(db)
		createDatabse(db)
		log.Println("DB & Table is Created...")
	}
}

func createDatabse(db *sql.DB) {
	var dataBase databasecenter.DB
	dataBase.CreateTable(db, tableName,
		"deviceID varchar(20) PRIMARY KEY ,"+
			"subject varchar(80)")
}

func getInfo() DeviceInformation {
	//deviceID := GetMacAddr()
	return DeviceInformation{
		DeviceID:  GetMacAddr(),
		TopicList: []string{"global", "motor", "pion", "sensor"}}
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

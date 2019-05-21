package contact

import (
	"encoding/json"
	"fmt"

	"github.com/nats-io/go-nats"
)

// ReciveMessage is get servers
func ReciveMessage(m *nats.Msg) {
	fmt.Println("Subject : ", m.Subject)
	fmt.Println("Received : ", string(m.Data))

	if m.Data != nil {
		parseMessage(m.Subject, &m.Data)
	}
}

func parseMessage(subject string, message *[]byte) {
	if subject == motor {
		motorOp(message)
	}
}

func motorOp(message *[]byte) {
	var motorJSON MotorControlMessageJSON
	json.Unmarshal(*message, &motorJSON)
	fmt.Println(motorJSON)
}

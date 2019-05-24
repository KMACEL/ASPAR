package contact

// SendMessage is
func SendMessage(subject string, message interface{}) {
	go client.Publish(subject, message)
}

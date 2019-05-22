package main

import (
	"sync"

	"github.com/KMACEL/Aspar/server/nats/contact"
)

func main() {
	var wg sync.WaitGroup
	go contact.ConnectNats()

	wg.Add(1)
	wg.Wait()
}

package main

import (
	"sync"

	"github.com/KMACEL/ASPAR/server/nats/contact"
	"github.com/KMACEL/ASPAR/server/pion"
)

func main() {
	var wg sync.WaitGroup
	go contact.ConnectNats()
	go pion.Pion()

	wg.Add(1)
	wg.Wait()
}

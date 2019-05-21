package main

import (
	"sync"

	"github.com/KMACEL/Aspar/client/contact"
)

func main() {
	var wg sync.WaitGroup

	go contact.ConnectNats()

	wg.Add(1)
	wg.Wait()
}

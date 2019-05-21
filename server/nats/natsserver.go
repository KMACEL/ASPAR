package main

import (
	"fmt"

	"github.com/nats-io/go-nats"
)

var (
	// Benim sunucuma kurduğum NATS server. Kurması çok kolay. Siz kendi bilgisayarınızda da kurabilirsiniz.
	// Kurmak için ;
	// 		go get github.com/nats-io/gnatsd
	// Çalıştırmak için ;
	//		gnatsd
	// Uyarı : Eğer kendi bilgisayarınız üzerinden test yapacaksanız, mke.systems yerine localhost yazmanız gerekmektedir
	url     = "nats://localhost:4222"
	subject = "motor"
)

func main() {
	nc, _ := nats.Connect(url)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	// Publisher - Gönderen
	for true {
		c.Publish(subject, `{"subject":"motor","type":"wheel","motors":[{"id":1,"direction":1,"pwm":50},{"id":2,"direction":1,"pwm":50},{"id":3,"direction":0,"pwm":80},{"id":4,"direction":0,"pwm":80}]`)
	}

	// Async Subscriber - Asenkron Alıcı
	c.Subscribe(subject, func(s string) {
		fmt.Printf("Gelen Mesaj: %s\n", s)
	})

	fmt.Scanln()
}

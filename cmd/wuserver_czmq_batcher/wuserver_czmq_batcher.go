//
//  Weather update server.
//  Binds PUB socket to tcp://*:5556
//  Publishes random weather updates
//

package main

import (
	czmq "github.com/zeromq/goczmq"

	"fmt"
	"math/rand"
	"time"
)

func main() {
	SendChan := make(chan [][]byte, 500)
	go func() {
		publisher, err := czmq.NewPub("tcp://*:5556")
		if err != nil {
			panic("")
		}
		defer publisher.Destroy()
		for {
			msgs := <-SendChan
			for _, msg := range msgs {
				publisher.SendFrame(msg, czmq.FlagNone)
			}
		}
	}()

	//  Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// loop for a while apparently
	for {
		msgs := make([][]byte, 0)
		for i := 0; i < 100; i++ {
			//  Get values that will fool the boss
			zipcode := rand.Intn(100000)
			temperature := rand.Intn(215) - 80
			relhumidity := rand.Intn(50) + 10

			//  Send message to all subscribers
			msg := fmt.Sprintf("%05d %d %d", zipcode, temperature, relhumidity)
			msgs = append(msgs, []byte(msg))
		}
		SendChan <- msgs
	}
}

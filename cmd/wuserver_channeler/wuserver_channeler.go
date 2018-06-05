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

	//  Prepare our publisher
	pub := czmq.NewPubChanneler("tcp://*:5556")
	defer pub.Destroy()

	//  Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// loop for a while apparently
	for {

		//  Get values that will fool the boss
		zipcode := rand.Intn(100000)
		temperature := rand.Intn(215) - 80
		relhumidity := rand.Intn(50) + 10

		//  Send message to all subscribers
		msg := fmt.Sprintf("%05d %d %d", zipcode, temperature, relhumidity)
		pub.SendChan <- [][]byte{[]byte(msg)}
	}
}
